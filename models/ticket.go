package models

import (
	"time"
)

// Ticket represents a parking ticket issued to a vehicle
type Ticket struct {
	ID          string        `json:"id"`
	VehicleID   string        `json:"vehicle_id"`
	EntryTime   time.Time     `json:"entry_time"`
	ExitTime    time.Time     `json:"exit_time,omitempty"`
	TotalHours  int           `json:"total_hours,omitempty"`
	Fee         int           `json:"fee,omitempty"`
	IsPaid      bool          `json:"is_paid"`
	ParkingSpot *ParkingSpace `json:"parking_spot"`
}

// NewTicket creates a new parking ticket with the current time as entry time
func NewTicket(vehicleID string, parkingSpot ParkingSpace) *Ticket {
	return &Ticket{
		ID:          generateID(), // You would implement this function
		VehicleID:   vehicleID,
		EntryTime:   time.Now(),
		IsPaid:      false,
		ParkingSpot: &parkingSpot,
	}
}

// CalculateFee calculates the parking fee based on duration
func (t *Ticket) CalculateFee() int {
	if t.TotalHours <= 2 {
		t.Fee = int(t.ParkingSpot.ParkingLot.RatePerHour)
	} else {
		t.Fee = int(t.ParkingSpot.ParkingLot.RatePerHour)*(t.TotalHours-2) + int(t.ParkingSpot.ParkingLot.RatePerHour)
	}
	return t.Fee
}

// Checkout sets the exit time and calculates the fee
func (t *Ticket) Checkout(totalHour int) {
	t.ExitTime = time.Now()
	t.TotalHours = totalHour
	t.CalculateFee()
}

// MarkAsPaid marks the ticket as paid
func (t *Ticket) MarkAsPaid() {
	t.IsPaid = true
}

// Helper function to generate a unique ID (placeholder)
func generateID() string {
	// In a real application, use a proper UUID generator
	return time.Now().Format("20060102150405")
}
