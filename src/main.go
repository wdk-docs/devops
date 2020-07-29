package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}
		// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})
	r.LoadHTMLGlob("templates/*")
	r.GET("/index",func (c *gin.Context)  {
		c.HTML(http.StatusOK,"index.tmpl",gin.H{"title":"Main website"})
	})
	r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
