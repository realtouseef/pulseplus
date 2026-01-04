# PulsePlus

PulsePlus is a lightweight URL status checker API built with [Go Fiber](https://gofiber.io/). It allows you to check the availability and HTTP status code of a website via a RESTful endpoint.

## Features

- **REST API**: Simple and fast HTTP API.
- **Real-time Checks**: Pings the requested URL and returns its status code.
- **Error Handling**: Graceful error handling for invalid requests or unreachable URLs.

## Project Structure

```
pulseplus/
├── main.go           # Entry point: sets up Fiber app and routes
├── handlers/
│   └── ping.go       # Handler for the ping endpoint
├── helpers/
│   └── helpers.go    # Utility functions
└── go.mod            # Go module definition
```

## Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) 1.25 or higher

### Running the Server

1. Clone the repository:
   ```bash
   git clone https://github.com/realtouseef/pulseplus.git
   cd pulseplus
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Start the server:
   ```bash
   go run main.go
   ```

   The server will start listening on port `3000`.

## API Documentation

### Ping a URL

Checks the status code of a provided URL.

- **Endpoint**: `POST /api/v1/ping`
- **Content-Type**: `application/json`

#### Request Body

```json
{
  "url": "https://github.com"
}
```

#### Example Request

```bash
curl -X POST http://localhost:3000/api/v1/ping \
  -H "Content-Type: application/json" \
  -d '{"url": "https://github.com"}'
```

#### Example Response

```text
status code for https://github.com is 200
```
