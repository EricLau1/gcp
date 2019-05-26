package translator

import (
	"testing"
)

var content = []string{
	"Hello World",
	"Hola Mundo",
	"你好，世界",
	"こんにちは世界",
	"Hallo Welt",
	"Bonjour le monde",
	"Привет, мир",
	"Ciao mondo",
}

type Language struct {
	Expression string
	Source     string
}

type Languages []Language

var inputs = Languages{
	Language{
		Expression: "Hello World",
		Source:     "en",
	},
	Language{
		Expression: "Hola Mundo",
		Source:     "es",
	},
	Language{
		Expression: "你好，世界",
		Source:     "zh-CN",
	},
	Language{
		Expression: "こんにちは世界",
		Source:     "ja",
	},
	Language{
		Expression: "Hallo Welt",
		Source:     "de",
	},
	Language{
		Expression: "Bonjour le monde",
		Source:     "fr",
	},
	Language{
		Expression: "Привет, мир",
		Source:     "ru",
	},
	Language{
		Expression: "Ciao mondo",
		Source:     "it",
	},
}

func TestTranslate(t *testing.T) {

	API_KEY = "YOUR API KEY GOOGLE CLOUD PLATFORM"

	var content []string

	for _, input := range inputs {
		content = append(content, input.Expression)
	}

	response, err := NewClient().Content(content).To(DEFAULT_TARGET).Translate()

	if err != nil {
		t.Errorf("%s", err.Error())
	}

	for i, translation := range response {
		source := translation.Source.String()
		if source != inputs[i].Source {
			t.Errorf("input: (%s), expected (%s), received (%s)", inputs[i].Expression, inputs[i].Source, source)
		}
	}

}
