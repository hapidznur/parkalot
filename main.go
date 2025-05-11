package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hafidhidayatullah/ticketing/models"
)

func main() {
	// This is the entry point of the application
	// You can initialize your server, database connections, etc. here
	Input := `
	create_parking_lot 6
	park KA-01-HH-1234
	park KA-01-HH-9999
	park KA-01-BB-0001
	park KA-01-HH-7777
	park KA-01-HH-2701
	park KA-01-HH-3141
	leave KA-01-HH-3141 4
	status
	park KA-01-P-333
	park DL-12-AA-9999
	leave KA-01-HH-1234 4
	leave KA-01-BB-0001 6
	leave DL-12-AA-9999 2
	park KA-09-HH-0987
	park CA-09-IO-1111
	park KA-09-HH-0123
	status
	`

	// 	output := `Allocated slot number: 1
	// Allocated slot number: 2
	// Allocated slot number: 3
	// Allocated slot number: 4
	// Allocated slot number: 5
	// Allocated slot number: 6
	// Registration number KA-01-HH-3141 with Slot Number 6 is free with Charge $30
	// Slot No. Registration No.
	// 1 KA-01-HH-1234
	// 2 KA-01-HH-9999
	// 3 KA-01-BB-0001
	// 4 KA-01-HH-7777
	// 5 KA-01-HH-2701
	// Allocated slot number: 6
	// Sorry, parking lot is full
	// Registration number KA-01-HH-1234 with Slot Number 1 is free with Charge $30 Registration number
	// KA-01-BB-0001 with Slot Number 3 is free with Charge $50 Registration number DL-12-AA-9999 not
	// found
	// Allocated slot number: 1
	// Allocated slot number: 3
	// Sorry, parking lot is full
	// Slot No. Registration No.
	// 1 KA-09-HH-0987
	// 2 KA-01-HH-9999
	// 3 CA-09-IO-1111
	// 4 KA-01-HH-7777
	// 5 KA-01-HH-2701
	// `

	output := []string{}

	var parkingLot *models.ParkingLot
	tickets := make(map[string]*models.Ticket)
	for _, command := range strings.Split(Input, "\n") {
		command = strings.TrimSpace(command)
		if command == "" {
			continue
		}
		switch {
		case strings.HasPrefix(command, "create_parking_lot"):
			parts := strings.Fields(command)
			if len(parts) != 2 {
				message := "Invalid command format\n"
				output = append(output, message)
				continue
			}
			capacity, err := strconv.Atoi(parts[1])
			if err != nil {
				message := "Invalid capacity: " + err.Error() + "\n"
				output = append(output, message)
				continue
			}
			parkingLot = models.NewParkingLot("Main Parking Lot", capacity, 10.0)
			continue
		case strings.HasPrefix(command, "park"):
			vechileID := strings.Split(command, " ")[1]
			if vechileID == "" {
				output = append(output, "Invalid vehicle ID")
				continue
			}
			space := parkingLot.Occupy(vechileID)
			if space == nil {
				output = append(output, "Sorry, parking lot is full\n")
				continue
			}
			tickets[vechileID] = models.NewTicket(vechileID, *space)
			message := fmt.Sprintf("Allocated slot number: %d\n", parkingLot.OccupiedSpace())
			output = append(output, message)
			continue
		case strings.HasPrefix(command, "leave"):
			leave := strings.Split(command, " ")
			if len(leave) != 3 {
				continue
			}

			if leave[1] == "" {
				continue
			}
			if leave[2] == "" {
				continue
			}

			hourly, err := strconv.Atoi(leave[2])
			if err != nil {
				continue
			}
			if ticket, oke := tickets[leave[1]]; oke {
				ok := parkingLot.FreeSpace(ticket.ParkingSpot.ID)
				if !ok {
					panic("Failed to free space")
				}
				ticket.Checkout(hourly)
				ticket.MarkAsPaid()
				message := fmt.Sprintf("Registration number %s with Slot Number %d is free with Charge $%d",
					ticket.VehicleID, ticket.ParkingSpot.ID, ticket.Fee)
				output = append(output, message)

			} else {
				message := fmt.Sprintf("Registration number %s not found", leave[1])
				output = append(output, message)
			}
			continue
			// Handle leave command
		case strings.HasPrefix(command, "status"):
			var message []string
			// Handle status command
			header := "Slot No. Registration No.\n"
			message = append(message, header)
			for _, space := range parkingLot.Spaces {
				if space.IsOccupied {
					message = append(message, fmt.Sprintf("%d %s\n", space.ID, space.VehicleID))
				}
			}
			output = append(output, message...)
			continue
		default:
			output = append(output, "Unknown command: "+command+"\n")
			continue
		}
	}

	fmt.Print(strings.Join(output, ""))
}
