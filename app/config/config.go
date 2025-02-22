package config

import (
	"fmt"
	"os"
)

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
		panic(fmt.Sprintf("Environment variable '%s' is not set.", varName)) // Log (slog.Panic)
	}

	return value
}
