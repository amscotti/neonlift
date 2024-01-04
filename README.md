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

You can specify the duration for standing and sitting using the `-stand` and `-sit` flags, respectively:

```sh
go run main.go -stand=45m -sit=15m
```

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

You can also pass the `-stand` and `-sit` flags to customize the durations:

```sh
docker run --rm neonlift -stand=45m -sit=15m
```

## TODO / Improvements

- Restructure code.
- Add unit test
- Options for notification other than beeping.

I'm always looking to improve Neon Lift, so if you're keen to contribute, have at it! Fork the repo, work your magic, and send a pull request my way. 