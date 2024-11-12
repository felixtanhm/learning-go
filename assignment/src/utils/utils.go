package utils

import (
	"assignment/models"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"sync"
)

var venueMutex sync.Mutex
var bookingMutex sync.Mutex

const venueDB = "./data/venueDB.csv"
const bookingDB = "./data/bookingDB.csv"

func ReadVenues(venuesP *[]models.Venue) error {
	defer handlePanic()
	file, err := os.Open(venueDB)
	if err != nil {
		return fmt.Errorf("error opening venue CSV: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("error reading venue CSV: %v", err)
	}
	for i, row := range records {
		if i == 0 {
			continue
		}
		capacity, _ := strconv.Atoi(row[3])
		newVenue := models.Venue{
			Uuid:     row[0],
			Name:     row[1],
			Type:     row[2],
			Capacity: capacity,
		}
		*venuesP = append(*venuesP, newVenue)
	}
	return nil
}

func WriteVenues(venuesP *[]models.Venue) error {
	venueMutex.Lock()
	defer venueMutex.Unlock()
	defer handlePanic()
	file, err := os.Create(venueDB)
	if err != nil {
		return fmt.Errorf("error opening venue CSV: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"ID", "Name", "Type", "Capacity"}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("error writing header to venue CSV: %v", err)
	}

	for _, venue := range *venuesP {
		capacity := strconv.Itoa(venue.Capacity)
		record := []string{
			venue.Uuid, venue.Name, venue.Type, capacity,
		}
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("error writing to venue CSV: %v", err)
		}
	}

	return nil
}

func handlePanic() error {
	if r := recover(); r != nil {
		return fmt.Errorf("unhandled system error: %v", r)
	}
	return nil
}
