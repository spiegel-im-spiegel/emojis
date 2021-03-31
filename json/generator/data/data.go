package data

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/spiegel-im-spiegel/emojis/unames"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/fetch"
)

const emojidataFile = "https://www.unicode.org/Public/UCD/latest/ucd/emoji/emoji-data.txt"

func dataListFile() (io.ReadCloser, error) {
	u, err := fetch.URL(emojidataFile)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("url", emojidataFile))
	}
	resp, err := fetch.New().Get(u)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("url", emojidataFile))
	}
	return resp.Body(), nil
}

func parseData(list map[rune]EmojiData) (map[rune]EmojiData, error) {
	r, err := dataListFile()
	if err != nil {
		return list, errs.Wrap(err)
	}
	defer r.Close()

	names, err := unames.New()
	if err != nil {
		return list, errs.Wrap(err)
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 || strings.ContainsAny(string(text[0]), "# \t") {
			continue
		}
		flds := strings.Split(text, ";")
		if len(flds) < 2 {
			continue
		}
		from, to, err := getRuneRange(flds[0])
		if err != nil {
			continue
		}
		for r := from; r <= to; r++ {
			name := names.Name(r)
			if len(name) > 0 {
				ed, ok := list[r]
				if !ok {
					ed = EmojiData{Code: r, Name: name}
				}
				list[r] = setProperty(ed, flds[1])
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, errs.Wrap(err)
	}
	return list, nil
}

func getRuneRange(s string) (rune, rune, error) {
	var from, to string
	if strings.Contains(s, "..") {
		flds := strings.Split(s, "..")
		if len(flds) < 2 {
			return 0, 0, os.ErrInvalid
		}
		from = strings.TrimSpace(flds[0])
		to = strings.TrimSpace(flds[1])
	} else {
		from = strings.TrimSpace(s)
		to = from
	}
	fromR, err := strconv.ParseUint(from, 16, 32)
	if err != nil {
		return 0, 0, os.ErrInvalid
	}
	toR, err := strconv.ParseUint(to, 16, 32)
	if err != nil {
		return 0, 0, os.ErrInvalid
	}
	return rune(fromR), rune(toR), nil
}

func setProperty(e EmojiData, s string) EmojiData {
	var prop string
	if strings.Contains(s, "#") {
		flds := strings.Split(s, "#")
		if len(flds) < 1 {
			return e
		}
		prop = strings.TrimSpace(flds[0])
	} else {
		prop = strings.TrimSpace(s)
	}
	switch {
	case strings.EqualFold(prop, "Emoji"):
		e.Emoji = true
	case strings.EqualFold(prop, "Emoji_Presentation"):
		e.EmojiPresentation = true
	case strings.EqualFold(prop, "Emoji_Modifier"):
		e.EmojiModifier = true
	case strings.EqualFold(prop, "Emoji_Modifier_Base"):
		e.EmojiModifierBase = true
	case strings.EqualFold(prop, "Emoji_Component"):
		e.EmojiComponent = true
	case strings.EqualFold(prop, "Extended_Pictographic"):
		e.ExtendedPictographic = true
	}
	return e
}
