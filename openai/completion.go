package openai

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

var URI = "https://api.openai.com/v1/completions"

var ApiKey string
var Param RequestBody

func Chat(msg string) (string, error) {

	Param.Prompt = msg
	Param.Temperature = 0.9
	Param.MaxToken = 1024

	buf, err := json.Marshal(Param)
	if err != nil {
		return "", err
	}

	client := resty.New()
	respJson, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+ApiKey).
		SetBody(string(buf)).
		Post(URI)

	if err != nil {
		return "", err
	}

	var result ResponBody

	fmt.Println(string(respJson.Body()))

	err = json.Unmarshal(respJson.Body(), &result)
	if err != nil {
		return "", err
	}

	if len(result.Choices) == 0 {
		return string(respJson.Body()), nil
	}

	text := strings.Trim(result.Choices[0].Text, "\n")
	text = strings.ReplaceAll(text, "\n", " ")
	return text, nil
}
