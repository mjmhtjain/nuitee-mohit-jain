# Hotel Rates API

## Project Overview
This project implements an API using the Gin framework in Golang that serves as a foundational component of liteAPI services. The API processes requests, interacts with the Hotelbeds API to retrieve rate information, and returns customized responses according to specified formats.

## Technical Requirements

### Core Technologies
- **Language**: Golang
- **Framework**: Gin
- **Integration**: Hotelbeds API

### Key Features
1. **API Integration**
   - Process predefined requests
   - Communicate with Hotelbeds API
   - Fetch rate information

2. **Data Transformation**
   - Transform Hotelbeds API responses to liteAPI format
   - Include both request and response data in final output
   - Maintain data exchange transparency

## Documentation
- Hotelbeds API Documentation: [Hotels Booking API](https://developer.hotelbeds.com/documentation/hotels/booking-api/)

## Setup and Testing
1. Clone the repository

2. Set up environment variables by creating a `.env` file in the root directory:
   ```env
   HOTEL_BEDS_BASE_URL=https://api.test.hotelbeds.com
   HOTELBEDS_API_KEY=your_api_key_here
   HOTELBEDS_API_SECRET=your_api_secret_here
   PORT=8080 
   ```

3. Install dependencies:
   ```bash
   go mod tidy
   ```

4. Run tests:
   ```bash
   go test ./...
   ```

5. Run the application:
   ```bash
   go run main.go
   ```

   The API will be available at `http://localhost:8080` (or your configured PORT)

## Repository Structure
```
.
├── .gitignore             # Git ignore rules
├── README.md              # Project documentation
├── go.mod                 # Go module definition
├── go.sum                 # Go module checksums
├── main.go                # Application entry point
└── cmd/                  # Application source code
    └── internals/        # Internal packages
        ├── handler/      # HTTP request handlers
        ├── service/      # Business logic layer
        ├── dto/          # Data transfer objects
        ├── client/       # External API clients
        └── router/       # Route definitions
```
