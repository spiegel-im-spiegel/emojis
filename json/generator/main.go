package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/spiegel-im-spiegel/emojis/json/generator/data"
)

func outputEmojiData() {
	list, err := data.Parse()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	if err := data.EncodeJSON(os.Stdout, list); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func outputEmojiSequencesData() {
}

func main() {
	var flagSequence bool
	flag.BoolVar(&flagSequence, "s", false, "import emoji sequences data")
	flag.Parse()

	if flagSequence {
		outputEmojiSequencesData()
	} else {
		outputEmojiData()
	}
}
