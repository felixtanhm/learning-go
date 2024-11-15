package print

import (
	"assignment/models"
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
		menuState <- "browseVenues"
	case 2:
		menuState <- "searchVenues"
	case 3:
		menuState <- "bookVenue"
	case 4:
		menuState <- "adminMenu"
	case 5:
		menuState <- "exit"
	default:
		fmt.Println("Invalid option, please select from the provided options.")
		menuState <- "userMenu"
	}
}

func BrowseVenues(menuState chan string, venuesP *models.BST, page int) {
	if page < 1 {
		fmt.Println("Invalid Page.")
		menuState <- "browseVenues"
	}
	limit := 5.0
	maxPages := int(math.Ceil(venuesP.Count() / limit))
	venuesList := venuesP.GetList("inOrder", page, int(limit))
	fmt.Println("----------------------------")
	fmt.Println("Results:")
	for _, venue := range venuesList {
		fmt.Printf("%v | Type: %v | Capacity: %v\n", venue.Name, venue.Type, venue.Capacity)
	}

	reader := bufio.NewReader(os.Stdin)

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

	command, _ := reader.ReadString('\n')
	command = strings.TrimSpace(command)
	if command == "Back" {
		menuState <- "userMenu"
	} else if command == "Next" {
		BrowseVenues(menuState, venuesP, page+1)
	} else if command == "Prev" {
		BrowseVenues(menuState, venuesP, page-1)
	} else {
		fmt.Println("Invalid Input.")
		menuState <- "browseVenues"
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

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("----------------------------")
	fmt.Println("> Options:")
	fmt.Println("> Search for a venue by name")
	fmt.Println("> \"Back\" to go back to menu.")

	command, _ := reader.ReadString('\n')
	command = strings.TrimSpace(command)
	if command == "Back" {
		menuState <- "userMenu"
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
		menuState <- "bookVenue"
		return
	} else {
		fmt.Printf("You have made a booking for the %v venue. The admin will reach out to the email provided to confirm your booking.\n", venueName)
	}

	menuState <- "userMenu"
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
		menuState <- "addVenue"
	case 2:
		menuState <- "getVenueBookings"
	case 3:
		menuState <- "userMenu"
	default:
		fmt.Println("Invalid option, please select from the provided options.")
		menuState <- "adminMenu"
	}
}

func AddVenue(menuState chan string, venuesP *models.BST) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("> Enter your venue name:")
	venueName, _ := reader.ReadString('\n')
	venueName = strings.TrimSpace(venueName)

	fmt.Println("> Enter your venue type (Room or Hall):")
	venueType, _ := reader.ReadString('\n')
	venueType = strings.ToLower(strings.TrimSpace(venueType))
	if venueType != "room" && venueType != "hall" {
		fmt.Println("Invalid venue type. Please enter either 'Room' or 'Hall'.")
		menuState <- "addVenue"
		return
	}

	fmt.Println("> Enter your venue capacity:")
	venueCapacityStr, _ := reader.ReadString('\n')
	venueCapacityStr = strings.TrimSpace(venueCapacityStr)
	venueCapacity, err := strconv.Atoi(venueCapacityStr)
	if err != nil {
		fmt.Println("Invalid integer entered.")
		menuState <- "addVenue"
		return
	}

	channel := make(chan error, 1)
	go venues.AddVenue(venueName, venueType, venueCapacity, venuesP, channel)
	if err, ok := <-channel; ok && err != nil {
		fmt.Printf("error adding venue: %v", err)
		menuState <- "addVenue"
		return
	} else {
		fmt.Printf("Your new venue, %v, has been added to the list of venues!", venueName)
	}
	menuState <- "adminMenu"
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

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("----------------------------")
	fmt.Println("> Options:")
	fmt.Println("> Input the venue you want to retrieve the bookings for.")
	fmt.Println("> \"Back\" to go back to menu.")

	command, _ := reader.ReadString('\n')
	command = strings.TrimSpace(command)
	if command == "Back" {
		menuState <- "adminMenu"
	} else {
		GetVenueBookings(menuState, venuesP, command)
	}
}
