package types

import (
	"strconv"
	"strings"
)

type SequencesType int

const (
	TypeUnknown SequencesType = iota
	TypeBasicEmoji
	TypeEmojiKeycapSequence
	TypeRGIEmojiFlagSequence
	TypeRGIEmojiTagSequence
	TypeRGIEmojiModifierSequence
	TypeRGIEmojiZWJSequence
)

var sequencesTypeMap = map[SequencesType]string{
	TypeBasicEmoji:               "Basic_Emoji",
	TypeEmojiKeycapSequence:      "Emoji_Keycap_Sequence",
	TypeRGIEmojiFlagSequence:     "RGI_Emoji_Flag_Sequence",
	TypeRGIEmojiTagSequence:      "RGI_Emoji_Tag_Sequence",
	TypeRGIEmojiModifierSequence: "RGI_Emoji_Modifier_Sequence",
	TypeRGIEmojiZWJSequence:      "RGI_Emoji_ZWJ_Sequence",
}

func GetSequenceType(s string) SequencesType {
	s = strings.TrimSpace(s)
	for k, v := range sequencesTypeMap {
		if strings.EqualFold(s, v) {
			return k
		}
	}
	return TypeUnknown
}

func (t SequencesType) String() string {
	return sequencesTypeMap[t]
}

//UnmarshalJSON returns result of Unmarshal for json.Unmarshal()
func (t *SequencesType) UnmarshalJSON(b []byte) error {
	s, err := strconv.Unquote(string(b))
	if err != nil {
		s = string(b)
	}
	*t = GetSequenceType(s)
	return nil
}

//MarshalJSON returns string
func (t *SequencesType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte("\"\""), nil
	}
	return []byte(strconv.Quote(t.String())), nil
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
