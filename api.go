package mstranslator

type Client struct {
	tranlsateProvider *TranslateProvider
	languageProvider  *LanguageProvider
	authenticator     *Authenticator
}

func NewClient(clientID, clientSecret string) *Client {
	auth := NewAuthenticator(clientID, clientSecret)
	
	if auth == nil {
		return nil
	}

	auth.GetToken()

	return &Client{
		tranlsateProvider: NewTranslateProvider(auth),
		languageProvider: NewLanguageProvider(auth),
		authenticator: auth,
	}
}

func (c *Client) Translate(text, from, to string) (string, error) {
	return c.tranlsateProvider.Translate(text, from, to)
}

func (c *Client) TransformText(lang, category, text string) (string, error) {
	return c.tranlsateProvider.TransformText(lang, category, text)
}

func (c *Client) Speak(text, lang, outFormat string) ([]byte, error) {
	return c.tranlsateProvider.Speak(text, lang, outFormat)
}

func (c *Client) Detect(text string) (string, error) {
	return c.languageProvider.Detect(text)
}

func (c *Client) DetectArray(text []string) ([]string, error) {
	return c.languageProvider.DetectArray(text)
}

func (c *Client) GetTranslations(text, from, to string, maxTranslations int) ([]ResponseTranslationMatch, error) {
	return c.languageProvider.GetTranslations(text, from, to, maxTranslations)
}

func (c *Client) GetLanguageNames(codes []string) ([]string, error) {
	return c.languageProvider.GetLanguageNames(codes)
}

func (c *Client) GetLanguagesForTranslate() ([]string, error) {
	return c.languageProvider.GetLanguagesForTranslate()
}

func (c *Client) GetLanguagesForSpeak() ([]string, error) {
	return c.languageProvider.GetLanguagesForSpeak()
}
