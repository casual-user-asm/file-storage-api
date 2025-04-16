package storage

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var minioClient *minio.Client
var ctx = context.Background()

func InitMinio() {
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKey := os.Getenv("MINIO_ROOT_USER")
	secretKey := os.Getenv("MINIO_ROOT_PASSWORD")
	bucketName := os.Getenv("MINIO_BUCKET")
	location := "us-east-1"
	var err error

	minioClient, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalf("Error to init minio client %v", err)
	}

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}
}

func UploadFile(objectName string, file io.Reader, fileSize int64, content string) error {
	bucketName := os.Getenv("MINIO_BUCKET")
	contentType := content
	info, err := minioClient.PutObject(ctx, bucketName, objectName, file, fileSize, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)

	return nil
}

func DownloadFile(objectName string) error {
	bucketName := os.Getenv("MINIO_BUCKET")
	_, err := minioClient.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Printf("Successfully download  %s", objectName)
	return nil
}

func ListAllFiles() ([]string, error) {
	bucketName := os.Getenv("MINIO_BUCKET")
	var fileNames []string

	allFiles := minioClient.ListObjects(ctx, bucketName, minio.ListObjectsOptions{
		Recursive: true,
	})

	for file := range allFiles {
		if file.Err != nil {
			return nil, file.Err
		}
		fileNames = append(fileNames, file.Key)
	}
	return fileNames, nil
}
