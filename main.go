package main

import (
	"encoding/json"
	"fmt"
	"go-gcp-translator/translator"
	"log"
	"os"

	"golang.org/x/text/language"
)

func main() {

	// Look at your API KEY in the GOOGLE CLOUD PLATFORM

	// Define env var with your api key
	apiKey := os.Getenv("GCP_TRANSLATION_API_KEY")
	if apiKey == "" {
		apiKey = "YOUR API KEY"
	}

	translator.API_KEY = apiKey

	content := []string{
		"Hello World",
		"Hola Mundo",
		"你好，世界",
		"こんにちは世界",
		"Hallo Welt",
		"Bonjour le monde",
		"Привет, мир",
		"Ciao mondo",
	}

	translations, err := translator.NewClient().Content(content).To(language.BrazilianPortuguese).Translate()
	if err != nil {
		log.Fatal(err)
	}

	data, err := json.MarshalIndent(translations, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))

}
