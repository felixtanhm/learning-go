# Venue Booking System

## Overview
The Venue Booking System is a command-line application designed to facilitate venue management and bookings. It offers features for both users and administrators, with persistent data storage and efficient data handling.

---

## Features
1. **Browse Venues**: View all available venues with pagination support.
2. **Search Venues**: Search for venues by name.
3. **Book Venue**: Book a specific venue by providing your email address.
4. **Admin Menu**:
   - Add new venues.
   - View bookings for a specific venue.
   - Manage the venue database.
5. **Persistent Data**: All venue and booking information is stored persistently in CSV files.

---

## Data Structures
### Binary Search Tree (BST)
- **Purpose**: Stores and organizes venues by their names.
- **Advantages**: Efficient search, insertion, and traversal operations (average-case complexity: O(log n)).
- **Implementation**: Built from scratch to meet assignment requirements.

### Linked List
- **Purpose**: Each venue maintains a linked list of its bookings.
- **Advantages**: Dynamic memory allocation and efficient management of venue-specific bookings.
- **Implementation**: Pointer-based, built without relying on Go's built-in data structures.

---

## Data Files
The application uses two CSV files for data persistence:

### `venueDB.csv`
- **Format**: `ID, Name, Type, Capacity`
- **Example**:
`123e4567-e89b-12d3-a456-426614174000, Main Hall, hall, 200`


### `bookingDB.csv`
- **Format**: `ID, VenueID, UserEmail`
- **Example**:
`123e4567-e89b-12d3-a456-426614174001, 123e4567-e89b-12d3-a456-426614174000, user@example.com`


---

## Error Handling and Concurrency
### Error Handling
- **Input Validation**:
- Ensures all inputs (e.g., email addresses, venue names, capacities) are valid.
- **File Operations**:
- Captures errors during file reading and writing with meaningful error messages.
- **Custom Errors**:
- Detects and reports invalid operations (e.g., duplicate venue names or invalid bookings).

### Concurrency
- **Goroutines**:
- Used for menu navigation and background tasks.
- **Channels**:
- Enables communication between goroutines for asynchronous operations.
- **Mutex**:
- Ensures thread-safe file writes to prevent data corruption.

---

## Instructions to Run the Application
1. **Prerequisites**:
 - Install [Go](https://golang.org/dl/) on your system.
 - Ensure `venueDB.csv` is in the `data/` directory.

2. **Steps**:
 - Clone or extract the project files to a local directory.
 - Navigate to the project directory in your terminal.
 - Run the application using the following command:
   ```
   go run main.go
   ```
 - Follow the prompts to browse venues, book venues, or manage data.
