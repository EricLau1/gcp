
Dep:

    go get cloud.google.com/go/translate
    go get golang.org/x/text/language
    go get google.golang.org/api/option

Ref:

https://github.com/GoogleCloudPlatform/golang-samples/blob/master/translate/snippets/snippet.go


Example:

	// Look at your API KEY in the GOOGLE CLOUD PLATFORM

	// Set environment variable with your api key
	apiKey := os.Getenv("GCP_TRANSLATION_API_KEY")
	if apiKey == "" {
		apiKey = "API KEY VALID"
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