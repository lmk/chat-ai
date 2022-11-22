package main

import (
	"chat-ai/openai"
	"chat-ai/translater"
	"fmt"
	"strings"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func StartWeb() {
	r := gin.Default()
	r.POST("/api/v1/chat", func(c *gin.Context) {

		text := c.PostForm("text")
		reqKrMsg := c.PostForm("say")

		reqEnMsg, err := translater.Ko2En(reqKrMsg)
		if err != nil {
			Error.Println("Ko2En", err)
			c.JSON(500, gin.H{
				"text": "fail translat '" + reqEnMsg + "'",
			})
			return
		}

		text += openai.ME + reqEnMsg + "\n"

		resEnMsg, err := openai.Chat(reqEnMsg)
		if err != nil {
			Error.Println("Chat", err)
			c.JSON(500, gin.H{
				"text": "fail chat '" + reqEnMsg + "'",
			})
			return
		}

		resEnMsg = strings.ReplaceAll(strings.Trim(resEnMsg, "\n"), "\n", " ")
		text += openai.AI + resEnMsg + "\n"

		resKrMsg, err := translater.En2Ko(resEnMsg)
		if err != nil {
			Error.Println("En2Ko", err)
			c.JSON(500, gin.H{
				"text": "fail translat '" + resEnMsg + "'",
			})
			return
		}

		Info.Println(text)

		c.JSON(200, gin.H{
			"text":   text,
			"req-en": reqEnMsg,
			"req-kr": reqKrMsg,
			"res-en": resEnMsg,
			"res-kr": resKrMsg,
		})
	})
	r.Use(static.Serve("/", static.LocalFile("public", true)))

	r.Run(fmt.Sprintf(":%d", conf.ListenPort))
}
