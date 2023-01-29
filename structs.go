package mstranslator

import "encoding/xml"

type XMLIntValue struct {
	Text int `xml:",chardata"`
}

type XMLStringValue struct {
	Text string `xml:",chardata"`
}

type ResponseTranslationMatch struct {
	Count               XMLIntValue    `xml:"Count"`
	MatchDegree         XMLIntValue    `xml:"MatchDegree"`
	MatchedOriginalText XMLStringValue `xml:"MatchedOriginalText"`
	Rating              XMLIntValue    `xml:"Rating"`
	TranslatedText      XMLStringValue `xml:"TranslatedText"`
}

type ResponseTranslations struct {
	TransMatch []ResponseTranslationMatch `xml:"TranslationMatch"`
}

type GetTranslationResponse struct {
	XMLName      xml.Name             `xml:"GetTranslationResponse"`
	Translations ResponseTranslations `xml:"Translations"`
}

type ResponseToken struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
	Scope       string `json:"scope"`
}

type TransformTextResponse struct {
	ErrorCondition   int    `json:"ec"`
	ErrorDescriptive string `json:"em"`
	Sentence         string `json:"sentence"`
}

type ResponseXML struct {
	XMLName   xml.Name `xml:"string"`
	Namespace string   `xml:"xmlns,attr"`
	Value     string   `xml:",innerxml"`
}

type ResponseArray struct {
	XMLName           xml.Name `xml:"ArrayOfstring"`
	Namespace         string   `xml:"xmlns,attr"`
	InstanceNamespace string   `xml:"xmlns:i,attr"`
	Strings           []string `xml:"string"`
}
