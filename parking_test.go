package main

import (
	"testing"

	"github.com/hafidhidayatullah/ticketing/models"
)

func TestCreateParking(t *testing.T) {
	parking := models.NewParkingLot("Test Parking", 10, 5.0)
	if parking.Name != "Test Parking" {
		t.Errorf("Expected parking name to be 'Test Parking', got '%s'", parking.Name)
	}
	if parking.Capacity != 10 {
		t.Errorf("Expected parking capacity to be 10, got %d", parking.Capacity)
	}
	if parking.RatePerHour != 5.0 {
		t.Errorf("Expected parking rate per hour to be 5.0, got %f", parking.RatePerHour)
	}
}
