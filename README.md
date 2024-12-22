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

## API Credentials

### Hotelbeds API
- **API Key**: `db11033c50b5ed53ab7b815cb1b2eaee`
- **Secret**: `704773c03`

## Documentation
- Hotelbeds API Documentation: [Hotels Booking API](https://developer.hotelbeds.com/documentation/hotels/booking-api/)

## Setup and Testing
1. Clone the repository
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Run the application:
   ```bash
   go run main.go
   ```

## Project Timeline
- Development time: 12 hours
- Submission: GitHub repository with complete source code and documentation

## Repository Structure
```
.
├── README.md
├── go.mod
├── main.go
└── [other project files]
```

## Testing
- A Postman collection is provided for testing the API endpoints
- The collection includes example requests and expected response formats

## Guidelines for Contribution
1. Ensure code is well-documented
2. Follow Go best practices
3. Include appropriate error handling
4. Write clean, maintainable code
5. Add necessary comments for complex logic