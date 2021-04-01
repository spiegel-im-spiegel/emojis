package emojis

import (
	"bytes"
	_ "embed"

	"github.com/spiegel-im-spiegel/emojis/json"
	"github.com/spiegel-im-spiegel/errs"
)

//go:embed json/emoji-data.json
var jsonEmojidataText []byte

// NewEmojiCodeList returns list of json.EmojiData data.
func NewEmojiCodeList() ([]json.EmojiData, error) {
	list, err := json.DecodeEmojiData(bytes.NewReader(jsonEmojidataText))
	if err != nil {
		return nil, errs.Wrap(err)
	}
	return list, nil
}

// MappingEmojiCode maps json.EmojiData data.
func MappingEmojiCode(list []json.EmojiData) map[rune]*json.EmojiData {
	emap := map[rune]*json.EmojiData{}
	for i := 0; i < len(list); i++ {
		emap[list[i].Code] = &list[i]
	}
	return emap
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
