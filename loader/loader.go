package loader

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

func LoadConfig(path string) {
	viper.SetEnvPrefix("service_runs_in")
	viper.BindEnv("env")

	if len(strings.TrimSpace(viper.GetString("env"))) <= 0 {
		log.Fatalln("没有设定[SERVICE_RUNS_IN_ENV]环境变量的值，请先使用export SERVICE_RUNS_IN_ENV = dev/test/prod设置环境变量 ")
	}

	log.Println("服务运行环境为:", viper.GetString("env"))

	if err := parseConfigFile(path); err != nil {
		log.Fatalln("config文件解析错误:", err)
	}
}
