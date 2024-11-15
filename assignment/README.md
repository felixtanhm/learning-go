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

## Data Structures and Their Application

### 1. **Binary Search Tree (BST)**
- **Application**:
  - The Binary Search Tree (BST) is used to store and manage venue data.
  - Venues are sorted by their names in the BST, which allows efficient operations such as searching for a venue, adding a new venue, or listing venues in sorted order.
  - Traversal methods (e.g., in-order, pre-order, post-order) are used to retrieve venues in different formats, including paginated views.

- **Suitability**:
  - **Efficient Search**: Searching for a venue by name is a frequent operation in the system. The BST provides average-case time complexity of \(O(\log n)\), making these operations fast.
  - **Scalability**: The BST structure can handle large numbers of venues without significant performance degradation.
  - **Custom Implementation**: Since built-in data structures are not allowed, the custom BST implementation meets the requirement and ensures flexibility for operations specific to the application.

---

### 2. **Linked List**
- **Application**:
  - Each venue maintains a linked list to store its bookings.
  - The linked list dynamically stores bookings associated with a venue, allowing operations such as adding a booking, retrieving all bookings, or removing a booking.

- **Suitability**:
  - **Dynamic Storage**: The linked list allows dynamic memory allocation, which is ideal for scenarios where the number of bookings varies for each venue.
  - **Efficiency**: Adding new bookings is efficient, with a time complexity of \(O(1)\) for appending to the tail.
  - **Pointer-Based Implementation**: This satisfies the requirement for a pointer-based data structure while ensuring that each venue's bookings are independently managed.

---

### Why These Data Structures Were Chosen
1. **BST for Venues**:
   - Venues need to be searchable, sortable, and retrievable efficiently.
   - The BST ensures that insertion and retrieval operations are optimized for the given use case, especially when dealing with a large dataset.

2. **Linked List for Bookings**:
   - Bookings are tied to individual venues, making a per-venue linked list a natural fit.
   - The dynamic nature of a linked list makes it ideal for scenarios where the number of bookings is unpredictable.

3. **Separation of Concerns**:
   - The BST handles the global organization of venues, while the linked list manages bookings on a per-venue basis. This division ensures modularity and clear data ownership.

4. **Custom Implementation**:
   - Both data structures are implemented from scratch without relying on Go’s standard library, demonstrating a deep understanding of data structures and their applications.

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
