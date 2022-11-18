package translater

import (
	"encoding/json"
	"errors"

	"github.com/go-resty/resty/v2"
)

var NaverClientId string
var NaverClientSecret string

func post(msg string, src string, dst string) (string, error) {

	reqUrl := NewReqUrl("").
		SetParam("source", src).
		SetParam("target", dst).
		SetParam("text", msg)

	client := resty.New()
	respJson, err := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8").
		SetHeader("X-Naver-Client-Id", NaverClientId).
		SetHeader("X-Naver-Client-Secret", NaverClientSecret).
		SetBody(reqUrl.Get()).
		Post("https://openapi.naver.com/v1/papago/n2mt")

	if err != nil {
		return "", err
	}

	var result NaverPapagoRespon

	err = json.Unmarshal(respJson.Body(), &result)
	if err != nil {
		return "", err
	}

	if result.Message.MsgType == "" {
		return "", errors.New(string(respJson.Body()))
	}

	return result.Message.Result.TranslatedText, nil
}

func Ko2En(msg string) (string, error) {
	return post(msg, "ko", "en")
}

func En2Ko(msg string) (string, error) {
	return post(msg, "en", "ko")
}
