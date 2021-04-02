package data

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/spiegel-im-spiegel/emojis/json"
	"github.com/spiegel-im-spiegel/emojis/json/generator/unames"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/fetch"
)

const emojiDataFile = "https://www.unicode.org/Public/UCD/latest/ucd/emoji/emoji-data.txt"

func dataListFile() (io.ReadCloser, error) {
	u, err := fetch.URL(emojiDataFile)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("url", emojiDataFile))
	}
	resp, err := fetch.New().Get(u)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("url", emojiDataFile))
	}
	return resp.Body(), nil
}

func parseData(list map[rune]json.EmojiData) (map[rune]json.EmojiData, error) {
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
					ed = json.EmojiData{
						Code: r,
						Name: name,
					}
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

func setProperty(e json.EmojiData, s string) json.EmojiData {
	if 0x1F1E6 <= e.Code && e.Code <= 0x1F1FF {
		e.RegionalIndicator = true
	}
	var prop EmojiProperty
	if strings.Contains(s, "#") {
		flds := strings.Split(s, "#")
		if len(flds) < 1 {
			return e
		}
		prop = GetEmojiProperty(strings.TrimSpace(flds[0]))
	} else {
		prop = GetEmojiProperty(strings.TrimSpace(s))
	}
	switch prop {
	case PropEmojiCharacter:
		e.Emoji = true
	case PropEmojiPresentation:
		e.EmojiPresentation = true
	case PropEmojiModifier:
		e.EmojiModifier = true
	case PropEmojiModifierBase:
		e.EmojiModifierBase = true
	case PropEmojiComponent:
		e.EmojiComponent = true
	case PropExtendedPictographic:
		e.ExtendedPictographic = true
	}
	return e
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
