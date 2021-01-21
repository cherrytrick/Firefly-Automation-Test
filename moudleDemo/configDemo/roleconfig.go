package main

import (
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/logx"
)

type RoleConfig struct {
	//自动化测试
	AutomationRole bool `json:"automation_role,default=false"`
	//压力测试
	StressRole bool `json:"stress_role,default=false"`
	//接口测试
	ApiRole bool `json:"api_role,default=true"`
}

func main() {
	th := &RoleConfig{}
	err := conf.LoadConfig("./etc/worker-role.yaml", th)
	if err != nil {
		logx.Error(err)
	}
	logx.Infof("%+v", th)
}
