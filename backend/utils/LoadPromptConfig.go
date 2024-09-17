package utils

import (
	"encoding/json"
	"io/ioutil"
)

type PromptConfig struct {
	Prompt string `json:"prompt"`
}

func LoadPromptConfig(filePath string) (string, error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	var config PromptConfig
	err = json.Unmarshal(file, &config)
	if err != nil {
		return "", err
	}

	return config.Prompt, nil
}
