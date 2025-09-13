package server

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"subscription-aggregator-api/manager"
	"subscription-aggregator-api/models"
	"subscription-aggregator-api/storage"

	"github.com/go-chi/chi"
)

const (
	ErrSubscriptionNotFound  = "Subscription not found"
	ErrSubscriptionsNotFound = "Subscriptions not found"
	ErrInternalServerError   = "Internal Server Error"

	StatusCreated = "created"
	StatusUpdated = "updated"
)

// ErrorResponse описывает тело ошибки
// swagger:model ErrorResponse
type ErrorResponse struct {
	Error string `json:"error"`
}

// Response описывает ответ с результатом операции
// swagger:model Response
type Response struct {
	Status string `json:"status"`
}

// TotalSumResponse описывает ответ с суммой подписок
// swagger:model TotalSumResponse
type TotalSumResponse struct {
	TotalSum int `json:"total_sum"`
}

func writeJSON[T any](w http.ResponseWriter, status int, data T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		slog.Error("Failed to write JSON response", "error", err)
	}
}

func writeErrorJSON(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, ErrorResponse{Error: message})
}

// @Summary      Создать подписку
// @Description  Создаёт новую подписку
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        subscription  body      models.Subscription  true  "Подписка"
// @Success      201           {object}  ResultResponse
// @Failure      400           {object}  ErrorResponse
// @Failure      405           {object}  ErrorResponse
// @Failure      500           {object}  ErrorResponse
// @Router       /subscriptions [post]
func (s *Server) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var subscription models.Subscription
	if err := json.NewDecoder(r.Body).Decode(&subscription); err != nil {
		slog.Error("Failed to decode Subscription from JSON", "error", err)
		writeErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := s.manager.CreateSubscription(subscription); err != nil {
		s.handleSubscriptionError(w, err)
		return
	}

	slog.Info("Subscription created successfully", "service_name", subscription.ServiceName, "user_id", subscription.UserID)
	writeJSON(w, http.StatusCreated, Response{Status: StatusCreated})
}

// @Summary      Получить информацию о подписке
// @Description  Возвращает подписку по ID
// @Tags         subscriptions
// @Produce      json
// @Param        id   path      string  true  "ID подписки"
// @Success      200  {object}  models.Subscription
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /subscriptions/{id} [get]
func (s *Server) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	subscription, err := s.manager.GetSubscription(id)
	if err != nil {
		s.handleSubscriptionError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, subscription)
	slog.Info("Subscription retrieved successfully", "id", id)
}

// @Summary      Получить список подписок
// @Description  Возвращает список всех подписок
// @Tags         subscriptions
// @Produce      json
// @Success      200  {array}   models.Subscription
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /subscriptions [get]
func (s *Server) GetList(w http.ResponseWriter, r *http.Request) {
	subscriptions, err := s.manager.GetAllSubscriptions()
	if err != nil {
		s.handleSubscriptionError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, subscriptions)
	slog.Info("Subscriptions list retrieved successfully", "count", len(subscriptions))
}

// @Summary      Обновить информацию о подписке
// @Description  Обновляет подписку по ID
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "ID подписки"
// @Param        subscription  body      models.Subscription  true  "Обновлённая подписка"
// @Success      200           {object}  ResultResponse
// @Failure      400           {object}  ErrorResponse
// @Failure      404           {object}  ErrorResponse
// @Failure      500           {object}  ErrorResponse
// @Router       /subscriptions/{id} [put]
func (s *Server) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	defer r.Body.Close()

	var subscription models.Subscription
	if err := json.NewDecoder(r.Body).Decode(&subscription); err != nil {
		slog.Error("Failed to decode Subscription from JSON", "error", err)
		writeErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := s.manager.UpdateSubscription(id, subscription); err != nil {
		s.handleSubscriptionError(w, err)
		return
	}

	slog.Info("Subscription updated successfully", "id", id)
	writeJSON(w, http.StatusOK, Response{Status: StatusUpdated})
}

// @Summary      Удалить подписку
// @Description  Удаляет подписку по ID
// @Tags         subscriptions
// @Produce      json
// @Param        id   path      string  true  "ID подписки"
// @Success      204  "No Content"
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /subscriptions/{id} [delete]
func (s *Server) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := s.manager.DeleteSubscription(id); err != nil {
		s.handleSubscriptionError(w, err)
		return
	}

	slog.Info("Subscription deleted successfully", "id", id)
	w.WriteHeader(http.StatusNoContent)
}

// @Summary      Получить сумму всех подписок
// @Description  Возвращает сумму цен всех подписок
// @Tags         subscriptions
// @Produce      json
// @Success      200  {object}  TotalSumResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /subscriptions/sum [get]
func (s *Server) GetSum(w http.ResponseWriter, r *http.Request) {
	totalSum, err := s.manager.GetAllSubscriptionsSum()
	if err != nil {
		s.handleSubscriptionError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, TotalSumResponse{TotalSum: totalSum})
	slog.Info("Total subscription price sum retrieved successfully", "total_sum", totalSum)
}

func (s *Server) handleSubscriptionError(w http.ResponseWriter, err error) {
	var badReqErr *manager.BadRequestError

	switch {
	case errors.As(err, &badReqErr):
		slog.Error(err.Error(), "error", err)
		writeErrorJSON(w, http.StatusBadRequest, err.Error())
	case errors.Is(err, storage.ErrSubscriptionNotFound):
		slog.Error(err.Error())
		writeErrorJSON(w, http.StatusNotFound, ErrSubscriptionNotFound)
	case errors.Is(err, storage.ErrNoSubscriptions):
		slog.Error(err.Error())
		writeErrorJSON(w, http.StatusNotFound, ErrSubscriptionsNotFound)
	default:
		slog.Error("internal server error", "error", err)
		writeErrorJSON(w, http.StatusInternalServerError, ErrInternalServerError)
	}
}
