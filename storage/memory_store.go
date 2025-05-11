package storage

import (
	"errors"
	"sync"

	"github.com/hafidhidayatullah/ticketing/models"
)

// MemoryStore provides an in-memory implementation of ticket and parking storage
type MemoryStore struct {
	tickets     map[string]*models.Ticket
	parkingLots map[string]*models.ParkingLot
	mutex       sync.RWMutex
}

// NewMemoryStore creates a new in-memory storage instance
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		tickets:     make(map[string]*models.Ticket),
		parkingLots: make(map[string]*models.ParkingLot),
	}
}

// Ticket storage operations

// SaveTicket stores a ticket in memory
func (s *MemoryStore) SaveTicket(ticket *models.Ticket) error {
	if ticket == nil {
		return errors.New("cannot save nil ticket")
	}
	if ticket.ID == "" {
		return errors.New("ticket must have an ID")
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.tickets[ticket.ID] = ticket
	return nil
}

// GetTicket retrieves a ticket by ID
func (s *MemoryStore) GetTicket(ticketID string) (*models.Ticket, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	ticket, exists := s.tickets[ticketID]
	if !exists {
		return nil, errors.New("ticket not found")
	}

	return ticket, nil
}

// UpdateTicket updates an existing ticket
func (s *MemoryStore) UpdateTicket(ticket *models.Ticket) error {
	if ticket == nil {
		return errors.New("cannot update nil ticket")
	}
	if ticket.ID == "" {
		return errors.New("ticket must have an ID")
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Check if the ticket exists
	_, exists := s.tickets[ticket.ID]
	if !exists {
		return errors.New("ticket not found")
	}

	s.tickets[ticket.ID] = ticket
	return nil
}

// DeleteTicket removes a ticket from storage
func (s *MemoryStore) DeleteTicket(ticketID string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	_, exists := s.tickets[ticketID]
	if !exists {
		return errors.New("ticket not found")
	}

	delete(s.tickets, ticketID)
	return nil
}

// GetAllTickets returns all tickets in storage
func (s *MemoryStore) GetAllTickets() []*models.Ticket {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	tickets := make([]*models.Ticket, 0, len(s.tickets))
	for _, t := range s.tickets {
		tickets = append(tickets, t)
	}

	return tickets
}

// GetTicketsByVehicle returns all tickets for a specific vehicle
func (s *MemoryStore) GetTicketsByVehicle(vehicleID string) []*models.Ticket {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var tickets []*models.Ticket
	for _, t := range s.tickets {
		if t.VehicleID == vehicleID {
			tickets = append(tickets, t)
		}
	}

	return tickets
}

// Parking lot storage operations

// SaveParkingLot stores a parking lot in memory
func (s *MemoryStore) SaveParkingLot(lot *models.ParkingLot) error {
	if lot == nil {
		return errors.New("cannot save nil parking lot")
	}
	if lot.ID == "" {
		return errors.New("parking lot must have an ID")
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.parkingLots[lot.ID] = lot
	return nil
}

// GetParkingLot retrieves a parking lot by ID
func (s *MemoryStore) GetParkingLot(lotID string) (*models.ParkingLot, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	lot, exists := s.parkingLots[lotID]
	if !exists {
		return nil, errors.New("parking lot not found")
	}

	return lot, nil
}

// UpdateParkingLot updates an existing parking lot
func (s *MemoryStore) UpdateParkingLot(lot *models.ParkingLot) error {
	if lot == nil {
		return errors.New("cannot update nil parking lot")
	}
	if lot.ID == "" {
		return errors.New("parking lot must have an ID")
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Check if the parking lot exists
	_, exists := s.parkingLots[lot.ID]
	if !exists {
		return errors.New("parking lot not found")
	}

	s.parkingLots[lot.ID] = lot
	return nil
}

// DeleteParkingLot removes a parking lot from storage
func (s *MemoryStore) DeleteParkingLot(lotID string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	_, exists := s.parkingLots[lotID]
	if !exists {
		return errors.New("parking lot not found")
	}

	delete(s.parkingLots, lotID)
	return nil
}

// GetAllParkingLots returns all parking lots in storage
func (s *MemoryStore) GetAllParkingLots() []*models.ParkingLot {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	lots := make([]*models.ParkingLot, 0, len(s.parkingLots))
	for _, l := range s.parkingLots {
		lots = append(lots, l)
	}

	return lots
}
