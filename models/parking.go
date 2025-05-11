package models

import "fmt"

// ParkingLot represents a parking facility
type ParkingLot struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Capacity    int            `json:"capacity"`
	Spaces      []ParkingSpace `json:"spaces"`
	RatePerHour int            `json:"rate_per_hour"`
}

// ParkingSpace represents an individual parking space
type ParkingSpace struct {
	ID         int         `json:"id"`
	Number     string      `json:"number"`
	IsOccupied bool        `json:"is_occupied"`
	VehicleID  string      `json:"vehicle_id,omitempty"`
	ParkingLot *ParkingLot `json:"parking_lot"`
}

// NewParkingLot creates a new parking lot with the specified capacity
func NewParkingLot(name string, capacity int, ratePerHour int) *ParkingLot {
	lot := &ParkingLot{
		ID:          generateID(),
		Name:        name,
		Capacity:    capacity,
		RatePerHour: ratePerHour,
		Spaces:      make([]ParkingSpace, capacity),
	}

	// Initialize all parking spaces
	for i := 0; i < capacity; i++ {
		spaceNumber := fmt.Sprintf("%03d", i+1) // Format: 001, 002, etc.
		lot.Spaces[i] = ParkingSpace{
			ID:         i + 1,
			Number:     spaceNumber,
			IsOccupied: false,
			ParkingLot: lot,
		}
	}

	return lot
}

// OccupySpace marks a specific parking space as occupied
func (p *ParkingLot) OccupySpace(spaceID int, vehicleID string) bool {
	for i := range p.Spaces {
		if p.Spaces[i].ID == spaceID && !p.Spaces[i].IsOccupied {
			p.Spaces[i].IsOccupied = true
			p.Spaces[i].VehicleID = vehicleID
			return true
		}
	}
	return false
}

// OccupySpace marks a specific parking space as occupied
func (p *ParkingLot) Occupy(vehicleID string) *ParkingSpace {
	for i := range p.Spaces {
		if !p.Spaces[i].IsOccupied {
			p.Spaces[i].IsOccupied = true
			p.Spaces[i].VehicleID = vehicleID
			return &p.Spaces[i]
		}
	}
	return nil
}

// FreeSpace marks a specific parking space as unoccupied
func (p *ParkingLot) FreeSpace(spaceID int) bool {
	for i := range p.Spaces {
		if p.Spaces[i].ID == spaceID && p.Spaces[i].IsOccupied {
			p.Spaces[i].IsOccupied = false
			p.Spaces[i].VehicleID = ""
			return true
		}
	}
	return false
}

// AvailableSpaces returns the number of available parking spaces
func (p *ParkingLot) AvailableSpaces() int {
	count := 0
	for _, space := range p.Spaces {
		if !space.IsOccupied {
			count++
		}
	}
	return count
}

func (p *ParkingLot) OccupiedSpace() int {
	count := 0
	for _, space := range p.Spaces {
		if space.IsOccupied {
			count++
		}
	}
	return count
}
