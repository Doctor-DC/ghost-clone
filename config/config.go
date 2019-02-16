package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

// InitConfig initializes configuration from toml file from the config directory and
// sets some important variable for later use. It uses viper for that
func InitConfig() {
	if len(os.Args) > 1 && os.Args[1] == "development" || len(os.Args) == 1 {
		viper.AddConfigPath("config")
		viper.SetConfigName("config.development")
	} else if len(os.Args) > 1 && os.Args[1] == "production" {
		viper.AddConfigPath("/go/bin/config")
		viper.SetConfigName("config.production")
	} else if len(os.Args) > 2 && os.Args[2] == "testing" {
		viper.AddConfigPath("config")
		viper.SetConfigName("config.testing")
	}

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	PRIVATE_KEY = viper.GetString("server.keys_dir") + "/app.rsa"
	PUBLIC_KEY = viper.GetString("server.keys_dir") + "/app.rsa.pub"
	DB_NAME = viper.GetString("database.name")
	ROOT_DIR, _ = os.Getwd()
}

var (
	DB_NAME                = "temp"
	POST_COLLECTION_NAME   = "posts"
	TAG_COLLECTION_NAME    = "tags"
	AUTHOR_COLLECTION_NAME = "authors"
	PRIVATE_KEY            = os.Getenv("KEYS_DIR") + "/app.rsa"
	PUBLIC_KEY             = os.Getenv("KEYS_DIR") + "/app.rsa.pub"
	DOMAIN                 = "http://api.bitneni.com"
	VERSION                = "v1"
	CACHE_DURATION         = "5"
	ROOT_DIR               = ""
)
