package services

import (
	"context"
	"fileserver/config"
	"fmt"
	"io"
	"log"
	"path/filepath"

	"github.com/google/uuid"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var mioClient *minio.Client

func InitMinio() {
	client, err := minio.New(config.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.MinioAccessKey, config.MinioSecretKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalf("Error initializing Minio client: %v", err)
	}
	mioClient = client
	
}

func UploadFile(subfolder string, customUuid string, fileContent io.Reader, originalFilename string) (string, error) {
	// exists, err := mioClient.BucketExists(context.Background(), bucketName)
	// if err != nil {
	// 	return fmt.Errorf("unable to check if bucket exists: %w", err)
	// }

	// if !exists {
	// 	err := mioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
	// 	if err != nil {
	// 		return fmt.Errorf("unable to create bucket: %w", err)
	// 	}
	// 	fmt.Println("Bucket created successfully!")
	// }
	if customUuid == "" {
		customUuid = uuid.New().String()
	}
	
	extension := filepath.Ext(originalFilename)
	if extension == "" {
		extension = ".bin"
	}

	filePath := customUuid + extension
	if subfolder != ""{
		filePath = subfolder + "/" + filePath
	}
	_, err := mioClient.PutObject(
		context.Background(), 
		config.BucketName,     
		filePath,     
		fileContent,    
		-1, 
		minio.PutObjectOptions{
			ContentType: "application/octet-stream", 
		},
	)
	if err != nil {
		return "", fmt.Errorf("unable to upload file: %w", err)
	}

	return filePath, nil
}

func GetFile(bucketName, objectName string) ([]byte, error) {
	object, err := mioClient.GetObject(context.Background(), bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer object.Close()
	file, err := io.ReadAll(object)
	return file, err
}

func ServeFile(bucketName, objectName string) ([]byte, string, error) {
	object, err := mioClient.GetObject(context.Background(), bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, "", fmt.Errorf("unable to retrieve file from MinIO: %w", err)
	}
	defer object.Close()

	file, err := io.ReadAll(object)
	if err != nil {
		return nil, "", fmt.Errorf("unable to read file content: %w", err)
	}

	return file, "application/octet-stream", nil
}