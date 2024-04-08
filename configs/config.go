package configs

import (
	"strings"

	"github.com/spf13/viper"
)

func CleanRSAKey(key string) string {
	returnKey := strings.ReplaceAll(key, "-----BEGIN PRIVATE KEY-----", "")
	returnKey = strings.ReplaceAll(returnKey, "-----END PRIVATE KEY-----", "")
	returnKey = strings.ReplaceAll(returnKey, " ", "")
	returnKey = strings.ReplaceAll(returnKey, "\n", "")
	return returnKey
}

type Configuration struct {
	ServerPort             string
	DBHost                 string
	DBName                 string
	DBUser                 string
	DBPass                 string
	DBPort                 string
	AccessTokenPrivateKey  string
	AccessTokenPublicKey   string
	RefreshTokenPrivateKey string
	RefreshTokenPublicKey  string
	JwtExpiration          int
}

func NewConfiguration() *Configuration {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

    viper.SetDefault("SERVER_PORT", ":8080")
	viper.SetDefault("PG_HOST", "localhost")
	viper.SetDefault("PG_PORT", 5432)
	viper.SetDefault("PG_USER", "postgres")
	viper.SetDefault("PG_PASSWORD", "password")
	viper.SetDefault("PG_DATABASE", "alcohol-app")

	viper.SetDefault("ACCESS_PRIVATE", "")
	viper.SetDefault("ACCESS_PUBLIC", "")
	viper.SetDefault("REFRESH_PRIVATE", "")
	viper.SetDefault("REFRESH_PUBLIC", "")

	viper.SetDefault("JWT_EXPIRATION", 30)

	accessPrivate := CleanRSAKey(viper.GetString("ACCESS_PRIVATE"))
	accessPublic := CleanRSAKey(viper.GetString("ACCESS_PUBLIC"))
	refreshPrivate := CleanRSAKey(viper.GetString("REFRESH_PRIVATE"))
	refreshPublic := CleanRSAKey(viper.GetString("REFRESH_PUBLIC"))

	config := &Configuration{
		ServerPort:             viper.GetString("SERVER_PORT"),
		DBHost:                 viper.GetString("PG_HOST"),
		DBName:                 viper.GetString("PG_DATABASE"),
		DBUser:                 viper.GetString("PG_USER"),
		DBPass:                 viper.GetString("PG_PASSWORD"),
		DBPort:                 viper.GetString("PG_PASSWORD"),
		AccessTokenPrivateKey:  accessPrivate,
		AccessTokenPublicKey:   accessPublic,
		RefreshTokenPrivateKey: refreshPrivate,
		RefreshTokenPublicKey:  refreshPublic,
		JwtExpiration:          viper.GetInt("JWT_EXPIRATION"),
	}

	return config
}
