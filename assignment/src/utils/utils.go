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

const venueDB = "./data/venueDB.csv"
const bookingDB = "./data/bookingDB.csv"

func LoadData(venuesP *models.BST) error {
	venueMap := make(map[string]*models.Venue)
	if err := readVenues(venuesP, venueMap); err != nil {
		return err
	}
	if err := readBookings(venueMap); err != nil {
		return err
	}
	return nil
}

func readVenues(venuesP *models.BST, venuesMap map[string]*models.Venue) error {
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
			Uuid:        row[0],
			Name:        row[1],
			Type:        row[2],
			Capacity:    capacity,
			BookingList: models.CreateLinkedList(),
		}
		venuesP.Insert(newVenue)
		venuesMap[row[0]] = &newVenue
	}
	return nil
}

func WriteVenues(venuesP *models.BST) error {
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

	venuesList := venuesP.GetAll("inOrder")
	for _, venue := range venuesList {
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

func readBookings(venuesMap map[string]*models.Venue) error {
	defer handlePanic()
	file, err := os.Open(bookingDB)
	if err != nil {
		return fmt.Errorf("error opening booking CSV: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("error reading booking CSV: %v", err)
	}
	for i, row := range records {
		if i == 0 {
			continue
		}
		newBooking := models.Booking{
			Uuid:      row[0],
			VenueID:   row[1],
			UserEmail: row[2],
		}
		venue := venuesMap[row[1]]
		venue.BookingList.Insert(newBooking)
	}
	return nil
}

func WriteBookings(venuesP *models.BST) error {
	venueMutex.Lock()
	defer venueMutex.Unlock()
	defer handlePanic()
	file, err := os.Create(bookingDB)
	if err != nil {
		return fmt.Errorf("error opening booking CSV: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"ID", "VenueID", "UserEmail"}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("error writing header to booking CSV: %v", err)
	}

	venuesList := venuesP.GetAll("inOrder")
	for _, venue := range venuesList {
		if venue.BookingList != nil {
			bookingsList := venue.BookingList.GetAll()
			for _, booking := range bookingsList {
				record := []string{
					booking.Uuid, booking.VenueID, booking.UserEmail,
				}
				if err := writer.Write(record); err != nil {
					return fmt.Errorf("error writing to booking CSV: %v", err)
				}
			}
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
