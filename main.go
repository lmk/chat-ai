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
	// openai.Param.Stream = conf.OpenAI.Stream
	// openai.Param.Logprobs = conf.OpenAI.Logprobs
	// openai.Param.Stop = conf.OpenAI.Stop

}

// func testChat() {
// 	reader := bufio.NewReader(os.Stdin)
// 	for {
// 		fmt.Print("Enter text: ")
// 		text, _ := reader.ReadString('\n')
// 		text = strings.Trim(text, " \n\t")
// 		if len(text) == 0 {
// 			break
// 		}

// 		text, err := translater.Ko2En(text)
// 		if err != nil {
// 			Error.Fatalln(err)
// 		}

// 		response, err := openai.Chat(text)
// 		if err != nil {
// 			Error.Fatalln(err)
// 		}

// 		text, err = translater.En2Ko(response)
// 		if err != nil {
// 			Error.Fatalln(err)
// 		}

// 		fmt.Printf("AI: %s (%s)\n", text, response)

// 	}
// }

func main() {

	InitLogger(os.Stdout, os.Stdout, os.Stdout, os.Stderr)

	Info.Println("Start App")

	InitConfig()

	StartWeb()

}
