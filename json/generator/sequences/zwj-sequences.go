package sequences

import (
	"bufio"
	"io"
	"strings"

	emj "github.com/kyokomi/emoji/v2"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/fetch"
)

const emojiZwjSequencesFile = "https://www.unicode.org/Public/emoji/13.1/emoji-zwj-sequences.txt"

func zwjSequencesListFile() (io.ReadCloser, error) {
	u, err := fetch.URL(emojiZwjSequencesFile)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("url", emojiZwjSequencesFile))
	}
	resp, err := fetch.New().Get(u)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("url", emojiZwjSequencesFile))
	}
	return resp.Body(), nil
}

func parseZwjSequences(list map[string]EmojiSequence) (map[string]EmojiSequence, error) {
	r, err := zwjSequencesListFile()
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
		if len(flds) < 2 {
			continue
		}
		seq, err := getRuneSequence(flds[0])
		if err != nil {
			continue
		}
		list[seq] = EmojiSequence{Sequence: seq, Name: getDescription(flds[2]), SequenceType: getSequenceType(flds[1]), Shortcodes: emj.RevCodeMap()[seq]}

	}

	if err := scanner.Err(); err != nil {
		return nil, errs.Wrap(err)
	}
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
