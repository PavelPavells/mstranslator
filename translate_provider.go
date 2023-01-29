package mstranslator

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type TranslateProvider struct {
	authenticator *Authenticator
}

func NewTranslateProvider(auth *Authenticator) *TranslateProvider {
	return &TranslateProvider{authenticator: auth}
}

func (t *TranslateProvider) Translate(text, from, to string) (string, error) {
	token := t.authenticator.GetToken()

	uri := fmt.Sprintf("%s?text=%s&from=%s&to=%s", TranslationURL, url.QueryEscape(text), url.QueryEscape(from), url.QueryEscape(to))

	client := http.Client{}
	request, err := http.NewRequest("GET", uri, nil)
	request.Header.Add("Content-Type", "text/plain")
	request.Header.Add("Authorization", "Bearer " + token)

	response, err := client.Do(request)

	if err != nil {
		return "", nil
	}

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		return "", nil
	}

	translation := &ResponseXML{}
	err = xml.Unmarshal(body, &translation)

	if err != nil {
		return "", nil
	}

	return translation.Value, nil
}

func (t *TranslateProvider) TransformText(lang, category, text string) (string, error) {
	token := t.authenticator.GetToken()

	uri := fmt.Sprintf(
		"%s?sentence=%s&category=%s&language=%s",
		TransformTextURL,
		url.QueryEscape(text),
		url.QueryEscape(category),
		url.QueryEscape(lang))

	client := &http.Client{}
	request, err := http.NewRequest("GET", uri, nil)
	request.Header.Add("Content-Type", "text/plain")
	request.Header.Add("Authorization", "Bearer "+token)

	response, err := client.Do(request)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return "", err
	}

	// Microsoft Server json response contain BOM, need to trim.
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))

	transTransform := TransformTextResponse{}
	err = json.Unmarshal(body, &transTransform)
	if err != nil {
		return "", err
	}

	return transTransform.Sentence, nil
}

func (t *TranslateProvider) Speak(text, lang, outFormat string) ([]byte, error) {
	token := t.authenticator.GetToken()

	uri := fmt.Sprintf(
		"%s?text=%s&language=%s&format=%s",
		SpeakURL,
		url.QueryEscape(text),
		url.QueryEscape(lang),
		url.QueryEscape(outFormat))

	client := &http.Client{}
	request, err := http.NewRequest("GET", uri, nil)
	request.Header.Add("Authorization", "Bearer "+token)

	var retBuffer []byte
	response, err := client.Do(request)
	if err != nil {
		return retBuffer, err
	}

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return retBuffer, err
	}

	return body, nil
}
