package initialize

import (
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/wytfy/gredir/global"
	"log"
	"os"
)

// InitViper 初始化配置
func InitViper(path ...string) {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "specify the config file.")
		flag.Parse()
		if config == "" {
			log.Printf("using the default config.")
			config = "./config.yaml"
		} else {
			log.Printf("using the config file set by -c.\n")
		}
	} else {
		config = path[0]
		log.Printf("using the InitViper(...) config: %v", config)
	}
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("fail to read config file")
	}

	if err := v.Unmarshal(&global.CONF); err != nil {
		fmt.Println(err)
	}
}

// InitLogrus 初始化日志组件
func InitLogrus(path string) {
	global.LOGGER = logrus.New()
	global.LOGGER.SetLevel(logrus.InfoLevel)
	writer, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		global.LOGGER.SetOutput(os.Stdout)
		fmt.Println("using the stdout as the logger output.")
	}
	global.LOGGER.SetOutput(writer)
	fmt.Printf("Log redirect to %v.\n", path)
}
