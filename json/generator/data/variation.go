package data

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/spiegel-im-spiegel/emojis/json"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/fetch"
)

const (
	emojiVariationSequencesFile = "https://www.unicode.org/Public/UCD/latest/ucd/emoji/emoji-variation-sequences.txt"
	textPresentationSelector    = 0xFE0E
	emojiPresentationSelector   = 0xFE0F
)

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

func parseVariation(list map[rune]json.EmojiData) (map[rune]json.EmojiData, error) {
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
		sq := string([]rune{fst, snd})
		switch snd {
		case textPresentationSelector:
			ed.VariationTextStyle = sq
		case emojiPresentationSelector:
			ed.VariationEmojiStyle = sq
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

/* MIT License
 *
 * Copyright 2021 Spiegel
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
