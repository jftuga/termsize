// +build !windows,!plan9,!appengine,!wasm

// Adopted from https://github.com/jessevdk/go-flags

package termsize

import (
	"golang.org/x/sys/unix"
)

func GetTerminalColumns() int {
	ws, err := unix.IoctlGetWinsize(0, unix.TIOCGWINSZ)
	if err != nil {
		return 80
	}
	return int(ws.Col)
}

func GetTerminalRows() int {
	ws, err := unix.IoctlGetWinsize(0, unix.TIOCGWINSZ)
	if err != nil {
		return 80
	}
	return int(ws.Row)
}
