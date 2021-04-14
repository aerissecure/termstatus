package termstatus

import (
	"golang.org/x/sys/unix"
)

// IsProcessBackground reports whether the current process is running in the
// background. fd must be a file descriptor for the terminal.
func IsProcessBackground(fd uintptr) bool {
	bg, err := isProcessBackground(fd)
	if err != nil {
		return false
	}
	return bg
}

func isProcessBackground(fd uintptr) (bool, error) {
	pid, err := unix.IoctlGetInt(int(fd), unix.TIOCGPGRP)
	return pid != unix.Getpgrp(), err
}
