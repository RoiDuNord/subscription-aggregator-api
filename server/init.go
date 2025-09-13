package server

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"subscription-aggregator-api/config"
	"subscription-aggregator-api/models"
	"syscall"

	_ "subscription-aggregator-api/docs"

	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type SubscriptionManager interface {
	CreateSubscription(subscription models.Subscription) error
	GetSubscription(id string) (models.Subscription, error)
	GetAllSubscriptions() ([]models.Subscription, error)
	UpdateSubscription(id string, updatedSubscription models.Subscription) error
	DeleteSubscription(id string) error
	GetAllSubscriptionsSum() (totalSum int, err error)
}

type Server struct {
	ctx        context.Context
	manager    SubscriptionManager
	httpServer *http.Server
}

func Init(ctx context.Context, manager SubscriptionManager) *Server {
	slog.Info("Server initialized")
	return &Server{
		ctx:     ctx,
		manager: manager,
	}
}

// @title Subscription Aggregator API
// @version 1.0
// @description API for Managing Subscriptions
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http https

func (s *Server) setupRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/subscriptions", s.Create)
	router.Get("/subscriptions/{id}", s.Get)
	router.Get("/subscriptions", s.GetList)
	router.Get("/subscriptions/sum", s.GetSum)
	router.Put("/subscriptions/{id}", s.Update)
	router.Delete("/subscriptions/{id}", s.Delete)
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	slog.Info("Endpoints successfully configured")
	return router
}

func (s *Server) MustRun(srvCfg config.ServerConfig) error {
	router := s.setupRouter()
	address := fmt.Sprintf("%s:%d", srvCfg.Host, srvCfg.Port)

	httpServer := &http.Server{
		Addr:    address,
		Handler: router,
	}

	slog.Info("Starting HTTP server", "address", address)
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %s", err.Error())
		}
	}()

	return s.gracefulShutdown(httpServer)
}

func (s *Server) gracefulShutdown(server *http.Server) error {
	shutdownSignals := make(chan os.Signal, 1)
	signal.Notify(shutdownSignals, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)

	select {
	case <-shutdownSignals:
		slog.Info("Received shutdown signal")
	case <-s.ctx.Done():
		slog.Info("Context timeout reached")
	}

	if err := server.Shutdown(s.ctx); err != nil {
		slog.Error("Failed to gracefully shut down server", "error", err)
		return err
	}

	slog.Info("Server shut down successfully")
	return nil
}
