package ui

import (
	"fmt"
	"strings"

	"github.com/amscotti/neonlift/model"
	"github.com/charmbracelet/lipgloss"
)

// View encapsulates UI rendering functionality
type View struct {
	Styles Styles
}

// NewView creates a new View with default styles
func NewView() *View {
	return &View{
		Styles: DefaultStyles(),
	}
}

// RenderModel renders the model state into a viewable string
func (v *View) RenderModel(m model.Model) string {
	var fullView string
	title := v.Styles.Title.Render("Neon Lift - Escape the Pixel Slump")
	instructions := v.Styles.Instructions.Render("Press 'Enter' to start, 'Space' to pause, 'Q' to quit")
	progress := v.Styles.Progress.Render(m.Progress.View())

	count := ""
	for i := uint8(0); i < m.CycleCount; i++ {
		if i%2 == 0 {
			count += "● "
		} else {
			count += "○ "
		}
	}
	count = strings.TrimSpace(count)

	// Content based on state
	switch m.State {
	case model.Sitting:
		timer := v.Styles.Timer.Render(fmt.Sprintf("Sitting down! Time left: %s", model.FormatTime(m.Timer)))
		fullView = lipgloss.JoinVertical(lipgloss.Left, title, v.Styles.Count.Render(count), timer, progress, instructions)
	case model.Standing:
		timer := v.Styles.Timer.Render(fmt.Sprintf("Standing up! Time left: %s", model.FormatTime(m.Timer)))
		fullView = lipgloss.JoinVertical(lipgloss.Left, title, v.Styles.Count.Render(count), timer, progress, instructions)
	case model.Waiting:
		waiting := v.Styles.Timer.Render("Please change your position")
		fullView = lipgloss.JoinVertical(lipgloss.Left, title, v.Styles.Count.Render(count), waiting, v.Styles.Progress.Render(), instructions)
	case model.Start:
		waiting := v.Styles.Timer.Render("Welcome, please begin standing")
		fullView = lipgloss.JoinVertical(lipgloss.Left, title, v.Styles.Count.Render(count), waiting, v.Styles.Progress.Render(), instructions)
	default:
		unknown := v.Styles.Timer.Render("Unknown state")
		fullView = lipgloss.JoinVertical(lipgloss.Left, title, v.Styles.Count.Render(count), unknown, v.Styles.Progress.Render(), instructions)
	}

	// Apply app-wide styling
	return v.Styles.App.Render(fullView)
}