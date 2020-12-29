// +build plan9 appengine wasm

// Adopted from https://github.com/jessevdk/go-flags

package termsize

func GetTerminalColumns() int {
	return 80
}

func GetTerminalRows() int {
	return 24
}
