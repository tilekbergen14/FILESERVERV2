package config

import (
	"log"
	"os"
)

var (
	MinioEndpoint = os.Getenv("MINIO_ENDPOINT")
	MinioAccessKey = os.Getenv("MINIO_ACCESS_KEY")
	MinioSecretKey = os.Getenv("MINIO_SECRET_KEY")
	MinioBucket = os.Getenv("MINIO_BUCKET")
)

func LoadConfig() {
	MinioEndpoint = os.Getenv("MINIO_ENDPOINT")
	MinioAccessKey = os.Getenv("MINIO_ACCESS_KEY")
	MinioSecretKey = os.Getenv("MINIO_SECRET_KEY")
	MinioBucket = os.Getenv("MINIO_BUCKET")
	if MinioEndpoint == "" || MinioAccessKey == "" || MinioSecretKey == "" || MinioBucket == ""{
		log.Fatal("Error: Missing configuration values")
	}
}
