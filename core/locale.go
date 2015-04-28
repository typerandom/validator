package core

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type Locale struct {
	Messages map[string]string
}

func NewLocale() *Locale {
	return &Locale{
		Messages: make(map[string]string),
	}
}

func (this *Locale) Set(key string, value string) {
	this.Messages[key] = value
}

func (this *Locale) Get(key string) (string, error) {
	if val, ok := this.Messages[key]; ok {
		return val, nil
	}
	return "", errors.New("Locale " + key + " does not exist.")
}

func (this *Locale) LoadJson(filePath string) error {
	rawJson, err := ioutil.ReadFile(filePath)

	if err != nil {
		return err
	}

	var data map[string]string

	if err = json.Unmarshal(rawJson, &data); err != nil {
		return err
	}

	for key, value := range data {
		this.Set(key, value)
	}

	return nil
}
