package tts

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var baseUrl = "https://translate.google.com/translate_tts?ie=UTF-8&q=%s&tl=%s&client=tw-ob"

type Config struct {
	Language string
	Speak    string
}

type Speech struct {
	bytes.Buffer
}

func Speak(t Config) (*Speech, error) {
	req := fmt.Sprintf(baseUrl, url.QueryEscape(t.Speak), url.QueryEscape(t.Language))
	res, err := http.Get(req)
	if err != nil {
		return nil, err
	}

	speech := &Speech{}
	if _, err := io.Copy(&speech.Buffer, res.Body); err != nil {
		return nil, err
	}

	return speech, nil
}
