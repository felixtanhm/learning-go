package venues

import (
	"assignment/models"
	"assignment/src/bst"
	"assignment/src/utils"
	"fmt"

	"github.com/google/uuid"
)

func AddVenue(name string, venueType string, capacity int, venuesP *bst.BST, errChannel chan error) {
	defer close(errChannel)

	newUuid := uuid.New().String()
	newVenue := models.Venue{
		Uuid:     newUuid,
		Name:     name,
		Type:     venueType,
		Capacity: capacity,
	}
	venuesP.Insert(newVenue)
	if err := utils.WriteVenues(venuesP); err != nil {
		errChannel <- fmt.Errorf("error writing to venue CSV: %v", err)
	}
	errChannel <- nil
}
