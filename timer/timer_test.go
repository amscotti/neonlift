package timer

import (
	"testing"
	"time"
)

func TestNewTimer(t *testing.T) {
	duration := 5 * time.Minute
	timer := NewTimer(duration, false)

	if timer.Duration != duration {
		t.Errorf("Expected duration to be %v, got %v", duration, timer.Duration)
	}

	if timer.Remaining != duration {
		t.Errorf("Expected remaining time to be %v, got %v", duration, timer.Remaining)
	}

	if timer.Paused {
		t.Errorf("Expected paused to be false, got true")
	}
}

func TestReset(t *testing.T) {
	duration := 5 * time.Minute
	timer := NewTimer(duration, false)
	
	// Simulate time passing
	timer.Remaining = 2 * time.Minute
	timer.Paused = true
	
	// Reset the timer
	timer.Reset()
	
	if timer.Remaining != duration {
		t.Errorf("Expected remaining time to be %v after reset, got %v", duration, timer.Remaining)
	}
	
	if timer.Paused {
		t.Errorf("Expected paused to be false after reset, got true")
	}
}

func TestPauseResume(t *testing.T) {
	timer := NewTimer(5*time.Minute, false)
	
	// Test pause
	timer.Pause()
	if !timer.Paused {
		t.Errorf("Expected timer to be paused after Pause(), got not paused")
	}
	
	// Test resume
	cmd := timer.Resume()
	if timer.Paused {
		t.Errorf("Expected timer to not be paused after Resume(), got paused")
	}
	if cmd == nil {
		t.Errorf("Expected Resume() to return a command, got nil")
	}
	
	// Test resume when already running
	cmd = timer.Resume()
	if cmd != nil {
		t.Errorf("Expected Resume() to return nil when timer already running, got non-nil")
	}
}

func TestFormatTime(t *testing.T) {
	testCases := []struct {
		remaining time.Duration
		expected  string
	}{
		{5 * time.Minute, "05:00"},
		{65 * time.Minute, "65:00"},
		{1*time.Minute + 30*time.Second, "01:30"},
		{90*time.Minute + 45*time.Second, "90:45"},
		{0, "00:00"},
	}
	
	for _, tc := range testCases {
		timer := NewTimer(10*time.Minute, false)
		timer.Remaining = tc.remaining
		
		result := timer.FormatTime()
		if result != tc.expected {
			t.Errorf("FormatTime() with %v remaining = %s; expected %s", tc.remaining, result, tc.expected)
		}
	}
}

func TestProgress(t *testing.T) {
	duration := 100 * time.Second
	timer := NewTimer(duration, false)
	
	testCases := []struct {
		remaining time.Duration
		expected  float64
	}{
		{100 * time.Second, 0.0},     // Just started
		{50 * time.Second, 0.5},      // Halfway
		{0 * time.Second, 1.0},       // Completed
		{75 * time.Second, 0.25},     // 25% complete
	}
	
	for _, tc := range testCases {
		timer.Remaining = tc.remaining
		progress := timer.Progress()
		
		if progress != tc.expected {
			t.Errorf("Progress() with %v remaining = %f; expected %f", tc.remaining, progress, tc.expected)
		}
	}
}