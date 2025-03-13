package main

import (
	"flag"
	"log"
	"time"

	"github.com/amscotti/neonlift/app"
	"github.com/amscotti/neonlift/model"
	"github.com/amscotti/neonlift/notification"
	tea "github.com/charmbracelet/bubbletea"
)

var (
	standingDuration = flag.Duration("stand", 30*time.Minute, "Duration for standing")
	sittingDuration  = flag.Duration("sit", 60*time.Minute, "Duration for sitting")
	notifySound      = flag.Bool("sound", true, "Enable sound notifications")
	notifyDesktop    = flag.Bool("desktop", true, "Enable desktop notifications")
	notifyIcon       = flag.String("icon", "", "Path to notification icon (for desktop notifications)")
)

func main() {
	flag.Parse()

	// Initialize model with command line parameters
	m := model.NewModel(*standingDuration, *sittingDuration)

	// Create application instance
	application := app.NewApp(m)

	// Set up notifications based on command line flags
	var notifiers []notification.Notifier
	if *notifySound {
		notifiers = append(notifiers, notification.DefaultSoundNotifier())
	}
	if *notifyDesktop {
		notifiers = append(notifiers, notification.NewDesktopNotifier(*notifyIcon))
	}
	if len(notifiers) > 0 {
		application.SetNotifier(notification.NewComboNotifier(notifiers...))
	}

	// Run the application
	p := tea.NewProgram(application)
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
