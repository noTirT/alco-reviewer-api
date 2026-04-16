package configs

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"noTirT/alcotracker/util"
	"os"
	"path"
	"strings"
	"time"
)

func CleanRSAKey(key string) string {
	returnKey := strings.ReplaceAll(key, "-----BEGIN PRIVATE KEY-----", "")
	returnKey = strings.ReplaceAll(returnKey, "-----END PRIVATE KEY-----", "")
	returnKey = strings.ReplaceAll(returnKey, " ", "")
	returnKey = strings.ReplaceAll(returnKey, "\n", "")
	return returnKey
}

type Configuration struct {
	ServerPort    string
	DBHost        string
	DBName        string
	DBUser        string
	DBPass        string
	DBPort        string
	JwtExpiration int
}

func NewConfiguration(configFile string) *Configuration {
	workDir, err := os.Getwd()

	if err != nil {
		panic("Working directory could not be read")
	}

	viper.SetConfigFile(configFile)
	err = viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	viper.SetDefault("SERVER_PORT", ":8080")
	viper.SetDefault("PG_HOST", "localhost")
	viper.SetDefault("PG_PORT", 5432)
	viper.SetDefault("PG_USER", "postgres")
	viper.SetDefault("PG_PASSWORD", "password")
	viper.SetDefault("PG_DATABASE", "alcohol-app")

	viper.SetDefault("JWT_EXPIRATION", 30)
	viper.SetDefault("PEM_KEY_EXPIRATION", 14)

	for _, tokenType := range []string{"access", "refresh"} {
		filePath := path.Join(workDir, fmt.Sprintf("/configs/%s-private.pem", tokenType))

		_, err := os.Stat(filePath)
		if errors.Is(err, os.ErrNotExist) {
			util.GenerateRSAKeyPairs(tokenType)
			continue
		}

		maxAge := time.Duration(viper.GetUint("PEM_KEY_EXPIRATION")) * 24 * time.Hour

		if util.IsKeyOld(tokenType, maxAge) {
			log.Printf("Regenerate %s-token due to old age\n", tokenType)
			util.GenerateRSAKeyPairs(tokenType)
			log.Printf("\n")
		}
	}

	config := &Configuration{
		ServerPort:    viper.GetString("SERVER_PORT"),
		DBHost:        viper.GetString("PG_HOST"),
		DBName:        viper.GetString("PG_DATABASE"),
		DBUser:        viper.GetString("PG_USER"),
		DBPass:        viper.GetString("PG_PASSWORD"),
		DBPort:        viper.GetString("PG_PASSWORD"),
		JwtExpiration: viper.GetInt("JWT_EXPIRATION"),
	}

	return config
}
