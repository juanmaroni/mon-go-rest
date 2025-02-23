package config

import (
	"fmt"
	"log/slog"
	"os"
)

var Logger *slog.Logger

type ServerConfig struct {
	Uri string
}

type MongoDbConfig struct {
	Uri string
	//Path string
}

// Environment variables
var (
	Server *ServerConfig
	MongoDb *MongoDbConfig
)

func LoadJSONLogger() {
	Logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

func LoadServerConfig() {
	Server = &ServerConfig{
		Uri: getEnvVar("SERVER_URI"),
	}
}

func LoadMongoDbConfig() {
	MongoDb = &MongoDbConfig{
		Uri: getEnvVar("MONGODB_URI"),
		//Path: getEnvVar("MONGODB_PATH"),
	}
}

func getEnvVar(varName string) string {
	value, exists := os.LookupEnv(varName)

	if !exists {
		msg := fmt.Sprintf("Environment variable '%s' is not set.", varName)
		Logger.Error(msg)
		panic(msg)
	}

	return value
}
