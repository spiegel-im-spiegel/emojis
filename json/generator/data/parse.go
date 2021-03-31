package data

import (
	"encoding/json"
	"io"
	"sort"

	"github.com/spiegel-im-spiegel/errs"
)

type EmojiData struct {
	Code                 rune
	Name                 string
	Emoji                bool
	EmojiPresentation    bool
	EmojiModifier        bool
	EmojiModifierBase    bool
	EmojiComponent       bool
	ExtendedPictographic bool
	PresentationText     bool
	PresentationEmoji    bool
}

// Parse returns emoji-data.
func Parse() (map[rune]EmojiData, error) {
	list := map[rune]EmojiData{}
	var err error
	list, err = parseData(list)
	if err != nil {
		return list, errs.Wrap(err)
	}
	list, err = parseVariation(list)
	if err != nil {
		return list, errs.Wrap(err)
	}
	return list, nil
}

// EncodeJSON outputs emoji-data with JSON format.
func EncodeJSON(w io.Writer, emap map[rune]EmojiData) error {
	list := []EmojiData{}
	for _, v := range emap {
		list = append(list, v)
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].Code < list[j].Code
	})
	enc := json.NewEncoder(w)
	if err := enc.Encode(list); err != nil {
		return errs.Wrap(err)
	}
	return nil
}
