//go:build leak
// +build leak

package deadlock

import (
	"testing"
	"time"
)

// go test -tags 'leak' -run TestRaceDetectorConflict -race -timeout 10h
func TestRaceDetectorConflict(t *testing.T) {
	defer restore()()
	Opts.DeadlockTimeout = 30 * time.Second
	var a Mutex
	for {
		a.Lock()
		a.Unlock()
	}
}
