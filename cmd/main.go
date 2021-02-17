// Tested on Windows 10, Raspberry Pi OS Linux, and MacOS Big Sur 11.1

package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"

	"github.com/jftuga/termsize"
)

var (
	Version string = strings.TrimSpace(version)
	//go:embed version.txt
	version string
)

func info() {
	fmt.Printf("termsize v%s\n", Version)
	fmt.Println("https://github.com/jftuga/termsize")
	fmt.Println()
}

func main() {
	if len(os.Args) == 2 && (os.Args[1] == "-h" || os.Args[1] == "-?" || os.Args[1] == "-v") {
		info()
		return
	}

	termsize.DefaultHeight = 120
	w := termsize.Width()
	h := termsize.Height()
	fmt.Printf("%dx%d\n", w, h)
}
