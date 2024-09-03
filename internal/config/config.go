package config

import (
	"fmt"
	"go.uber.org/fx"
	"time"

	"go.uber.org/config"
)

type Config struct {
	fx.Out

	Provider config.Provider
	//Http HttpConfig `yaml:"http"`
	//DB   DBConfig   `yaml:"postgres"`
}

func NewConfig() (Config, error) {
	loader, err := config.NewYAML(config.File("./configs/monolith.yml"))
	if err != nil {
		return Config{}, fmt.Errorf("fail load monolith config: %w", err)
	}
	return Config{Provider: loader}, nil
}

type HttpConfig struct {
	Port int `yaml:"port"`
}

func NewHttpConfig(provider config.Provider) (*HttpConfig, error) {
	var cfg HttpConfig
	if err := provider.Get("monolith").Get("http").Populate(&cfg); err != nil {
		return nil, fmt.Errorf("http config: %w", err)
	}
	return &cfg, nil
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DB       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`

	MaxOpenConns    int           `yaml:"maxOpenConns"`
	MaxIdleConns    int           `yaml:"maxIdleConns"`
	ConnMaxLifetime time.Duration `yaml:"connMaxLifetime"`
	ConnMaxIdleTime time.Duration `yaml:"connMaxIdleTime"`
}

func NewDBConfig(provider config.Provider) (*DBConfig, error) {
	var cfg DBConfig
	if err := provider.Get("monolith").Get("postgres").Populate(&cfg); err != nil {
		return nil, fmt.Errorf("postgres config: %w", err)
	}
	return &cfg, nil
}

func (c *DBConfig) URI() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.User, c.Password, c.Host, c.Port, c.DB)
}

//func LoadConfig() (Config, error) {
//	// TODO: configure
//	viper.AddConfigPath(".")
//	// viper.SetConfigName("app")
//	viper.SetConfigFile(".env")
//	viper.SetConfigType("env")
//
//	viper.AutomaticEnv()
//
//	if err := viper.ReadInConfig(); err != nil {
//		return Config{}, err
//	}
//
//	var config Config
//	err := viper.Unmarshal(&config)
//	return config, err
//}
