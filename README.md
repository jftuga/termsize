# termsize
Go module to return the terminal width and height

The [Releases Page](https://github.com/jftuga/termsize/releases) contains command-line binaries for Windows, MacOS and Linux.

## NOTES
When a terminal is not available such as when using a pipe or redirecting to a file, then the `DefaultWidth` and 
`DefaultHeight` as defined in `defaults.go` will be used.  These defaults can be changed by setting 
`termsize.DefaultWidth` and `termsize.DefaultHeight`.

## INSTALL

Go version 1.16 or newer is required for compilation.

```
go get github.com/jftuga/termsize
```

## ACKNOWLEDGEMENTS
This code is adapted from https://github.com/jessevdk/go-flags.
* [BSD-3-Clause License](https://github.com/jftuga/termsize/blob/main/LICENSE)

## Example

```go
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
```

## Output

```
155x46
```
