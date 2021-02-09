package main

import (
	"log"
	"net/http"
	"registrationCenter/discover"
	"registrationCenter/register"
	"registrationCenter/synserver"
)

func main() {
	//初始化集群配置
	synserver.InitConfig("application.ini")
	//服务注册
	register.RegisterWeb()
	//同步服务
	synserver.SynserverWeb()
	//发现服务
	discover.DiscoverHttp()

	log.Fatal(http.ListenAndServe(synserver.LocalUrl, nil))
}
