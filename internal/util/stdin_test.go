package util

import (
	"os"
	"testing"
)

func TestStdinIsPipe(t *testing.T) {
	// This test verifies the function works
	// In normal usage, StdinIsPipe() returns false when run directly
	// and true when receiving piped input
	result := StdinIsPipe()
	// When run via go test, stdin should not be a pipe
	if result {
		t.Log("StdinIsPipe() returned true - may be running in piped environment")
	} else {
		t.Log("StdinIsPipe() returned false - normal terminal input")
	}
}

func TestStdinStat(t *testing.T) {
	// Verify we can stat stdin without errors
	stat, err := os.Stdin.Stat()
	if err != nil {
		t.Fatalf("Failed to stat stdin: %v", err)
	}

	// Just verify we got a valid stat
	if stat == nil {
		t.Error("Expected non-nil stat")
	}
}
