package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type venue struct {
	Name         string
	Type         string
	Capacity     int
	Availability bool
}

func main() {
	venueList := []venue{}
	readCSV(&venueList)
	fmt.Println(venueList)
	newVenue := venue{
		Name:         "Hello",
		Type:         "room",
		Capacity:     123,
		Availability: false,
	}
	writeCSV(newVenue)
}

func handlePanic() error {
	if r := recover(); r != nil {
		return fmt.Errorf("Unhandled system error: %v", r)
	}
	return nil
}

func readCSV(venuesP *[]venue) error {
	fileName := "go_fundamentals_assignment.csv"
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("Error opening csv data file: %v", err)
	}
	defer file.Close()
	fmt.Println("hello")

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("Error reading csv data file: %v", err)
	}
	for i, row := range records {
		if i == 0 {
			continue
		}
		capacity, _ := strconv.Atoi(row[2])
		availability, _ := strconv.ParseBool(row[3])
		newVenue := venue{
			Name:         row[0],
			Type:         row[1],
			Capacity:     capacity,
			Availability: availability,
		}
		*venuesP = append(*venuesP, newVenue)
	}
	return nil
}

func writeCSV(newVenue venue) error {
	fileName := "go_fundamentals_assignment.csv"
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("Error opening csv data file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	capacity := strconv.Itoa(newVenue.Capacity)
	availability := strconv.FormatBool(newVenue.Availability)
	record := []string{
		newVenue.Name, newVenue.Type, capacity, availability,
	}

	if err := writer.Write(record); err != nil {
		return fmt.Errorf("Error writing to csv data file: %v", err)
	}
	return nil
}
