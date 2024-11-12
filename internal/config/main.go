package config

import (
	"net/http"

	"github.com/kong/go-kong/kong"
	"github.com/spf13/viper"
)

type AppConfig struct {
	CurrentContext ServerContext      `mapstructure:"context"`
	Servers        map[string]*Server `mapstructure:"servers"`
}

type ServerContext struct {
	Current string `mapstructure:"current"`
}

type Server struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"pass"`
	Url      string `mapstructure:"url"`
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.kongcli")
	viper.AddConfigPath(".")
}

func ParseConfig() (conf *AppConfig, err error) {
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	servers := &AppConfig{}
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return servers, err
		} else {
			return servers, err
		}
	}
	err = viper.Unmarshal(servers)
	return servers, err
}

func AddToConfig(name string, url string, user string, password string) error {
	// servers := make(map[string]*Server)
	servers, err := ParseConfig()

	if err != nil {
		return err
	}

	server := &Server{
		User:     user,
		Password: password,
		Url:      url,
	}
	servers.Servers[name] = server

	return nil

}

func GenerateNew() {
	servers := make(map[string]*Server)

}

func LoadConfig(path string) {}

func CreateClient(url string) (*kong.Client, error) {
	httpClient := &http.Client{}
	client, err := kong.NewClient(&url, httpClient)

	if err != nil {
		return nil, err
	}

	return client, nil
}
