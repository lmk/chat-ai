package main

import (
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
)

var conf AppConfig

func main() {

	InitLogger(os.Stdout, os.Stdout, os.Stdout, os.Stderr)

	err := conf.readConfig("app.yaml")
	if err != nil {
		Error.Fatalln(err)
	}

	inputMsg := "어떻게 지내?"

	reqUrl := NewReqUrl("").
		SetParam("source", "ko").
		SetParam("target", "en").
		SetParam("text", inputMsg)

	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8").
		SetHeader("X-Naver-Client-Id", conf.NaverClientId).
		SetHeader("X-Naver-Client-Secret", conf.NaverClientSecret).
		SetBody(reqUrl.Get()).
		Post("https://openapi.naver.com/v1/papago/n2mt")

	if err != nil {
		Error.Fatalf("Get Fail :%s", err)
	}

	fmt.Println(resp)
}
