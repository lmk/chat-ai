package main

import (
	"chat-ai/openai"
	"chat-ai/translater"
	"fmt"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func StartWeb() {
	r := gin.Default()
	r.POST("/api/v1/chat", func(c *gin.Context) {

		reqKrMsg := c.PostForm("text")

		reqEnMsg, err := translater.Ko2En(reqKrMsg)
		if err != nil {
			c.JSON(500, gin.H{
				"text": "fail translat '" + reqEnMsg + "'",
			})
			return
		}

		resEnMsg, err := openai.Chat(reqEnMsg)
		if err != nil {
			c.JSON(500, gin.H{
				"text": "fail chat '" + reqEnMsg + "'",
			})
			return
		}

		resKrMsg, err := translater.En2Ko(resEnMsg)
		if err != nil {
			c.JSON(500, gin.H{
				"text": "fail translat '" + resEnMsg + "'",
			})
			return
		}

		c.JSON(200, gin.H{
			"req-message-en": reqEnMsg,
			"req-message-kr": reqKrMsg,
			"res-message-en": resEnMsg,
			"res-message-kr": resKrMsg,
		})
	})
	r.Use(static.Serve("/", static.LocalFile("public", true)))

	r.Run(fmt.Sprintf(":%d", conf.ListenPort))
}
