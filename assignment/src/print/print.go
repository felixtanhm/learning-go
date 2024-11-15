package print

import (
	"assignment/models"
	"assignment/src/constants"
	"assignment/src/utils"
	"assignment/src/venues"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func UserMenu(menuState chan string) {
	fmt.Println("")
	fmt.Println("********************************************************************")
	fmt.Println("*                                                                  *")
	fmt.Println("*  Welcome to Felix's Venue Booking App                            *")
	fmt.Println("*                                                                  *")
	fmt.Println("*  1. Browse Venues                                                *")
	fmt.Println("*  2. Search for a Venue                                           *")
	fmt.Println("*  3. Book a Venue                                                 *")
	fmt.Println("*  4. Admin Menu                                                   *")
	fmt.Println("*  5. Close App                                                    *")
	fmt.Println("*                                                                  *")
	fmt.Println("*  Input menu item's number to navigate                            *")
	fmt.Println("*                                                                  *")
	fmt.Println("*  For the latest version, visit:                                  *")
	fmt.Println("*  https://github.com/felixtanhm/learning-go/tree/main/assignment  *")
	fmt.Println("*                                                                  *")
	fmt.Println("********************************************************************")
	fmt.Println("")
	fmt.Println("> Select an option: ")

	userNav := 0
	fmt.Scan(&userNav)
	switch userNav {
	case 1:
		menuState <- constants.BrowseVenues
	case 2:
		menuState <- constants.SearchVenues
	case 3:
		menuState <- constants.BookVenue
	case 4:
		menuState <- constants.AdminMenu
	case 5:
		menuState <- constants.Exit
	default:
		fmt.Println("Invalid option, please select from the provided options.")
		menuState <- constants.UserMenu
	}
}

func BrowseVenues(menuState chan string, venuesP *models.BST, page int) {
	if page < 1 {
		fmt.Println("Invalid Page.")
		menuState <- constants.BrowseVenues
	}
	limit := 5.0
	maxPages := int(math.Ceil(venuesP.Count() / limit))
	venuesList := venuesP.GetList("inOrder", page, int(limit))
	if len(venuesList) < 1 {
		fmt.Println("No venues available.")
		menuState <- constants.UserMenu
		return
	}
	fmt.Println("----------------------------")
	fmt.Println("Results:")
	for _, venue := range venuesList {
		fmt.Printf("%v | Type: %v | Capacity: %v\n", venue.Name, venue.Type, venue.Capacity)
	}

	fmt.Println("----------------------------")
	fmt.Printf("Viewing Page %d of %d.\n", page, maxPages)
	fmt.Println("----------------------------")
	fmt.Println("> Options:")

	if maxPages > page {
		fmt.Printf("> \"Next\" for next page\n")
	}
	if page > 1 {
		fmt.Printf("> \"Prev\" for previous page\n")
	}
	fmt.Println("> \"Back\" to go back to menu.")
	fmt.Println("----------------------------")

	command := utils.ReadInput()
	switch command {
	case "Back":
		menuState <- constants.UserMenu
	case "Next":
		BrowseVenues(menuState, venuesP, page+1)
	case "Prev":
		BrowseVenues(menuState, venuesP, page-1)
	default:
		fmt.Println("Invalid Input.")
		menuState <- constants.BrowseVenues
	}
}

func SearchVenues(menuState chan string, venuesP *models.BST, searchParam string) {
	if searchParam != "" {
		venue := venuesP.GetOne(searchParam)
		fmt.Println("----------------------------")
		fmt.Println("Result:")
		if venue != nil {
			fmt.Printf("%v | Type: %v | Capacity: %v\n", venue.Name, venue.Type, venue.Capacity)
		} else {
			fmt.Printf("No venues with the name \"%v\" was found.\n", searchParam)
			SearchVenues(menuState, venuesP, "")
		}
	}

	fmt.Println("----------------------------")
	fmt.Println("> Options:")
	fmt.Println("> Search for a venue by name")
	fmt.Println("> \"Back\" to go back to menu.")

	command := utils.ReadInput()
	if command == "Back" {
		menuState <- constants.UserMenu
	} else {
		SearchVenues(menuState, venuesP, command)
	}
}

func BookVenue(menuState chan string, venuesP *models.BST) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("> Enter the venue name that you would like to book:")
	venueName, _ := reader.ReadString('\n')
	venueName = strings.TrimSpace(venueName)

	fmt.Println("Retrieving venue. Please wait...")
	retrievedVenue := venuesP.GetOne(venueName)
	if retrievedVenue == nil {
		fmt.Println("There doesn't seem to be such a venue.")
		menuState <- "bookVenue"
		return
	}
	fmt.Printf("You are making a booking for %v\n", retrievedVenue.Name)

	fmt.Println("> Enter your email:")
	userEmail, _ := reader.ReadString('\n')
	userEmail = strings.TrimSpace(userEmail)

	channel := make(chan error, 1)
	go venues.BookVenue(userEmail, venuesP, retrievedVenue, channel)
	if err, ok := <-channel; ok && err != nil {
		fmt.Printf("error booking venue: %v", err)
		menuState <- constants.BookVenue
		return
	} else {
		fmt.Printf("You have made a booking for the %v venue. The admin will reach out to the email provided to confirm your booking.\n", venueName)
	}

	menuState <- constants.UserMenu
}

func AdminMenu(menuState chan string) {
	fmt.Println("")
	fmt.Println("********************************************")
	fmt.Println("*                                          *")
	fmt.Println("*  [Admin Menu] Venue Booking App          *")
	fmt.Println("*                                          *")
	fmt.Println("*  1. Add Venue                            *")
	fmt.Println("*  2. List Bookings by Venue               *")
	fmt.Println("*  3. Back                                 *")
	fmt.Println("*                                          *")
	fmt.Println("*  Input menu item's number to navigate    *")
	fmt.Println("*                                          *")
	fmt.Println("********************************************")
	fmt.Println("")
	fmt.Println("Select an option: ")

	userNav := 0
	fmt.Scan(&userNav)
	switch userNav {
	case 1:
		menuState <- constants.AddVenue
	case 2:
		menuState <- constants.GetVenueBookings
	case 3:
		menuState <- constants.UserMenu
	default:
		fmt.Println("Invalid option, please select from the provided options.")
		menuState <- constants.AdminMenu
	}
}

func AddVenue(menuState chan string, venuesP *models.BST) {
	fmt.Println("> Enter your venue name:")
	venueName := utils.ReadInput()

	fmt.Println("> Enter your venue type (Room or Hall):")
	venueType := strings.ToLower(utils.ReadInput())
	if venueType != "room" && venueType != "hall" {
		fmt.Println("Invalid venue type. Please enter either 'Room' or 'Hall'.")
		menuState <- constants.AddVenue
		return
	}

	fmt.Println("> Enter your venue capacity:")
	venueCapacityStr := utils.ReadInput()
	venueCapacity, err := strconv.Atoi(venueCapacityStr)
	if err != nil {
		fmt.Println("Invalid integer entered.")
		menuState <- constants.AddVenue
		return
	}

	channel := make(chan error, 1)
	go venues.AddVenue(venueName, venueType, venueCapacity, venuesP, channel)
	if err, ok := <-channel; ok && err != nil {
		fmt.Printf("error adding venue: %v", err)
		menuState <- constants.AddVenue
		return
	} else {
		fmt.Printf("Your new venue, %v, has been added to the list of venues!", venueName)
	}
	menuState <- constants.AdminMenu
}

func GetVenueBookings(menuState chan string, venuesP *models.BST, searchParam string) {
	if searchParam != "" {
		venue := venuesP.GetOne(searchParam)
		if venue != nil {
			fmt.Printf("Retrieving bookings for %v. Please wait...\n", venue.Name)
			bookings := venue.BookingList.GetAll()
			fmt.Println("----------------------------")
			fmt.Println("Result:")
			for i, booking := range bookings {
				fmt.Printf("%v. User: %v\n", i+1, booking.UserEmail)
			}
		} else {
			fmt.Printf("No venues with the name \"%v\" was found.\n", searchParam)
			GetVenueBookings(menuState, venuesP, "")
		}
	}

	fmt.Println("----------------------------")
	fmt.Println("> Options:")
	fmt.Println("> Input the venue you want to retrieve the bookings for.")
	fmt.Println("> \"Back\" to go back to menu.")

	command := utils.ReadInput()
	if command == "Back" {
		menuState <- constants.AdminMenu
	} else {
		GetVenueBookings(menuState, venuesP, command)
	}
}
