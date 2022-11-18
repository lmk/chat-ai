package main

import (
	"os"

	"test-api/translater"
)

var conf AppConfig

func main() {

	InitLogger(os.Stdout, os.Stdout, os.Stdout, os.Stderr)

	err := conf.readConfig("app.yaml")
	if err != nil {
		Error.Fatalln(err)
	}

	translater.NaverClientId = conf.NaverClientId
	translater.NaverClientSecret = conf.NaverClientSecret

	result, err := translater.Ko2En("안녕하세요")
	if err != nil {
		Error.Fatalln(err)
	}

	Info.Println(result)

}
