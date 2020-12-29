// Tested on Windows 10, Raspberry Pi OS Linux, and MacOS Big Sur 11.1

package main

import (
	"fmt"

	"github.com/jftuga/termsize"
)

func main() {
	w := termsize.GetTerminalColumns()
	h := termsize.GetTerminalRows()
	fmt.Printf("%dx%d\n", w, h)
}
