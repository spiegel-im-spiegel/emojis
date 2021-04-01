# [emojis] -- List of Emoji-Sequences

[![check vulns](https://github.com/spiegel-im-spiegel/emojis/workflows/vulns/badge.svg)](https://github.com/spiegel-im-spiegel/emojis/actions)
[![lint status](https://github.com/spiegel-im-spiegel/emojis/workflows/lint/badge.svg)](https://github.com/spiegel-im-spiegel/emojis/actions)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/emojis/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/spiegel-im-spiegel/emojis.svg)](https://github.com/spiegel-im-spiegel/emojis/releases/latest)

This package is required Go 1.16 or later.

## Import

```go
import "github.com/spiegel-im-spiegel/emojis"
```

## Usage

```go
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
	list, err := emojis.NewEmojiList()
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
		fmt.Printf("| %v | %v |  %v | %v |%s |\n", ec.Code, dump(ec.Code), ec.Name, ec.SequenceType, bldr.String())
	}
}
```

## Refer

- [UTS #51: Unicode Emoji](http://www.unicode.org/reports/tr51/)
- [Emoji List, v13.1](http://www.unicode.org/emoji/charts/emoji-list.html)

### Unicode Names List (Latest)

- [https://www.unicode.org/Public/UCD/latest/ucd/](https://www.unicode.org/Public/UCD/latest/ucd/)
  - [UCD: Unicode NamesList File Format](https://www.unicode.org/Public/UCD/latest/ucd/NamesList.html)
  - [NamesList.txt](https://www.unicode.org/Public/UCD/latest/ucd/NamesList.txt)

### Emoji Data (Latest)

- [https://www.unicode.org/Public/UCD/latest/ucd/emoji/](https://www.unicode.org/Public/UCD/latest/ucd/emoji/)
  - [ReadMe.txt](https://www.unicode.org/Public/UCD/latest/ucd/emoji/ReadMe.txt)
  - [emoji-data.txt](https://www.unicode.org/Public/UCD/latest/ucd/emoji/emoji-data.txt)
  - [emoji-variation-sequences.txt](https://www.unicode.org/Public/UCD/latest/ucd/emoji/emoji-variation-sequences.txt)

### Emoji Sequences (v13.1)

- [https://unicode.org/Public/emoji/](https://unicode.org/Public/emoji/)
  - [13.1/](https://unicode.org/Public/emoji/13.1/)
    - [ReadMe.txt](https://unicode.org/Public/emoji/13.1/ReadMe.txt)
    - [emoji-sequences.txt](https://unicode.org/Public/emoji/13.1/emoji-sequences.txt)
    - [emoji-zwj-sequences.txt](https://unicode.org/Public/emoji/13.1/emoji-zwj-sequences.txt)

[emojis]: https://github.com/spiegel-im-spiegel/emojis "spiegel-im-spiegel/emojis: List of Emoji-Sequences"
