package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spiegel-im-spiegel/emojis/unames/json/generator/namelist"
)

func main() {
	list, err := namelist.Parse()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(list); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
