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

	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type Manager interface {
	CreateSubscription(subscription models.Subscription) error
	GetSubscriptionList() ([]models.Subscription, error)
	GetSubscription(id string) (models.Subscription, error)
	UpdateSubscription(id string, updatedSubscription models.Subscription) (models.Subscription, error)
	DeleteSubscription(id string) error
}

type Server struct {
	ctx     context.Context
	manager Manager
}

func Init(ctx context.Context, manager Manager) *Server {
	slog.Info("server initialized")
	return &Server{
		ctx:     ctx,
		manager: manager,
	}
}

func (s *Server) setupRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/subscriptions", s.Create)        // Создать новую подписку
	router.Get("/subscriptions", s.GetList)        // Получить список подписок (с фильтрами)
	router.Get("/subscriptions/{id}", s.Get)       // Получить одну подписку
	router.Delete("/subscriptions/{id}", s.Delete) // Удалить подписку cannot use s.Delete (value of type func() error) as http.HandlerFunc value in argument to router.Delete compilerIncompatibleAssign
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	slog.Info("router setup completed", "routes", []string{"POST /subscriptions", "GET /subscriptions", "GET /subscriptions/{id}", "DELETE /subscriptions/{id}"})
	return router
}

func (s *Server) MustRun(srvCfg config.ServerConfig) error {
	router := s.setupRouter()
	address := fmt.Sprintf("%s:%d", srvCfg.Host, srvCfg.Port)

	slog.Info("запуск HTTP сервера", "address", address)
	httpServer := &http.Server{
		Addr:    address,
		Handler: router,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// slog.Error("Ошибка запуска сервера", "error", err)
			log.Fatalf("Ошибка запуска сервера: %s", err.Error())
		}
	}()

	return s.gracefulShutdown(httpServer)
}

func (s *Server) gracefulShutdown(server *http.Server) error {
	shutdownSignals := make(chan os.Signal, 1)
	signal.Notify(shutdownSignals, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)

	select {
	case <-shutdownSignals:
		slog.Info("Получен сигнал завершения работы")
	case <-s.ctx.Done():
		slog.Info("Истекло время ожидания контекста")
	}

	if err := server.Shutdown(s.ctx); err != nil {
		slog.Error("Не удалось корректно завершить работу сервера", "error", err)
		return err
	}

	slog.Info("Сервер успешно завершил работу")
	return nil
}
