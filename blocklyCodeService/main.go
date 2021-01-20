package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//r.Static("resource/demos/blockfactory", "./")
	//r.LoadHTMLGlob("resource/demos/index.html")
	r.LoadHTMLFiles("resource/demos/index.html")
	//r.LoadHTMLFiles("D:\\Firefly-Automation-Test\\blocklyCodeService\\resource\\demos\\index.html")
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "demos/index",
		})
	})

	err := r.Run(":8888")
	fmt.Print(err)
}
