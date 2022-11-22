package main

import (
	"os"

	"chat-ai/openai"
	"chat-ai/translater"
)

var conf AppConfig

func InitConfig() {
	err := conf.readConfig("app.yaml")
	if err != nil {
		Error.Fatalln(err)
	}
	Info.Println(conf.makePretty())

	translater.NaverClientId = conf.NaverClientId
	translater.NaverClientSecret = conf.NaverClientSecret

	openai.ApiKey = conf.OpenAI.ApiKey
	openai.Param.Model = conf.OpenAI.Model
	openai.Param.MaxToken = conf.OpenAI.MaxTokens
	openai.Param.Temperature = conf.OpenAI.Temperature
	openai.Param.TopP = conf.OpenAI.TopP
	openai.Param.FrequencyPenalty = conf.OpenAI.FrequencyPenalty
	openai.Param.PresencePenalty = conf.OpenAI.PresencePenalty
}

func main() {

	InitLogger(os.Stdout, os.Stdout, os.Stdout, os.Stderr)

	Info.Println("Start App")

	InitConfig()

	StartWeb()

}
