package config

import (
	"fmt"
	"log"
	"net/http"
	"os"

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

func InitConfig() {
	homePath, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(homePath + "/.kongcli")
	err = os.MkdirAll(homePath+"/.kongcli", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

func ParseConfig() (conf *AppConfig, err error) {
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	servers := &AppConfig{}
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			conf, err = generateNew()
			if err != nil {
				return nil, err
			} else {
				return conf, nil
			}
		} else {
			return nil, err
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

	// Initialize the Servers map if it's nil
	if servers.Servers == nil {
		servers.Servers = make(map[string]*Server)
	}

	server := &Server{
		User:     user,
		Password: password,
		Url:      url,
	}
	servers.Servers[name] = server

	viper.Set("servers", servers.Servers)
	err = viper.WriteConfig()
	return err
}

func DeleteFromConfig(name string) error {
	config, err := ParseConfig()
	if err != nil {
		return err
	}

	delete(config.Servers, name)
	viper.Set("servers", config.Servers)
	return viper.WriteConfig()
}

func ListServers() error {
	config, err := ParseConfig()
	if err != nil {
		return err
	}

	for name, server := range config.Servers {
		fmt.Println("--------------------------------")
		fmt.Printf("Name: %s\n", name)
		if server.Url != "" {
			fmt.Printf("URL: %s\n", server.Url)
		}
		if server.User != "" {
			fmt.Printf("User: %s\n", server.User)
		}
	}

	return nil
}

func SetContext(name string) error {
	config, err := ParseConfig()
	if err != nil {
		return err
	}

	_, ok := config.Servers[name]
	if !ok {
		log.Fatalf("Server %s not found", name)
	}

	config.CurrentContext.Current = name
	viper.Set("context", config.CurrentContext)
	return viper.WriteConfig()
}

func generateNew() (*AppConfig, error) {
	servers := make(map[string]*Server)
	appConfig := &AppConfig{
		Servers: servers,
	}

	viper.Set("servers", appConfig.Servers)

	err := viper.SafeWriteConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileAlreadyExistsError); ok {
			return appConfig, nil
		} else {
			return nil, err
		}
	}
	return appConfig, nil
}

func LoadConfig(path string) (*AppConfig, error) {
	conf, err := ParseConfig()
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func CreateClient(url string) (*kong.Client, error) {
	httpClient := &http.Client{}
	client, err := kong.NewClient(&url, httpClient)

	if err != nil {
		return nil, err
	}

	return client, nil
}
