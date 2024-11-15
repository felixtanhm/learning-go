package main

import (
	"assignment/models"
	"assignment/src/constants"
	"assignment/src/print"
	"assignment/src/utils"
	"fmt"
)

func main() {
	venueList := models.CreateBST()
	menuState := make(chan string)

	go func() {
		if err := utils.LoadData(venueList); err != nil {
			fmt.Println("Failed to load data:", err)
			menuState <- "exit" // Exit if data loading fails
		}
	}()
	go func() {
		menuState <- "userMenu"
	}()
	for displayedMenu := range menuState {
		switch displayedMenu {
		case constants.UserMenu:
			go print.UserMenu(menuState)
		case constants.BrowseVenues:
			go print.BrowseVenues(menuState, venueList, 1)
		case constants.SearchVenues:
			go print.SearchVenues(menuState, venueList, "")
		case constants.BookVenue:
			go print.BookVenue(menuState, venueList)
		case constants.AdminMenu:
			go print.AdminMenu(menuState)
		case constants.AddVenue:
			go print.AddVenue(menuState, venueList)
		case constants.GetVenueBookings:
			go print.GetVenueBookings(menuState, venueList, "")
		case constants.Exit:
			fmt.Println("Thanks! See you again! :)")
			close(menuState)
		}
	}
}
