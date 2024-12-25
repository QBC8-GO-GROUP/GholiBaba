package company

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/QBC8-GO-GROUP/GholiBaba/internal/company/domain"
	"github.com/QBC8-GO-GROUP/GholiBaba/internal/company/port"

	"github.com/QBC8-GO-GROUP/GholiBaba/pkg/logger"
)

var (
	ErrCompanyOnCreate           = errors.New("error on creating new company")
	ErrCompanyCreationValidation = errors.New("validation failed")
	ErrCompanyNotFound           = errors.New("company not found")
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}

// AddTravel creates a new travel entry
func (s *service) AddTravel(w http.ResponseWriter, r *http.Request) {
	var travel Travel
	if err := json.NewDecoder(r.Body).Decode(&travel); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	travel.ID = fmt.Sprintf("travel-%d", time.Now().UnixNano())
	s.mu.Lock()
	s.travels[travel.ID] = &travel
	s.mu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(travel)
}

func (rs *service) GetTravel(ctx context.Context, travelID domain.TravelID) (*domain.Travel, error) {
	travel, err := rs.repo.Get(ctx, travelID)
	if err != nil {
		logger.Error("travel not found", nil)
		return nil, err
	}
	if travel == nil || travel.ID == 0 {
		logger.Error("travel not found", nil)
		return nil, errors.New("travel no found")
	}
	logger.Info("successful get travel", nil)
	return travel, nil
}

func (rs *service) UpdateTravel(ctx context.Context, travel domain.Travel) error {
	err := rs.repo.Update(ctx, travel)
	if err != nil {
		logger.Error("error in updating travel", nil)
		return err
	}
	logger.Info("successful update travel", nil)
	return nil
}

func (rs *service) DeleteTravel(ctx context.Context, travelID domain.TravelID) error {
	err := rs.repo.Delete(ctx, travelID)
	if err != nil {
		logger.Error("error in deleting travel", nil)
		return err
	}
	logger.Info("successful delete travel", nil)
	return nil
}

// GetTravels lists all available travels
func (s *service) GetTravels(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()

	travels := make([]*Travel, 0, len(s.travels))
	for _, travel := range s.travels {
		travels = append(travels, travel)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(travels)
}

// BookTravel reserves a seat for a travel entry
func (s *service) BookTravel(w http.ResponseWriter, r *http.Request) {
	type BookingRequest struct {
		TravelID string `json:"travel_id"`
	}

	var req BookingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	travel, exists := s.travels[req.TravelID]
	if !exists {
		http.Error(w, "Travel not found", http.StatusNotFound)
		return
	}

	if travel.Available <= 0 {
		http.Error(w, "No available seats", http.StatusConflict)
		return
	}

	travel.Available--
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Booking successful"})
}

// CancelBooking cancels a booking and frees up a seat
func (s *service) CancelBooking(w http.ResponseWriter, r *http.Request) {
	type CancelRequest struct {
		TravelID string `json:"travel_id"`
	}

	var req CancelRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	travel, exists := s.travels[req.TravelID]
	if !exists {
		http.Error(w, "Travel not found", http.StatusNotFound)
		return
	}

	travel.Available++
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Booking canceled"})
}
