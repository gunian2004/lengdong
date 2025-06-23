package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

// Config 存储所有应用程序的配置
type Config struct {
	Server     ServerConfig     `mapstructure:"server"`
	Database   DatabaseConfig   `mapstructure:"database"`
	Redis      RedisConfig      `mapstructure:"redis"`
	JWT        JWTConfig        `mapstructure:"jwt"`
	Blockchain BlockchainConfig `mapstructure:"blockchain"`
	Ethereum   EthereumConfig   `mapstructure:"ethereum"`
}

// BlockchainConfig 存储区块链相关的配置
type BlockchainConfig struct {
	NetworkConfigPath string `mapstructure:"network_config_path"`
	ChannelID         string `mapstructure:"channel_id"`
	ChaincodeName     string `mapstructure:"chaincode_name"`
	OrgAdmin          string `mapstructure:"org_admin"`
	OrgName           string `mapstructure:"org_name"`
}

// EthereumConfig 存储以太坊相关的配置
type EthereumConfig struct {
	NodeURL         string `mapstructure:"node_url"`
	ContractAddress string `mapstructure:"contract_address"`
	PrivateKey      string `mapstructure:"private_key"`
}

var AppConfig Config

// LoadConfig 从配置文件中加载配置
func LoadConfig(configPath string) {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(configPath)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件失败: %s", err)
	}

	err := viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatalf("无法解析配置到结构体: %v", err)
	}
}
