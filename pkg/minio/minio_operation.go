package minio

import (
	"context"
	"errors"
	"mime/multipart"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"

	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MC struct {
	mcp        *minio.Client
	BucketName string
}

func NewMINIO() (*MC, error) {

	endpoint := os.Getenv("MINIO_HOST") + ":" + os.Getenv("MINIO_API_PORT")
	accessKey := os.Getenv("MINIO_ACCESS_KEY")
	secretKey := os.Getenv("MINIO_SECRET_KEY")

	// Initialize minio client object.
	var err error
	mcp, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false,
	})

	if err != nil {
		return nil, err
	}
	mc := &MC{mcp: mcp, BucketName: os.Getenv("MINIO_BUCKET_NAME")}

	if err = mc.CheckBucket(context.Background()); err != nil {
		return nil, err
	}
	return mc, nil
}

func (mc MC) CheckBucket(ctx context.Context) error {
	err := mc.mcp.MakeBucket(ctx, mc.BucketName, minio.MakeBucketOptions{Region: "ap-northeast-1"})

	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := mc.mcp.BucketExists(ctx, mc.BucketName)
		if errBucketExists != nil || !exists {
			return errors.New("fail to check bucket exists")
		}
	}
	return nil
}

func (mc MC) PutFile(ctx context.Context, file *multipart.FileHeader) (string, error) {

	object, err := file.Open()
	if err != nil {
		return "", err
	}
	id := uuid.New()
	fileSize := file.Size
	parts := strings.Split(file.Filename, ".")
	fileType := parts[len(parts)-1]
	newFileName := id.String() + "." + fileType

	info, err := mc.mcp.PutObject(ctx, mc.BucketName, newFileName, object, fileSize, minio.PutObjectOptions{ContentType: "multipart/form-data"})
	if err != nil {
		return info.Bucket, err
	}
	if err := object.Close(); err != nil {
		return info.Bucket, err
	}
	return newFileName, nil
}

func (mc MC) GetFileURL(ctx context.Context, id string) string {
	reqParams := make(url.Values)
	presignedURL, err := mc.mcp.PresignedGetObject(ctx, mc.BucketName, id, time.Second*24*60*60, reqParams)
	if err != nil {
		//default image if can find image by uuid
		return ""
	}
	return presignedURL.String()
}
