package app

import (
	"github.com/amscotti/neonlift/model"
	"github.com/amscotti/neonlift/notification"
	"github.com/amscotti/neonlift/ui"
	tea "github.com/charmbracelet/bubbletea"
)

// TimeExpiredMsg is sent when the timer expires
type TimeExpiredMsg struct{}

// App represents the main application
type App struct {
	model      model.Model
	view       *ui.View
	notifier   notification.Notifier
	lastState  model.State
	lastTimer  int
}

// NewApp creates a new app instance
func NewApp(m model.Model) *App {
	// Create a combo notifier with both sound and desktop notifications
	soundNotifier := notification.DefaultSoundNotifier()
	desktopNotifier := notification.NewDesktopNotifier("") // Empty string for default icon
	comboNotifier := notification.NewComboNotifier(soundNotifier, desktopNotifier)
	
	return &App{
		model:     m,
		view:      ui.NewView(),
		notifier:  comboNotifier,
		lastState: m.State,
		lastTimer: int(m.Timer.Seconds()),
	}
}

// SetNotifier changes the notification method
func (a *App) SetNotifier(n notification.Notifier) {
	a.notifier = n
}

// Init initializes the app
func (a App) Init() tea.Cmd {
	return a.model.Init()
}

// Update handles messages and updates the app state
func (a App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	updatedModel, cmd := a.model.Update(msg)
	
	// Type assertion to convert tea.Model back to model.Model
	a.model = updatedModel.(model.Model)

	// Check if timer has expired
	if a.lastState != model.Waiting && a.model.State == model.Waiting {
		// Timer just expired, send notification
		_ = a.notifier.Notify("Neon Lift", "Time to change position!")
	}

	a.lastState = a.model.State
	a.lastTimer = int(a.model.Timer.Seconds())

	return a, cmd
}

// View renders the current app state
func (a App) View() string {
	return a.view.RenderModel(a.model)
}