package venues

import (
	"assignment/models"
	"assignment/src/utils"
	"fmt"

	"github.com/google/uuid"
)

func AddVenue(name string, venueType string, capacity int, venuesP *[]models.Venue, errChannel chan error) {
	defer close(errChannel)

	newUuid := uuid.New().String()
	newVenue := models.Venue{
		Uuid:     newUuid,
		Name:     name,
		Type:     venueType,
		Capacity: capacity,
	}
	*venuesP = append(*venuesP, newVenue)
	if err := utils.WriteVenues(venuesP); err != nil {
		errChannel <- fmt.Errorf("error opening csv data file: %v", err)
	}
	errChannel <- nil
}
