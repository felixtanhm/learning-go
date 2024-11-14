package main

import (
	"assignment/models"
	"assignment/src/print"
	"assignment/src/utils"
	"fmt"
)

func main() {
	venueList := models.CreateBST()
	menuState := make(chan string)
	newVenueState := make(chan bool)

	go utils.ReadVenues(venueList)
	go func() {
		menuState <- "userMenu"
	}()
	go func() {
		if update, ok := <-newVenueState; update && ok {
			go utils.ReadVenues(venueList)
		}
	}()
	for displayedMenu := range menuState {
		switch displayedMenu {
		case "userMenu":
			go print.UserMenu(menuState)
		case "browseVenues":
			go print.BrowseVenues(menuState, venueList, 1)
		case "searchVenues":
			go print.SearchVenues(menuState, venueList, "")
		// case "bookVenue":
		// 	go print.BookVenue()
		case "adminMenu":
			go print.AdminMenu(menuState)
		case "addVenue":
			go print.AddVenue(menuState, venueList)
		case "exit":
			fmt.Println("Thanks! See you again! :)")
			close(menuState)
		}
	}
}
