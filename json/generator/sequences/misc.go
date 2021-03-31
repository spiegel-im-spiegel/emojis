package sequences

import (
	"os"
	"strconv"
	"strings"
)

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

func getSequenceType(s string) SequencesType {
	s = strings.TrimSpace(s)
	for k, v := range sequencesTypeMap {
		if strings.EqualFold(s, v) {
			return k
		}
	}
	return TypeUnknown
}

func getRuneSequence(s string) (string, error) {
	flds := strings.Fields(s)
	runes := []rune{}
	for _, s := range flds {
		r, err := strconv.ParseUint(s, 16, 32)
		if err != nil {
			return "", os.ErrInvalid
		}
		runes = append(runes, rune(r))
	}
	return string(runes), nil
}

func getDescription(s string) string {
	flds := strings.Split(s, "#")
	if len(flds) == 0 {
		return ""
	}
	return strings.ReplaceAll(strings.TrimSpace(flds[0]), `\x{23}`, "#")
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
