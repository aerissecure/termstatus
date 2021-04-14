package termstatus

import (
	"os"
	"testing"
)

func TestIsProcessBackground(t *testing.T) {
	tty, err := os.Open("/dev/tty")
	if err != nil {
		t.Skipf("can't open terminal: %v", err)
	}

	_, err = isProcessBackground(tty.Fd())
	if err != nil {
		t.Fatal(err)
	}

	_ = tty.Close()
}
