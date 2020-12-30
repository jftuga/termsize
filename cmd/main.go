// Tested on Windows 10, Raspberry Pi OS Linux, and MacOS Big Sur 11.1

package main

import (
	"fmt"

	"github.com/jftuga/termsize"
)

func main() {
	termsize.DefaultHeight = 133
	w := termsize.Width()
	h := termsize.Height()
	fmt.Printf("%dx%d\n", w, h)
}
