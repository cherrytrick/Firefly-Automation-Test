package route

import (
	"github.com/gin-gonic/gin"
	"worker/controller"
	"worker/route/middlewares"
)

func InitRouter() {
	router := gin.Default()
	// 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
	router.Use(middlewares.Cors())
	// 使用 session(cookie-based)
	//router.Use(sessions.Sessions("myyyyysession", session.Store))
	v1 := router.Group("v1")
	{
		v1.POST("/exec", controller.ExecScript)
	}

	router.Run(":8888")
}
