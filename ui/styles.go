package ui

import (
	"github.com/charmbracelet/lipgloss"
)

// Colors
const (
	NeonPink       = lipgloss.Color("#FF00FF")
	NeonBlue       = lipgloss.Color("#00FFFF")
	DarkBackground = lipgloss.Color("#120052")
)

// Styles holds all the application styles
type Styles struct {
	App          lipgloss.Style
	Title        lipgloss.Style
	Count        lipgloss.Style
	Timer        lipgloss.Style
	Progress     lipgloss.Style
	Instructions lipgloss.Style
}

// DefaultStyles returns the default styling for the application
func DefaultStyles() Styles {
	app := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(NeonPink).
		Background(DarkBackground).
		Width(80)

	title := lipgloss.NewStyle().
		Bold(true).
		Underline(true).
		Foreground(NeonPink).
		Background(DarkBackground).
		Width(80).
		Padding(2, 0).
		Align(lipgloss.Center)

	count := lipgloss.NewStyle().
		Bold(true).
		Foreground(NeonBlue).
		Background(DarkBackground).
		Width(80).
		Padding(0, 0).
		Align(lipgloss.Center)

	timer := lipgloss.NewStyle().
		Foreground(NeonBlue).
		Background(DarkBackground).
		Width(80).
		Align(lipgloss.Center).
		Padding(1, 0)

	progress := lipgloss.NewStyle().
		Foreground(NeonBlue).
		Background(DarkBackground).
		Width(80).
		Align(lipgloss.Center).
		Padding(1, 0)

	instructions := lipgloss.NewStyle().
		Foreground(NeonPink).
		Width(80).
		Align(lipgloss.Center)

	return Styles{
		App:          app,
		Title:        title,
		Count:        count,
		Timer:        timer,
		Progress:     progress,
		Instructions: instructions,
	}
}