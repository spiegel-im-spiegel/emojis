package emojis

import (
	"bytes"
	_ "embed"

	"github.com/spiegel-im-spiegel/emojis/json"
	"github.com/spiegel-im-spiegel/errs"
)

//go:embed json/emoji-sequences.json
var jsonEmojiSequenceText []byte

// NewEmojiSequenceList returns list of json.EmojiSequence data.
func NewEmojiSequenceList() ([]json.EmojiSequence, error) {
	list, err := json.DecodeEmojiSequence(bytes.NewReader(jsonEmojiSequenceText))
	if err != nil {
		return nil, errs.Wrap(err)
	}
	return list, nil
}

// MappingEmojiSequence maps json.EmojiSequence data.
func MappingEmojiSequence(list []json.EmojiSequence) map[string]*json.EmojiSequence {
	emap := map[string]*json.EmojiSequence{}
	for i := 0; i < len(list); i++ {
		emap[list[i].Sequence] = &list[i]
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
