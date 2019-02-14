package clip

import (
	"fmt"
	"io"
	"testing"
	"time"
)

func TestCopyPaste(t *testing.T) {
	c, err := New()
	if err != nil {
		t.Fatal(err)
	}

	// hopefully the user doesn't have more than 16 mb, because
	// we're going to use this to fix their clipboard after the test
	// I generally wouldn't allocate this much in a test, but people
	// would be pretty surprised if a Go test ended up clobbering
	// what they have on the clipboard
	var tmp [1024 * 1024 * 64]byte

	n, err := c.Read(tmp[:])
	if err != nil && err != io.EOF {
		t.Fatal("failed to read clipboard for backup", err)
	}
	defer func() {
		n2, err := c.Write(tmp[:n])
		if err != nil || n2 < n {
			println(n2, n)
			println("***couldn't restore your clipboard, dumping to stdout as hex***")
			fmt.Printf("%x\n", tmp[n])
		}
	}()

	now := time.Now().String()

	t.Run("Paste", func(t *testing.T) {
		n, err := fmt.Fprintln(c, now)
		if n < len([]byte(now)) {
			t.Fatalf("short write: have %d, want %d bytes", n, len([]byte(now)))
		}
		if err != nil {
			t.Fatal("error on write", err)
		}
	})
	t.Run("Copy", func(t *testing.T) {
		tmp := make([]byte, len([]byte(now)))
		c.Read(tmp[:])
		have := string(tmp)
		if have != now {
			t.Fatalf("have %q want %q", have, now)
		}
	})

}
