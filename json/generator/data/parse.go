package data

import (
	"encoding/json"
	"io"
	"sort"

	"github.com/spiegel-im-spiegel/errs"
)

// EmojiData is entity of "emoji-data.txt" and "emoji-variation-sequences.txt".
type EmojiData struct {
	Code                 rune
	Name                 string
	Emoji                bool     `json:",omitempty"`
	EmojiPresentation    bool     `json:"Emoji_Presentation,omitempty"`
	EmojiModifier        bool     `json:"Emoji_Modifier,omitempty"`
	EmojiModifierBase    bool     `json:"Emoji_Modifier_Base,omitempty"`
	EmojiComponent       bool     `json:"Emoji_Component,omitempty"`
	ExtendedPictographic bool     `json:"Extended_Pictographic,omitempty"`
	VariationTextStyle   string   `json:",omitempty"`
	VariationEmojiStyle  string   `json:",omitempty"`
	Shortcodes           []string `json:",omitempty"`
}

// Parse returns EmojiData list.
func Parse() (map[rune]EmojiData, error) {
	list := map[rune]EmojiData{}
	var err error
	list, err = parseData(list)
	if err != nil {
		return list, errs.Wrap(err)
	}
	list, err = parseVariation(list)
	if err != nil {
		return list, errs.Wrap(err)
	}
	return list, nil
}

// EncodeJSON outputs emoji-data with JSON format.
func EncodeJSON(w io.Writer, emap map[rune]EmojiData) error {
	list := []EmojiData{}
	for _, v := range emap {
		list = append(list, v)
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].Code < list[j].Code
	})
	enc := json.NewEncoder(w)
	if err := enc.Encode(list); err != nil {
		return errs.Wrap(err)
	}
	return nil
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
