package util

import (
	"os"
)

// StdinIsPipe returns true if stdin is receiving piped input
func StdinIsPipe() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode()&os.ModeCharDevice) == 0
}
