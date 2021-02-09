package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"worker/common"
	"worker/service"
)

func TestInsert(c *gin.Context) {
	var testService service.Test
	err := c.ShouldBindJSON(&testService)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := testService.Insert()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Insert() error!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "success",
		"data":    id,
	})

}

func ExecScript(c *gin.Context) {
	var scriptType service.Script
	err := c.ShouldBindJSON(&scriptType)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	name := scriptType.Name
	common.PutMessage(name)

	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "success",
		"data":    "",
	})

}
