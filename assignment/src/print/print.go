package print

import (
	"assignment/src/bst"
	"assignment/src/venues"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func UserMenu(menuState chan string) {
	fmt.Println("")
	fmt.Println("********************************************")
	fmt.Println("*                                          *")
	fmt.Println("*  Welcome to Felix's Venue Booking App    *")
	fmt.Println("*                                          *")
	fmt.Println("*  1. Browse Venues                        *")
	fmt.Println("*  2. Book a Venue                         *")
	fmt.Println("*  3. Admin Menu                           *")
	fmt.Println("*  4. Close App                            *")
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
		// Code to execute if variable == value1
		menuState <- "browseVenues"
	case 2:
		// Code to execute if variable == value2
	case 3:
		menuState <- "adminMenu"
	case 4:
		menuState <- "exit"
	default:
		fmt.Println("Invalid option, please select from the provided options.")
		menuState <- "userMenu"
	}
}

func BrowseVenues(menuState chan string, venuesP *bst.BST) {
	venuesList := venuesP.GetAll("inOrder")
	for _, venue := range venuesList {
		fmt.Printf("%v | Type: %v | Capacity: %v\n", venue.Name, venue.Type, venue.Capacity)
	}
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Type \"Back\" to go back to menu.")
	command, _ := reader.ReadString('\n')
	command = strings.TrimSpace(command)
	if command == "Back" {
		menuState <- "userMenu"
	} else {
		fmt.Println("Invalid Input.")
		menuState <- "browseVenues"
	}
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
		// Code to execute if variable == value1
		menuState <- "addVenue"
	case 2:
		// Code to execute if variable == value2
	case 3:
		menuState <- "userMenu"
	default:
		fmt.Println("Invalid option, please select from the provided options.")
		menuState <- "adminMenu"
	}
}

func AddVenue(menuState chan string, venuesP *bst.BST) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your venue name:")
	venueName, _ := reader.ReadString('\n')
	venueName = strings.TrimSpace(venueName)
	fmt.Println("Enter your venue type (Room or Hall):")
	venueType, _ := reader.ReadString('\n')
	venueType = strings.TrimSpace(venueType)
	fmt.Println("Enter your venue capacity:")
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
	if err, ok := <-channel; !ok && err != nil {
		fmt.Printf("error adding venue: %v", err)
		menuState <- "adminMenu"
		return
	} else {
		fmt.Printf("Your new venue, %v, has been added to the list of venues!", venueName)
	}
	menuState <- "userMenu"
}
