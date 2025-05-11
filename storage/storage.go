package storage

import (
	"github.com/hafidhidayatullah/ticketing/models"
)

// TicketStorage defines the interface for ticket storage operations
type TicketStorage interface {
	SaveTicket(ticket *models.Ticket) error
	GetTicket(ticketID string) (*models.Ticket, error)
	UpdateTicket(ticket *models.Ticket) error
	DeleteTicket(ticketID string) error
	GetAllTickets() []*models.Ticket
	GetTicketsByVehicle(vehicleID string) []*models.Ticket
}

// ParkingStorage defines the interface for parking lot storage operations
type ParkingStorage interface {
	SaveParkingLot(lot *models.ParkingLot) error
	GetParkingLot(lotID string) (*models.ParkingLot, error)
	UpdateParkingLot(lot *models.ParkingLot) error
	DeleteParkingLot(lotID string) error
	GetAllParkingLots() []*models.ParkingLot
}

// Store is a combined interface for all storage operations
type Store interface {
	TicketStorage
	ParkingStorage
}
