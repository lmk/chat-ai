package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	NaverClientId     string `yaml:"naverClientId"`
	NaverClientSecret string `yaml:"naverClientSecret"`
	SrcLangType       string `yaml:"srcLangType"`
	TarLangType       string `yaml:"tarLangType"`
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

func makePretty(conf *AppConfig) string {

	buf, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		Info.Fatalf(err.Error())
	}

	return string(buf)
}
