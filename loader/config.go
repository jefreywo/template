package loader

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

const defaultConfigPath = "/config"
const (
	EnvDev   = "dev"
	EnvTest  = "test"
	EnvProd  = "prod"
	EnvLocal = "local"
)

var V *viper.Viper

func parseConfigFile(path string) error {
	var configPath string
	if len(strings.TrimSpace(path)) > 0 {
		configPath = path + defaultConfigPath
	} else {
		configPath = GetExecutablePath() + defaultConfigPath
	}

	V = viper.New()
	V.SetConfigFile(viper.GetString("env"))
	V.AddConfigPath(configPath)
	if err := V.ReadInConfig(); err != nil {
		log.Fatalln("加载配置文件失败:", err)
	}
	return nil
}

func GetExecutablePath() string {
	ex, err := os.Executable()
	if err != nil {
		log.Fatalln("获取可执行文件路径失败:", err)
	}
	return filepath.Dir(ex)
}
