package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Ethereum EthereumConfig `mapstructure:"ethereum"`
}

// 服务器配置
type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

// 数据库配置
type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
	TimeZone string `mapstructure:"timezone"`
}

// Ethereum 配置
type EthereumConfig struct {
	WSURL         string `mapstructure:"ws_url"`
	RPCURL        string `mapstructure:"rpc_url"`
	ContractAddr  string `mapstructure:"contract_address"`
	StartBlock    uint64 `mapstructure:"start_block"`
	ScanBatchSize int    `mapstructure:"scan_batch_size"`
	RateLimit     int    `mapstructure:"rate_limit"`
	MaxRetries    int    `mapstructure:"max_retries"`
}

func init() {
	// 设置配置文件名称
	viper.SetConfigName("config")
	// 设置配置文件后缀
	viper.SetConfigType("yaml")
	// 设置配置文件路径
	viper.AddConfigPath("./config/")

	// 设置配置前缀
	viper.SetEnvPrefix("NFT_MARKETPLACE")
	viper.AutomaticEnv()

	// 默认配置
	viper.SetDefault("server.host", "0.0.0.0")
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("server.mode", "debug")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 5432)
	viper.SetDefault("database.sslmode", "disable")
	viper.SetDefault("database.timezone", "Asia/Shanghai")
	viper.SetDefault("ethereum.scan_batch_size", 2000)
	viper.SetDefault("ethereum.rate_limit", 500)
	viper.SetDefault("ethereum.max_retries", 3)
	viper.SetDefault("event.max_events", 100)

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: Error reading config file: %v", err)
		log.Println("Using default values and environment variables")
	} else {
		log.Printf("Using config file: %s", viper.ConfigFileUsed())
	}
}

func LoadConfig() (*Config, error) {
	var config Config

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
