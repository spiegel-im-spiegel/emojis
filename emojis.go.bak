package emojis

import (
	"sort"
	"strings"

	"github.com/spiegel-im-spiegel/emojis/types"
	"github.com/spiegel-im-spiegel/errs"
)

// Emoji is info for emoji character or sequence
type Emoji struct {
	Code         string
	Name         string
	SequenceType types.SequencesType
	Shortcodes   []string `json:",omitempty"`
}

// NewEmojiList returns list of Emojis.
func NewEmojiList() ([]Emoji, error) {
	esmap := map[string]Emoji{}

	clist, err := NewEmojiCodeList()
	if err != nil {
		return nil, errs.Wrap(err)
	}
	for i := 0; i < len(clist); i++ {
		code := string([]rune{clist[i].Code})
		if clist[i].EmojiModifierBase {
			esmap[code] = Emoji{
				Code:         code,
				Name:         clist[i].Name,
				SequenceType: types.TypeEmojiModifierBase,
				Shortcodes:   clist[i].Shortcodes,
			}
		} else if clist[i].ExtendedPictographic {
			esmap[code] = Emoji{
				Code:         code,
				Name:         clist[i].Name,
				SequenceType: types.TypeExtendedPictographic,
				Shortcodes:   clist[i].Shortcodes,
			}
		} else if clist[i].EmojiPresentation {
			esmap[code] = Emoji{
				Code:         code,
				Name:         clist[i].Name,
				SequenceType: types.TypeEmojiPresentation,
				Shortcodes:   clist[i].Shortcodes,
			}
		}
		if len(clist[i].VariationEmojiStyle) > 0 {
			esmap[clist[i].VariationEmojiStyle] = Emoji{
				Code:         clist[i].VariationEmojiStyle,
				Name:         clist[i].Name,
				SequenceType: types.TypeEmojiPresentationSequence,
				Shortcodes:   clist[i].ShortcodesVariationEmoji,
			}
		}
	}

	slist, err := NewEmojiSequenceList()
	if err != nil {
		return nil, errs.Wrap(err)
	}
	for i := 0; i < len(slist); i++ {
		if _, ok := esmap[slist[i].Sequence]; !ok {
			esmap[slist[i].Sequence] = Emoji{
				Code:         slist[i].Sequence,
				Name:         slist[i].Name,
				SequenceType: slist[i].SequenceType,
				Shortcodes:   slist[i].Shortcodes,
			}
		}
	}
	es := []Emoji{}
	for _, v := range esmap {
		es = append(es, v)
	}
	sort.Slice(es, func(i, j int) bool {
		return strings.Compare(es[i].Code, es[j].Code) < 0
	})
	return es, nil
}

// MappingEmojiList maps Emoji data.
func MappingEmojiList(list []Emoji) map[string]*Emoji {
	emap := map[string]*Emoji{}
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
