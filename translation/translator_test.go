package translation_test

import (
	"testing"

	"github.com/logicmason/cdig/translation"
)

func TestTranslate(t *testing.T) {
	// Arrange
	tt := []struct {
		Word        string
		Language    string
		Translation string
	}{
		{
			Word:        "hello",
			Language:    "english",
			Translation: "hello",
		},
		{
			Word:        "hello",
			Language:    "german",
			Translation: "hallo",
		},
		{
			Word:        "Hello",
			Language:    "german",
			Translation: "hallo",
		},
		{
			Word:        "hello",
			Language:    "German",
			Translation: "hallo",
		},
		{
			Word:        "hello ",
			Language:    "german",
			Translation: "hallo",
		},
		{
			Word:        "hello",
			Language:    "finnish",
			Translation: "hei",
		},
		{
			Word:        "hello",
			Language:    "klingon",
			Translation: "",
		},
		{
			Word:        "bye",
			Language:    "german",
			Translation: "",
		}, {
			Word:        "bye",
			Language:    "klingon",
			Translation: "",
		},
	}

	for _, test := range tt {
		res := translation.Translate(test.Word, test.Language)

		if res != test.Translation {
			t.Errorf(
				`expected "%s" in "%s" to be "%s" but received "%s"`,
				test.Word, test.Language, test.Translation, res)

		}
	}
}
