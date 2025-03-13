package notification

import (
	"github.com/gen2brain/beeep"
)

// Notifier is an interface for different notification methods
type Notifier interface {
	Notify(title, message string) error
}

// SoundNotifier provides sound notifications
type SoundNotifier struct {
	Frequency  float64
	Duration   int
	UseDefault bool
}

// DefaultSoundNotifier creates a notifier with default settings
func DefaultSoundNotifier() *SoundNotifier {
	return &SoundNotifier{
		UseDefault: true,
	}
}

// NewSoundNotifier creates a custom sound notifier
func NewSoundNotifier(freq float64, duration int) *SoundNotifier {
	return &SoundNotifier{
		Frequency:  freq,
		Duration:   duration,
		UseDefault: false,
	}
}

// Notify plays a beep sound
func (s *SoundNotifier) Notify(title, message string) error {
	if s.UseDefault {
		return beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	}
	return beeep.Beep(s.Frequency, s.Duration)
}

// DesktopNotifier provides desktop notifications
type DesktopNotifier struct {
	Icon string // Path to notification icon
}

// NewDesktopNotifier creates a new desktop notifier
func NewDesktopNotifier(iconPath string) *DesktopNotifier {
	return &DesktopNotifier{
		Icon: iconPath,
	}
}

// Notify sends a desktop notification
func (d *DesktopNotifier) Notify(title, message string) error {
	return beeep.Notify(title, message, d.Icon)
}

// ComboNotifier combines multiple notification methods
type ComboNotifier struct {
	Notifiers []Notifier
}

// NewComboNotifier creates a new combo notifier
func NewComboNotifier(notifiers ...Notifier) *ComboNotifier {
	return &ComboNotifier{
		Notifiers: notifiers,
	}
}

// Notify sends notifications through all configured notifiers
func (c *ComboNotifier) Notify(title, message string) error {
	var lastErr error
	for _, n := range c.Notifiers {
		if err := n.Notify(title, message); err != nil {
			lastErr = err
		}
	}
	return lastErr
}