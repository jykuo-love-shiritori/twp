package minio

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"net/url"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v7"

	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MC *minio.Client

func init() {

	endpoint := "localhost:" + os.Getenv("MINIO_API_PORT")
	accessKey := os.Getenv("MINIO_ACCESS_KEY")
	secretKey := os.Getenv("MINIO_SECRET_KEY")

	log.Printf("Endpoint: %s", endpoint)
	log.Printf("AccessKeyID: %s", accessKey)
	log.Printf("SecretAccessKey: %s", secretKey)
	// Initialize minio client object.
	var err error
	MC, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", MC) // minioClient is now setup
}

func CheckBuckets(ctx context.Context) error {
	fmt.Println("checkout buckets")
	buckets, err := MC.ListBuckets(ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, bucket := range buckets {
		fmt.Println(bucket)
	}
	return nil
}

func Putfile(ctx context.Context, src multipart.File, file_size int64, file_name string) error {
	fmt.Println("push file")
	_, err := MC.PutObject(ctx, "files", file_name, src, file_size, minio.PutObjectOptions{ContentType: "multipart/form-data"})
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func GeneratePresignedURL(ctx context.Context, file_name string) (string, error) {
	fmt.Println("gen url")
	reqParams := make(url.Values)
	presignedURL, err := MC.PresignedGetObject(ctx, "files", file_name, time.Second*24*60*60, reqParams)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return presignedURL.String(), nil
}

func GetFileURL(c echo.Context, file_name string) string {
	url, err := GeneratePresignedURL(c.Request().Context(), file_name)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return url
}
