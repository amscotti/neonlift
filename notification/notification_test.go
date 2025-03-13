package notification

import (
	"testing"
)

// MockNotifier implements the Notifier interface for testing
type MockNotifier struct {
	Called      bool
	Title       string
	Message     string
	ShouldError bool
}

func (m *MockNotifier) Notify(title, message string) error {
	m.Called = true
	m.Title = title
	m.Message = message
	
	if m.ShouldError {
		return &mockError{"Mock notification error"}
	}
	return nil
}

type mockError struct {
	message string
}

func (e *mockError) Error() string {
	return e.message
}

func TestDefaultSoundNotifier(t *testing.T) {
	notifier := DefaultSoundNotifier()
	
	if !notifier.UseDefault {
		t.Errorf("Expected UseDefault to be true, got false")
	}
}

func TestNewSoundNotifier(t *testing.T) {
	freq := 440.0
	duration := 500
	
	notifier := NewSoundNotifier(freq, duration)
	
	if notifier.UseDefault {
		t.Errorf("Expected UseDefault to be false, got true")
	}
	
	if notifier.Frequency != freq {
		t.Errorf("Expected Frequency to be %f, got %f", freq, notifier.Frequency)
	}
	
	if notifier.Duration != duration {
		t.Errorf("Expected Duration to be %d, got %d", duration, notifier.Duration)
	}
}