package unames

import (
	_ "embed"
	"encoding/json"

	"github.com/spiegel-im-spiegel/errs"
)

//go:embed json/nameslist.json
var jsonText []byte

// NamesMap is class Unicode name list.
type NamesMap struct {
	nmap map[rune]string
}

// New returns a new NamesMap instance.
func New() (NamesMap, error) {
	list := []struct {
		Code rune
		Name string
	}{}
	if err := json.Unmarshal(jsonText, &list); err != nil {
		return NamesMap{}, errs.Wrap(err)
	}
	nmap := map[rune]string{}
	for _, n := range list {
		nmap[n.Code] = n.Name
	}
	return NamesMap{nmap}, nil
}

// Name returns string of name for Unicode point.
func (nm NamesMap) Name(r rune) string {
	return nm.nmap[r]
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
