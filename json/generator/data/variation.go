package data

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/fetch"
)

const emojiVariationSequencesFile = "https://www.unicode.org/Public/UCD/latest/ucd/emoji/emoji-variation-sequences.txt"

func variationListFile() (io.ReadCloser, error) {
	u, err := fetch.URL(emojiVariationSequencesFile)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("url", emojiVariationSequencesFile))
	}
	resp, err := fetch.New().Get(u)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("url", emojiVariationSequencesFile))
	}
	return resp.Body(), nil
}

func parseVariation(list map[rune]EmojiData) (map[rune]EmojiData, error) {
	r, err := variationListFile()
	if err != nil {
		return list, errs.Wrap(err)
	}
	defer r.Close()

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 || strings.ContainsAny(string(text[0]), "# \t") {
			continue
		}
		flds := strings.Split(text, ";")
		if len(flds) < 1 {
			continue
		}
		fst, snd, err := getRuneSequence(flds[0])
		if err != nil {
			continue
		}
		ed, ok := list[fst]
		if !ok {
			continue
		}
		switch snd {
		case 0xFE0E:
			ed.PresentationText = true
		case 0xFE0F:
			ed.PresentationEmoji = true
		}
		list[fst] = ed
	}

	if err := scanner.Err(); err != nil {
		return nil, errs.Wrap(err)
	}
	return list, nil
}

func getRuneSequence(s string) (rune, rune, error) {
	flds := strings.Split(strings.TrimSpace(s), " ")
	if len(flds) < 2 {
		return 0, 0, os.ErrInvalid
	}
	fromR, err := strconv.ParseUint(flds[0], 16, 32)
	if err != nil {
		return 0, 0, os.ErrInvalid
	}
	toR, err := strconv.ParseUint(flds[1], 16, 32)
	if err != nil {
		return 0, 0, os.ErrInvalid
	}
	return rune(fromR), rune(toR), nil
}
