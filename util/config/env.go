package config

import (
	"os"
)

// GetPortEnv
func GetPortEnv() string {
	return os.Getenv("PORT")
}

// GetPGDataSourceEnv
func GetPGDataSourceEnv() string {
	return os.Getenv("PG_DATASOURCE")
}

func GetJWTSerectKeyEnv() string {
	return os.Getenv("JWT_SECRET_KEY")
}
