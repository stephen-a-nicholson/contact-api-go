# Contact API Go

A RESTful API built with Go for managing a digital contact book. This project provides endpoints for creating, retrieving, updating, and deleting contact information.

## Features

- Create a new contact
- Retrieve a contact by ID
- Update an existing contact
- Delete a contact

## Prerequisites

- Go (version 1.16 or higher)

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/stephen-a-nicholson/contact-api-go.git
   ```

2. Navigate to the project directory:
   ```
   cd contact-api-go
   ```

3. Install the dependencies:
   ```
   go mod download
   ```

## Usage

1. Start the server:
   ```
   go run cmd/server/main.go
   ```

   The server will start running on `http://localhost:8080`.

2. Use an API client (e.g., cURL) to interact with the API endpoints:

   - Create a contact:
     ```
     curl -X POST -H "Content-Type: application/json" -d '{"firstName":"John","lastName":"Doe","email":"john@example.com","phone":"1234567890"}' http://localhost:8080/contacts
     ```

   - Retrieve a contact by ID:
     ```
     curl -X GET http://localhost:8080/contacts/1
     ```

   - Update a contact:
     ```
     curl -X PUT -H "Content-Type: application/json" -d '{"firstName":"John","lastName":"Doe","email":"john.doe@example.com","phone":"9876543210"}' http://localhost:8080/contacts/1
     ```

   - Delete a contact:
     ```
     curl -X DELETE http://localhost:8080/contacts/1
     ```

## Project Structure

```
contact-api-go/
├── cmd/
│   └── server/
│       └── main.go
├── pkg/
│   ├── handlers/
│   │   └── contactHandler.go
│   ├── models/
│   │   └── contact.go
│   └── routes/
│       └── routes.go
├── go.mod
└── go.sum
```

- `cmd/`: Contains the main application entry point.
  - `server/`: Server-related code.
    - `main.go`: The main function that sets up the server and starts it.

- `pkg/`: Contains the package-level code.
  - `handlers/`: Handler functions for each API endpoint.
    - `contactHandler.go`: Defines the handler functions for contact-related operations.
  - `models/`: Data models and structs.
    - `contact.go`: Defines the `Contact` struct representing a contact.
  - `routes/`: Route definitions and route-related code.
    - `routes.go`: Defines the API routes and their corresponding handler functions.

- `go.mod`: Module dependencies and version constraints.

- `go.sum`: Checksums of the dependencies used in the project.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please feel free to open an issue or submit a pull request.

## Licence

This project is licensed under the [MIT Licence](LICENCE).