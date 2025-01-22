# gonsole

A minimal ANSI console library for Go.

## Example Use

```golang
package main

import (
	"fmt"
	"os"
	"strings"
	
	"github.com/mwhickson/gonsole"
)

func main() {

	console := gonsole.NewConsoleHandler()
	// console.DisplayTest() // DEBUG:
	
	console.ClearScreen()
	console.CursorToHome()

	for {
		console.SetForegroundYellow()
		fmt.Print("> ")
		console.SetForegroundWhite()

		line := console.ReadLine()

		if 0 == (strings.Compare(line, "")) {
			os.Exit(0)
		}

		console.SetForegroundYellow()
		fmt.Print(">>> ")
		console.SetForegroundCyan()
		fmt.Println(line)
		console.SetForegroundWhite()
	}

}
```

## Caveats

This is a "use at your own risk" library.

It's not guaranteed to do anything other than take up space on disk, but it will hopefully have its' uses.

*NOTE: The name is horrible, and likely not unique, but...*

## License

[MIT](LICENSE)