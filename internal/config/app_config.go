package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// 当前App的配置文件配置
const (
	LocalServiceType    = ServiceTypeDev
	LocalConfigDir      = FileDirDefault
	LocalFileNameLayout = FileNameLayoutDefault
	LocalConfigType     = FilePostfixDefault
)

const (
	ServiceTypeLabel = "SERVICE_TYPE"
	FileDirLabel     = "CONFIG_DIR"
	FilePostfixLabel = "CONFIG_POSTFIX"
)

const (
	ServiceTypeDev  = "dev"
	ServiceTypeProd = "prod"
)

const (
	FileDirDefault = "configs/"
)

const (
	FileNameLayoutDefault = "config.%s"
)

const (
	FilePostfixDefault = "yaml"
)

type AppConfig struct {
	Blockchain BlockchainConfig `mapstructure:"blockchain" json:"blockchain"`
}

var globalAppConfig *AppConfig

func LoadAppConfig() (*AppConfig, error) {
	// 设置 viper 的文件名、文件类型以及文件地址
	viper.SetConfigName(fmt.Sprintf(LocalFileNameLayout, LocalServiceType))
	viper.SetConfigType(LocalConfigType)
	viper.AddConfigPath(LocalConfigDir)

	// 自动扫描环境
	viper.AutomaticEnv()

	// 将配置文件读取到环境
	err := viper.ReadInConfig()

	if err != nil {
		log.Println("Read Config file error: ", err)
		return nil, err
	}

	var cfg *AppConfig

	err = viper.Unmarshal(&cfg)

	if err != nil {
		log.Println("Unmarshal Config file error: ", err)
		return nil, err
	}

	globalAppConfig = cfg

	return globalAppConfig, nil
}
