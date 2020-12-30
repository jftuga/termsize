// +build !windows,!plan9,!appengine,!wasm

// Adopted from https://github.com/jessevdk/go-flags

package termsize

import (
	"golang.org/x/sys/unix"
)

func Width() int {
	ws, err := unix.IoctlGetWinsize(0, unix.TIOCGWINSZ)
	if err != nil {
		return DefaultHeight
	}
	return int(ws.Col)
}

func Height() int {
	ws, err := unix.IoctlGetWinsize(0, unix.TIOCGWINSZ)
	if err != nil {
		return DefaultHeight
	}
	return int(ws.Row)
}
