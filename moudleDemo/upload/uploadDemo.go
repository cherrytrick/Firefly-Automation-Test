package main

import (
	"github.com/gin-gonic/gin"
)

var (
	uploadFileKey = "upload-key"
)

func main() {
	r := gin.Default()
	r.POST("/upload", uploadHandler)
	r.Run()
}

// FormFile调用了ParseMultipartForm，已经分配了32MB的内存，超出32MB的部分以临时文件保存起来了
// 当SaveUploadedFile的时候，相当于打开临时文件，再写入自己的文件中，效率并不高
// 建议直接使用MultipartReader，自行处理各个Part。ParseMultipartForm本质上也是用的MultipartReader
func uploadHandler(c *gin.Context) {
	header, err := c.FormFile(uploadFileKey)
	if err != nil {
		//ignore
	}
	dst := header.Filename
	// gin 简单做了封装,拷贝了文件流
	if err := c.SaveUploadedFile(header, dst); err != nil {
		// ignore
	}
}
