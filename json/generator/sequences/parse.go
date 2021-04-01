package sequences

import (
	"sort"

	"github.com/spiegel-im-spiegel/emojis/json"
	"github.com/spiegel-im-spiegel/errs"
)

// Parse returns EmojiSequence list.
func Parse() ([]json.EmojiSequence, error) {
	emap := map[string]json.EmojiSequence{}
	var err error
	emap, err = parseSequences(emap)
	if err != nil {
		return nil, errs.Wrap(err)
	}
	emap, err = parseZwjSequences(emap)
	if err != nil {
		return nil, errs.Wrap(err)
	}

	list := []json.EmojiSequence{}
	for _, v := range emap {
		list = append(list, v)
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].Sequence < list[j].Sequence
	})
	return list, nil
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
