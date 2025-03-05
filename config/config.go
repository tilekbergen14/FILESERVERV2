package config

import (
	"log"
	"os"
)

var (
	MinioEndpoint = os.Getenv("MINIO_ENDPOINT")
	MinioAccessKey = os.Getenv("MINIO_ACCESS_KEY")
	MinioSecretKey = os.Getenv("MINIO_SECRET_KEY")
	BucketName = os.Getenv("BUCKET_NAME")
)

func LoadConfig() {
	MinioEndpoint = os.Getenv("MINIO_ENDPOINT")
	MinioAccessKey = os.Getenv("MINIO_ACCESS_KEY")
	MinioSecretKey = os.Getenv("MINIO_SECRET_KEY")
	BucketName = os.Getenv("BUCKET_NAME")
	if MinioEndpoint == "" || MinioAccessKey == "" || MinioSecretKey == "" || BucketName == "" {
		log.Fatal("Error: Missing configuration values")
	}
}
