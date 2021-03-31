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
