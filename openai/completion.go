package openai

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

var URI = "https://api.openai.com/v1/completions"

const ME = " Human:"
const AI = " AI:"

var ApiKey string
var Param RequestBody

func init() {
	stop := []string{ME, AI}
	Param.Stop = stop
}

func Chat(msg string) (string, error) {

	Param.Prompt = msg

	//fmt.Println(Param)

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

	//fmt.Println(string(respJson.Body()))

	err = json.Unmarshal(respJson.Body(), &result)
	if err != nil {
		return "", err
	}

	if len(result.Choices) == 0 {
		return string(respJson.Body()), nil
	}

	return result.Choices[0].Text, nil
}
