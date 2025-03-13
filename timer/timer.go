package timer

import (
	"fmt"
	"time"
	
	tea "github.com/charmbracelet/bubbletea"
)

// TickMsg is sent on every tick
type TickMsg time.Time

// TimeExpiredMsg is sent when the timer expires
type TimeExpiredMsg struct{}

// Timer represents the core timer functionality
type Timer struct {
	Duration      time.Duration
	Remaining     time.Duration
	Paused        bool
	InitialPaused bool
}

// NewTimer creates a new timer with the given duration
func NewTimer(duration time.Duration, initiallyPaused bool) *Timer {
	return &Timer{
		Duration:      duration,
		Remaining:     duration,
		Paused:        initiallyPaused,
		InitialPaused: initiallyPaused,
	}
}

// Start returns a command to start the timer
func (t *Timer) Start() tea.Cmd {
	t.Paused = t.InitialPaused
	return tea.Tick(time.Second, func(time time.Time) tea.Msg {
		return TickMsg(time)
	})
}

// Reset resets the timer to its initial duration
func (t *Timer) Reset() {
	t.Remaining = t.Duration
	t.Paused = t.InitialPaused
}

// Pause pauses the timer
func (t *Timer) Pause() {
	t.Paused = true
}

// Resume resumes the timer if it was paused
func (t *Timer) Resume() tea.Cmd {
	if t.Paused {
		t.Paused = false
		return tea.Tick(time.Second, func(time time.Time) tea.Msg {
			return TickMsg(time)
		})
	}
	return nil
}

// Update updates the timer when a tick occurs
// Returns true if the timer expired, false otherwise
func (t *Timer) Update() (bool, tea.Cmd) {
	if !t.Paused && t.Remaining > 0 {
		t.Remaining -= time.Second
		
		if t.Remaining <= 0 {
			return true, nil
		}
		
		return false, tea.Tick(time.Second, func(time time.Time) tea.Msg {
			return TickMsg(time)
		})
	}
	
	return false, nil
}

// FormatTime formats the remaining time as MM:SS
func (t *Timer) FormatTime() string {
	minutes := t.Remaining / time.Minute
	seconds := (t.Remaining - minutes*time.Minute) / time.Second
	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}

// Progress returns the progress as a float between 0 and 1
func (t *Timer) Progress() float64 {
	if t.Duration <= 0 {
		return 0.0
	}
	
	elapsed := t.Duration - t.Remaining
	return float64(elapsed) / float64(t.Duration)
}

// IsPaused returns whether the timer is paused
func (t *Timer) IsPaused() bool {
	return t.Paused
}

// IsExpired returns whether the timer has expired
func (t *Timer) IsExpired() bool {
	return t.Remaining <= 0
}