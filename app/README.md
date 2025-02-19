# Memorization Tracker API

This is a RESTful API for tracking Quran memorization progress, implemented using the GIN framework and GORM for database handling.

## Getting Started

### Prerequisites

- Go 1.19 or later
- PostgreSQL database
- GIN framework
- GORM library

### Installation

1. Clone the repository:

   ```bash
   git clone ...
   cd your-repository
   docker compose -f ./docker/database/docker-compose.yaml up -d
   cd app
   ```

2. Install the required Go modules:

   ```bash
   go mod tidy
   ```

### Running the Server

To run the server, use the following command:

```bash
go build && ./go-auth
```

This will start the API server at http://localhost:3030.
