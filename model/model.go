package model

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

// State represents the current state of the application
type State int

const (
	Sitting State = iota
	Standing
	Waiting
	Start
)

// TickMsg is sent on every tick
type TickMsg time.Time

// Model represents the application state
type Model struct {
	State         State
	PreviousState State
	CycleCount    uint8
	Timer         time.Duration
	Paused        bool
	Progress      progress.Model
	StandDuration time.Duration
	SitDuration   time.Duration
}

// NewModel initializes a new model with the given durations
func NewModel(standDuration, sitDuration time.Duration) Model {
	return Model{
		State:         Start,
		PreviousState: Sitting,
		CycleCount:    0,
		Timer:         standDuration,
		Progress:      progress.New(progress.WithWidth(60), progress.WithDefaultGradient(), progress.WithoutPercentage()),
		StandDuration: standDuration,
		SitDuration:   sitDuration,
	}
}

// Init initializes the model and returns an initial command
func (m Model) Init() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg { return TickMsg(t) })
}

// Update handles messages and updates the model
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			return m, tea.Batch(tea.Quit)
		case " ":
			if m.State == Standing || m.State == Sitting {
				if m.Paused {
					m.Paused = false
					return m, tea.Tick(time.Second, func(t time.Time) tea.Msg { return TickMsg(t) })
				} else {
					m.Paused = true
					return m, nil
				}
			}
			return m, nil

		case "enter":
			if m.State == Waiting || m.State == Start {
				// Transition to the opposite of the previous state
				if m.PreviousState == Sitting {
					m.State = Standing
					m.Timer = m.StandDuration
				} else {
					m.State = Sitting
					m.Timer = m.SitDuration
				}
				return m, tea.Tick(time.Second, func(t time.Time) tea.Msg { return TickMsg(t) })
			}
		}
	case TickMsg:
		if (m.State == Standing || m.State == Sitting) && m.Timer > 0 {
			if !m.Paused {
				m.Timer -= time.Second
				initialDuration := m.StandDuration
				if m.State == Sitting {
					initialDuration = m.SitDuration
				}
				elapsedTime := initialDuration - m.Timer
				progressPercent := float64(elapsedTime) / float64(initialDuration)
				return m, tea.Batch(tea.Tick(time.Second, func(t time.Time) tea.Msg { return TickMsg(t) }), m.Progress.SetPercent(progressPercent))
			}
		} else if m.Timer <= 0 {
			// Notification will be handled externally

			m.PreviousState = m.State
			m.State = Waiting
			m.CycleCount++
			return m, m.Progress.SetPercent(0.0)
		}
	case progress.FrameMsg:
		progressModel, cmd := m.Progress.Update(msg)
		m.Progress = progressModel.(progress.Model)
		return m, cmd
	}

	return m, nil
}

// FormatTime takes a duration and returns it as a formatted string.
func FormatTime(d time.Duration) string {
	minutes := d / time.Minute
	seconds := (d - minutes*time.Minute) / time.Second
	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}

// View is a placeholder to satisfy the tea.Model interface
// The actual view implementation is handled in ui package
func (m Model) View() string {
	return ""
}