package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spiegel-im-spiegel/emojis"
)

func dump(s string) string {
	ss := []string{}
	for _, r := range s {
		ss = append(ss, fmt.Sprintf("`%U`", r))
	}
	return strings.Join(ss, " + ")
}

func main() {
	list, err := emojis.NewEmojiSequenceList()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println("| Char  | Code Point | Name | Type | Shortcodes |")
	fmt.Println("| :---: | ---------- | ---- | ---- | ---------- |")
	for _, ec := range list {
		var bldr strings.Builder
		for _, c := range ec.Shortcodes {
			bldr.WriteString(fmt.Sprintf(" `%s`", c))
		}
		fmt.Printf("| <abbr class='emoji-chars' title='%[3]v'>%[1]v</abbr> | %v |  %v | %v |%s |\n", ec.Sequence, dump(ec.Sequence), ec.Name, ec.SequenceType, bldr.String())
	}
}
