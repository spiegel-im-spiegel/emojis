package namelist

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/fetch"
)

const unicodeNamesListFile = "https://www.unicode.org/Public/UCD/latest/ucd/NamesList.txt"

func nameListFile() (io.ReadCloser, error) {
	u, err := fetch.URL(unicodeNamesListFile)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("url", unicodeNamesListFile))
	}
	resp, err := fetch.New().Get(u)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("url", unicodeNamesListFile))
	}
	return resp.Body(), nil
}

type UnicodeName struct {
	Code rune
	Name string
}

func Parse() ([]UnicodeName, error) {
	r, err := nameListFile()
	if err != nil {
		return nil, errs.Wrap(err)
	}
	defer r.Close()

	list := []UnicodeName{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 || strings.ContainsAny(string(text[0]), ";@ \t") {
			continue
		}
		flds := strings.Fields(text)
		if len(flds) < 2 {
			continue
		}
		code, err := strconv.ParseUint(flds[0], 16, 32)
		if err != nil {
			continue
		}
		name := strings.Join(flds[1:], " ")
		if strings.HasPrefix(name, "<") && strings.HasSuffix(name, ">") {
			continue
		}
		list = append(list, UnicodeName{Code: rune(code), Name: name})
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
