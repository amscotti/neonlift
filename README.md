# Neon Lift - Stand and Sit Timer

<p align="center">
    <img src="images/wizard_standing.jpg" alt="A wizard standing at a standing desk, using a computer" width="600">
</p>

Neon Lift is a simple command-line application to remind you to alternate between standing and sitting positions while working at a standing desk. In addition to its functionality, this was created to help learn more about [Bubble Tea](https://github.com/charmbracelet/bubbletea) and [Lip Gloss](https://github.com/charmbracelet/lipgloss).

[![asciicast](https://asciinema.org/a/VQ065W0QL9xH2bBDE6I55RN2b.svg)](https://asciinema.org/a/VQ065W0QL9xH2bBDE6I55RN2b)

## Features

- Customizable standing and sitting durations.
- Visual progress bar.
- Audible alert when it's time to change positions.
- Pause and resume functionality.

## Usage

### Running with Go

To run the application, use the following command:

```sh
go run main.go
```

You can customize the application using the following command-line flags:

```sh
go run main.go -stand=45m -sit=15m -sound=true -desktop=true -icon=/path/to/icon.png
```

Available options:
- `-stand`: Duration for standing (default: 30m)
- `-sit`: Duration for sitting (default: 60m)
- `-sound`: Enable sound notifications (default: true)
- `-desktop`: Enable desktop notifications (default: true)
- `-icon`: Path to notification icon for desktop notifications (optional)

Press 'Enter' to start, 'Space' to pause, and 'Q' to quit the application.


### Running with Docker

You can build and run the application using Docker. First, build the Docker image:

```sh
docker build -t neonlift .
```

Then, run the application in a Docker container:

```sh
docker run --rm neonlift
```

You can also pass command-line flags to customize the application:

```sh
docker run --rm neonlift -stand=45m -sit=15m -sound=true -desktop=false
```

## Developer Information

### Running Tests

To run all tests in the project:

```sh
go test ./...
```

To run tests for a specific package:

```sh
go test ./model
go test ./notification
go test ./timer
```

To run tests with coverage report:

```sh
go test -cover ./...
```

For a detailed HTML coverage report:

```sh
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```