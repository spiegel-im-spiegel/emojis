// +build run

package main

import (
	"fmt"
	"os"

	"github.com/spiegel-im-spiegel/emojis/unames"
)

func main() {
	nm, err := unames.New()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	var r rune = 0x10001
	for ; r < 0x1ffff; r++ {
		name := nm.Name(r)
		if len(name) == 0 {
			break
		}
		fmt.Printf("%U: %s\n", r, name)
	}
}
