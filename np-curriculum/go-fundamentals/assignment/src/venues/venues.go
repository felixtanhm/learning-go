package venues

import (
	"assignment/models"
	"assignment/src/utils"
	"fmt"

	"github.com/google/uuid"
)

func AddVenue(name string, venueType string, capacity int, venuesP *models.BST, errChannel chan error) {
	defer close(errChannel)

	if name == "" {
		errChannel <- fmt.Errorf("venue name cannot be empty")
		return
	}
	if venueType != "room" && venueType != "hall" {
		errChannel <- fmt.Errorf("invalid venue type: %v", venueType)
		return
	}
	if capacity <= 0 {
		errChannel <- fmt.Errorf("venue capacity must be a positive integer")
		return
	}
	existingVenue := venuesP.GetOne(name)
	if existingVenue != nil {
		errChannel <- fmt.Errorf("a venue with the name %v already exists", name)
		return
	}

	newUuid := uuid.New().String()
	newVenue := models.Venue{
		Uuid:        newUuid,
		Name:        name,
		Type:        venueType,
		Capacity:    capacity,
		BookingList: models.CreateLinkedList(),
	}
	if err := venuesP.Insert(newVenue); err != nil {
		errChannel <- fmt.Errorf("error adding venue: %v", err)
	}

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
