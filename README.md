# PulsePlus

PulsePlus is a lightweight, concurrent URL status checker written in Go. It efficiently pings a list of websites to verify their availability and reports their HTTP status codes.

## Features

- **Concurrent Checks**: Uses Go routines to check multiple URLs simultaneously for faster execution.
- **Synchronization**: Utilizes `sync.WaitGroup` to ensure all checks complete before the program exits.
- **Error Handling**: Includes a helper utility to gracefully handle and log errors with timestamps.

## Project Structure

```
pulseplus/
├── main.go           # Entry point: defines URLs and launches concurrent checks
├── helpers/
│   └── helpers.go    # Utility functions for error handling
└── go.mod            # Go module definition
```

## Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) 1.25 or higher

### Running the Application

1. Clone the repository:
   ```bash
   git clone https://github.com/realtouseef/pulseplus.git
   cd pulseplus
   ```

2. Run the program:
   ```bash
   go run main.go
   ```

### Example Output

```text
status code for https://github.com is 200
status code for https://www.linkedin.com is 200
```
*(Note: Output order may vary due to concurrency)*
