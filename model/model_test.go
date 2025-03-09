package model

import (
	"testing"
	"time"
)

func TestNewModel(t *testing.T) {
	standDuration := 30 * time.Minute
	sitDuration := 60 * time.Minute

	model := NewModel(standDuration, sitDuration)

	// Check initial values
	if model.State != Start {
		t.Errorf("Expected initial state to be Start, got %v", model.State)
	}

	if model.PreviousState != Sitting {
		t.Errorf("Expected previous state to be Sitting, got %v", model.PreviousState)
	}

	if model.CycleCount != 0 {
		t.Errorf("Expected cycle count to be 0, got %v", model.CycleCount)
	}

	if model.Timer != standDuration {
		t.Errorf("Expected timer to be %v, got %v", standDuration, model.Timer)
	}

	if model.Paused {
		t.Errorf("Expected paused to be false, got %v", model.Paused)
	}

	if model.StandDuration != standDuration {
		t.Errorf("Expected standing duration to be %v, got %v", standDuration, model.StandDuration)
	}

	if model.SitDuration != sitDuration {
		t.Errorf("Expected sitting duration to be %v, got %v", sitDuration, model.SitDuration)
	}
}

func TestFormatTime(t *testing.T) {
	testCases := []struct {
		input    time.Duration
		expected string
	}{
		{5 * time.Minute, "05:00"},
		{65 * time.Minute, "65:00"},
		{1*time.Minute + 30*time.Second, "01:30"},
		{90*time.Minute + 45*time.Second, "90:45"},
		{0, "00:00"},
	}

	for _, tc := range testCases {
		result := FormatTime(tc.input)
		if result != tc.expected {
			t.Errorf("FormatTime(%v) = %s; expected %s", tc.input, result, tc.expected)
		}
	}
}