package config

import (
	"log"
	"os"
)

var (
	MinioEndpoint = os.Getenv("MINIO_ENDPOINT")
	MinioAccessKey = os.Getenv("MINIO_ACCESS_KEY")
	MinioSecretKey = os.Getenv("MINIO_SECRET_KEY")
)

func LoadConfig() {
	MinioEndpoint = os.Getenv("MINIO_ENDPOINT")
	MinioAccessKey = os.Getenv("MINIO_ACCESS_KEY")
	MinioSecretKey = os.Getenv("MINIO_SECRET_KEY")
	if MinioEndpoint == "" || MinioAccessKey == "" || MinioSecretKey == "" {
		log.Fatal("Error: Missing configuration values")
	}
}
