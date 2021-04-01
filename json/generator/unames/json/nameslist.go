package json

import (
	jsonorg "encoding/json"
	"io"

	"github.com/spiegel-im-spiegel/errs"
)

// UnicodeName is Unicode name info
type UnicodeName struct {
	Code rune
	Name string
}

// EncodeUnicodeName outputs []UnicodeName data with JSON format.
func EncodeUnicodeName(w io.Writer, ulist []UnicodeName) error {
	if err := jsonorg.NewEncoder(w).Encode(ulist); err != nil {
		return errs.Wrap(err)
	}
	return nil
}

// DecodeUnicodeName returns []UnicodeName data from JSON format text.
func DecodeUnicodeName(r io.Reader) ([]UnicodeName, error) {
	list := []UnicodeName{}
	if err := jsonorg.NewDecoder(r).Decode(&list); err != nil {
		return list, errs.Wrap(err)
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
