package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	ListenPort        int    `yaml:"listenPort"`
	NaverClientId     string `yaml:"naverClientId"`
	NaverClientSecret string `yaml:"naverClientSecret"`
	OpenAI            openAI `yaml:"openai"`
}

type openAI struct {
	ApiKey           string  `yaml:"apiKey"`
	Model            string  `yaml:"model"`
	MaxTokens        int     `yaml:"max_tokens"`
	Temperature      float32 `yaml:"temperature"`
	TopP             float32 `yaml:"top_p"`
	FrequencyPenalty float32 `yaml:"frequency_penalty"`
	PresencePenalty  float32 `yaml:"presence_penalty"`
}

func (conf *AppConfig) readConfig(fileName string) error {

	Info.Println("Read " + fileName)

	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("cannot read config file %s, ReadFile: %v", fileName, err)
	}

	err = yaml.Unmarshal(buf, conf)
	if err != nil {
		return fmt.Errorf("invaild config file %s, Unmarshal: %v", fileName, err)
	}

	return nil
}

func (conf *AppConfig) makePretty() string {

	buf, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		Info.Fatalf(err.Error())
	}

	return string(buf)
}
