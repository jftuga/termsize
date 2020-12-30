// +build plan9 appengine wasm

// Adopted from https://github.com/jessevdk/go-flags

package termsize

func Width() int {
	return DefaultHeight
}

func Height() int {
	return DefaultHeight
}
