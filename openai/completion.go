package openai

import (
	"encoding/json"
	"strings"

	"github.com/go-resty/resty/v2"
)

var URI = "https://api.openai.com/v1/completions"

const ME = "You:"
const AI = "Robot:"

var ApiKey string
var Param RequestBody

func init() {
	stop := []string{ME}
	Param.Stop = stop
}

func Chat(msg string) (string, error) {

	Param.Prompt = msg

	buf, err := json.Marshal(Param)
	if err != nil {
		return "", err
	}

	//fmt.Println("req: " + string(buf))

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

	//fmt.Println("res: " + string(respJson.Body()))

	err = json.Unmarshal(respJson.Body(), &result)
	if err != nil {
		return "", err
	}

	if len(result.Choices) == 0 {
		return string(respJson.Body()), nil
	}

	return strings.Trim(result.Choices[0].Text, "\n"), nil
}

func StripPrefix(text string) string {
	i := strings.LastIndex(text, ":")

	if i >= 0 {
		text = strings.Trim(text[i+1:], " ")
	}

	return text
}
