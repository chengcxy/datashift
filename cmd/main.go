package main

import (
	"flag"
	"fmt"
	_ "github.com/chengcxy/datashift/clients"
	"github.com/chengcxy/datashift/scheduler"
	"github.com/chengcxy/gotools/configor"
)

var ConfigPath string
var Env string
var UsedEnv bool
var Config *configor.Config

func init() {
	flag.StringVar(&ConfigPath, "c", "./config", "配置文件路径")
	flag.StringVar(&Env, "e", "dev", "在什么环境运行")
	flag.BoolVar(&UsedEnv, "usedEnv", false, "是否使用环境变量")
	flag.Parse()
	Config = configor.NewConfig(ConfigPath, Env, UsedEnv)
	fmt.Println(ConfigPath, Env, UsedEnv)
	_, err := Config.Get("reader")
	fmt.Println(err)
}
func main() {
	s := &scheduler.Scheduler{
		Config: Config,
	}
	s.Run()
}
