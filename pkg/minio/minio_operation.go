package minio

import (
	"context"
	"log"
	"mime/multipart"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"

	"github.com/jykuo-love-shiritori/twp/pkg/constants"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MC struct {
	mcp *minio.Client
}

func NewMINIO() *MC {

	endpoint := os.Getenv("MINIO_HOST") + ":" + os.Getenv("MINIO_API_PORT")
	accessKey := os.Getenv("MINIO_ACCESS_KEY")
	secretKey := os.Getenv("MINIO_SECRET_KEY")

	// Initialize minio client object.
	var err error
	mcp, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false,
	})
	mc := &MC{mcp: mcp}

	if err != nil {
		log.Fatal(err)
	}
	if err = mc.CheckBuckets(context.Background(), os.Getenv("MINIO_BUCKET_NAME")); err != nil {
		log.Fatal(err)
	}
	return mc
}

func (mc MC) CheckBuckets(ctx context.Context, bucketName string) error {
	err := mc.mcp.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: "ap-northeast-1"})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := mc.mcp.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("bucket '%s' already have\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	}
	return nil
}

func (mc MC) PutFile(ctx context.Context, logger *zap.SugaredLogger, file *multipart.FileHeader) (string, error) {

	object, err := file.Open()
	if err != nil {
		return "", err
	}
	id := uuid.New()
	fileSize := file.Size
	parts := strings.Split(file.Filename, ".")
	fileType := parts[len(parts)-1]
	newFileName := id.String() + "." + fileType
	logger.Info(newFileName)

	info, err := mc.mcp.PutObject(ctx, constants.BUCKETNAME, newFileName, object, fileSize, minio.PutObjectOptions{ContentType: "multipart/form-data"})
	if err != nil {
		logger.Error(err)
		return "", err
	}
	if err := object.Close(); err != nil {
		logger.Error(err)
		return "", err
	}
	logger.Info(info)
	return newFileName, nil
}

func (mc MC) GetFileURL(c echo.Context, logger *zap.SugaredLogger, id string) string {
	reqParams := make(url.Values)
	presignedURL, err := mc.mcp.PresignedGetObject(c.Request().Context(), constants.BUCKETNAME, id, time.Second*24*60*60, reqParams)
	if err != nil {
		logger.Error(err)
		//default image if can find image by uuid
		return "https://imgur.com/UniMfif.png"
	}
	return presignedURL.String()
}
