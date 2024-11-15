package venues

import (
	"assignment/models"
	"assignment/src/utils"
	"fmt"

	"github.com/google/uuid"
)

func AddVenue(name string, venueType string, capacity int, venuesP *models.BST, errChannel chan error) {
	defer close(errChannel)

	newUuid := uuid.New().String()
	newVenue := models.Venue{
		Uuid:        newUuid,
		Name:        name,
		Type:        venueType,
		Capacity:    capacity,
		BookingList: models.CreateLinkedList(),
	}
	venuesP.Insert(newVenue)
	if err := utils.WriteVenues(venuesP); err != nil {
		errChannel <- fmt.Errorf("error writing to venue CSV: %v", err)
	}
}

func BookVenue(userEmail string, venuesP *models.BST, venue *models.Venue, errChannel chan error) {
	defer close(errChannel)

	newUuid := uuid.New().String()
	newBooking := models.Booking{
		Uuid:      newUuid,
		VenueID:   venue.Uuid,
		UserEmail: userEmail,
	}

	venue.BookingList.Insert(newBooking)
	if err := utils.WriteBookings(venuesP); err != nil {
		errChannel <- fmt.Errorf("error writing to booking csv %v", err)
	}
}
