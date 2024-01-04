package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/lipgloss"

	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gen2brain/beeep"
)

var (
	standingDuration = flag.Duration("stand", 30*time.Minute, "Duration for standing")
	sittingDuration  = flag.Duration("sit", 1*time.Hour, "Duration for sitting")
)

// formatTime takes a duration and returns it as a formatted string.
func formatTime(d time.Duration) string {
	minutes := d / time.Minute
	seconds := (d - minutes*time.Minute) / time.Second
	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}

type State int

const (
	Sitting State = iota
	Standing
	Waiting
	Start
)

// tickMsg is sent on every tick
type tickMsg time.Time

type model struct {
	state         State
	previousState State
	cycleCount    uint8
	timer         time.Duration
	paused        bool
	progress      progress.Model
}

func initialModel() model {
	return model{
		state:         Start,
		previousState: Sitting,
		cycleCount:    0,
		timer:         *standingDuration,
		progress:      progress.New(progress.WithWidth(60), progress.WithDefaultGradient(), progress.WithoutPercentage()),
	}
}

func (m model) Init() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg { return tickMsg(t) })
}

const neonPink = lipgloss.Color("#FF00FF")
const neonBlue = lipgloss.Color("#00FFFF")
const darkBackground = lipgloss.Color("#120052")

var appStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(neonPink).
	Background(darkBackground).
	Width(80) // Set a minimum width for the app window

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Underline(true).
			Foreground(neonPink).
			Background(darkBackground).
			Width(80).
			Padding(2, 0).
			Align(lipgloss.Center)

	countStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(neonBlue).
			Background(darkBackground).
			Width(80).
			Padding(0, 0).
			Align(lipgloss.Center)

	timerStyle = lipgloss.NewStyle().
			Foreground(neonBlue).
			Background(darkBackground).
			Width(80).
			Align(lipgloss.Center).
			Padding(1, 0)

	progressStyle = lipgloss.NewStyle().
			Foreground(neonBlue).
			Background(darkBackground).
			Width(80).
			Align(lipgloss.Center).
			Padding(1, 0)

	instructionStyle = lipgloss.NewStyle().
				Foreground(neonPink).
				Width(80).
				Align(lipgloss.Center)
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			return m, tea.Batch(tea.Quit)
		case " ":
			if m.state == Standing || m.state == Sitting {
				if m.paused {
					m.paused = false
					return m, tea.Tick(time.Second, func(t time.Time) tea.Msg { return tickMsg(t) })
				} else {
					m.paused = true
					return m, nil
				}
			}
			return m, nil

		case "enter":
			if m.state == Waiting || m.state == Start {
				// Transition to the opposite of the previous state
				if m.previousState == Sitting {
					m.state = Standing
					m.timer = *standingDuration // Use the standing duration from the command-line option
				} else {
					m.state = Sitting
					m.timer = *sittingDuration // Use the sitting duration from the command-line option
				}
				return m, tea.Tick(time.Second, func(t time.Time) tea.Msg { return tickMsg(t) })
			}
		}
	case tickMsg:
		if (m.state == Standing || m.state == Sitting) && m.timer > 0 {
			if !m.paused {
				m.timer -= time.Second
				initialDuration := *standingDuration
				if m.state == Sitting {
					initialDuration = *sittingDuration
				}
				elapsedTime := initialDuration - m.timer
				progressPercent := float64(elapsedTime) / float64(initialDuration)
				return m, tea.Batch(tea.Tick(time.Second, func(t time.Time) tea.Msg { return tickMsg(t) }), m.progress.SetPercent(progressPercent))
			}
		} else if m.timer <= 0 {
			err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
			if err != nil {
				panic(err)
			}

			m.previousState = m.state
			m.state = Waiting
			m.cycleCount++
			return m, m.progress.SetPercent(0.0)
		}
	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd
	}

	return m, nil
}

func (m model) View() string {
	var fullView string
	title := titleStyle.Render("Neon Lift - Escape the Pixel Slump")
	instructions := instructionStyle.Render("Press 'Enter' to start, 'Space' to pause, 'Q' to quit")
	progress := progressStyle.Render(m.progress.View())

	count := ""
	for i := uint8(0); i < m.cycleCount; i++ {
		if i%2 == 0 {
			count += "● "
		} else {
			count += "○ "
		}
	}
	count = strings.TrimSpace(count)

	// Content based on state
	switch m.state {
	case Sitting:
		timer := timerStyle.Render(fmt.Sprintf("Sitting down! Time left: %s", formatTime(m.timer)))
		fullView = lipgloss.JoinVertical(lipgloss.Left, title, countStyle.Render(count), timer, progress, instructions)
	case Standing:
		timer := timerStyle.Render(fmt.Sprintf("Standing up! Time left: %s", formatTime(m.timer)))
		fullView = lipgloss.JoinVertical(lipgloss.Left, title, countStyle.Render(count), timer, progress, instructions)
	case Waiting:
		waiting := timerStyle.Render("Please change your position")
		fullView = lipgloss.JoinVertical(lipgloss.Left, title, countStyle.Render(count), waiting, progressStyle.Render(), instructions)
	case Start:
		waiting := timerStyle.Render("Welcome, please begin standing")
		fullView = lipgloss.JoinVertical(lipgloss.Left, title, countStyle.Render(count), waiting, progressStyle.Render(), instructions)
	default:
		unknown := timerStyle.Render("Unknown state")
		fullView = lipgloss.JoinVertical(lipgloss.Left, title, countStyle.Render(count), unknown, progressStyle.Render(), instructions)
	}

	// Apply app-wide styling
	return appStyle.Render(fullView)
}

func main() {
	flag.Parse()

	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
