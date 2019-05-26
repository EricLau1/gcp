package translator

import (
	"context"
	"errors"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
)

type Google struct {
	Client  *translate.Client `json:","`
	Context context.Context   `json:","`
	ApiKey  string            `json:","`
}

type Translator struct {
	Engine       *Google                 `json:","`
	Text         []string                `json:"text"`
	Target       language.Tag            `json:"target"`
	Translations []translate.Translation `json:"translations"`
	IsValid      bool                    `json:"is_valid"`
	Errors       []string                `json:"errors"`
}

var (
	DEFAULT_TARGET = language.BrazilianPortuguese
	API_KEY        = "YOUR API KEY" // Look at your API KEY in the GOOGLE CLOUD PLATFORM
)

func NewClient() *Translator {

	var err error

	translator := &Translator{Engine: &Google{ApiKey: API_KEY}}

	translator.Engine.Context = context.Background()

	translator.Engine.Client, err = translate.NewClient(
		translator.Engine.Context,
		option.WithAPIKey(translator.Engine.ApiKey))
	if err != nil {
		translator.IsValid = false
		translator.Errors = append(translator.Errors, err.Error())
		return translator
	}
	translator.IsValid = true
	return translator
}

func (translator *Translator) Content(content []string) *Translator {
	translator.Text = append(translator.Text, content...)
	return translator
}

func (translator *Translator) To(target language.Tag) *Translator {
	translator.Target = target
	return translator
}

func (translator *Translator) Translate() ([]translate.Translation, error) {

	if translator.IsValid {
		resp, err := translator.Engine.Client.Translate(
			translator.Engine.Context,
			translator.Text,
			translator.Target,
			nil)

		if err != nil {
			return nil, err
		}
		return resp, nil
	}
	return nil, errors.New("Invalid translator")
}
