package main

import (
	"test-api/openai"
	"test-api/translater"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func StartWeb() {
	r := gin.Default()
	r.POST("/api/v1/chat", func(c *gin.Context) {

		text := c.PostForm("text")

		text, err := translater.Ko2En(text)
		if err != nil {
			c.JSON(500, gin.H{
				"text": "fail translat '" + text + "'",
			})
			return
		}

		resEn, err := openai.Chat(text)
		if err != nil {
			c.JSON(500, gin.H{
				"text": "fail chat '" + text + "'",
			})
			return
		}

		resKr, err := translater.En2Ko(resEn)
		if err != nil {
			c.JSON(500, gin.H{
				"text": "fail translat '" + resEn + "'",
			})
			return
		}

		c.JSON(200, gin.H{
			"message-en": resEn,
			"message-kr": resKr,
		})
	})
	r.Use(static.Serve("/", static.LocalFile("public", true)))

	r.Run(":8888")
}
