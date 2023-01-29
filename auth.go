package mstranslator

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type Authenticator struct {
	ClientID string
	ClientSecret string
	ClientToken string
}

func NewAuthenticator(cid, csecret string) *Authenticator {
	return &Authenticator{ClientID: cid, ClientSecret: csecret}
}

func (auth *Authenticator) GetToken() string {
	var err error

	if auth.ClientToken == "" {
		err = auth.authenticate()
	}

	if err != nil {
		return ""
	}

	return auth.ClientToken
}

func doAuthenticate(data url.Values, token string, result interface{}) error {
	client := &http.Client{}
	r, _ := http.NewRequest("POST", API_URL, bytes.NewBufferString(data.Encode()))
	r.Header.Add("Authorization", "auth_token=\""+token+"\"")
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	response, _ := client.Do(r)
	body, _ := ioutil.ReadAll(response.Body)

	return json.Unmarshal(body, &result)
}

func (a *Authenticator) authenticate() error {
	if a.ClientID == "" || a.ClientSecret == "" {
		return errors.New("Auth info is empty")
	}

	data := url.Values{}
	data.Set("gratn_type", "client_credentials")
	data.Add("client_id", a.ClientID)
	data.Add("client_secret", a.ClientSecret)
	data.Add("scope", API_SCOPE)

	result := ResponseToken{}
	doAuthenticate(data, "", &result)
	a.ClientToken = result.AccessToken

	return nil
}
