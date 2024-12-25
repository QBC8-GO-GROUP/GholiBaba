package domain

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type (
	TravelID uint
	OwnerID  uint
)

type Travel struct {
	ID          TravelID
	CreatedAt   time.Time
	DeletedAt   time.Time
	UpdatedAt   time.Time
	Owner       OwnerID   `json:"company_id"`
	Type        string    `json:"type"` // Bus, Train, Flight, Ship
	Source      string    `json:"source"`
	Destination string    `json:"destination"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Price       float64   `json:"price"`
	Seats       int       `json:"seats"`
	Available   int       `json:"available"`
}

// TravelAgencyService manages travel offerings and bookings
type TravelAgencyService struct {
	mu      sync.Mutex
	travels map[string]*Travel // ID to Travel
}

// NewTravelAgencyService initializes the service
func NewTravelAgencyService() *TravelAgencyService {
	return &TravelAgencyService{
		travels: make(map[string]*Travel),
	}
}

// AddTravel creates a new travel entry
func (s *TravelAgencyService) AddTravel(w http.ResponseWriter, r *http.Request) {
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

// GetTravels lists all available travels
func (s *TravelAgencyService) GetTravels(w http.ResponseWriter, r *http.Request) {
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
func (s *TravelAgencyService) BookTravel(w http.ResponseWriter, r *http.Request) {
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
func (s *TravelAgencyService) CancelBooking(w http.ResponseWriter, r *http.Request) {
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

func main() {
	agencyService := NewTravelAgencyService()

	http.HandleFunc("/travels", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			agencyService.AddTravel(w, r)
		} else if r.Method == http.MethodGet {
			agencyService.GetTravels(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/book", agencyService.BookTravel)
	http.HandleFunc("/cancel", agencyService.CancelBooking)

	fmt.Println("Travel Agency Service is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
