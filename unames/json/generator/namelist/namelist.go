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
