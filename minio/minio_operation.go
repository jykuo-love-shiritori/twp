package minio

import (
	"context"
	"mime/multipart"
	"net/url"
	"os"
	"time"

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

	return mc, nil
}

func (mc MC) PutFile(ctx context.Context, file *multipart.FileHeader, fileName string) (string, error) {

	object, err := file.Open()
	if err != nil {
		return "", err
	}

	info, err := mc.mcp.PutObject(ctx, mc.BucketName, fileName, object, file.Size, minio.PutObjectOptions{ContentType: "multipart/form-data"})
	if err != nil {
		return info.VersionID, err
	}
	if err := object.Close(); err != nil {
		return info.VersionID, err
	}
	return fileName, nil
}
func (mc MC) PutFileByPath(ctx context.Context, path string, fileName string) (string, error) {

	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	fileStat, err := file.Stat()
	if err != nil {
		return "", err
	}

	info, err := mc.mcp.PutObject(ctx, mc.BucketName, fileName, file, fileStat.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return info.VersionID, err
	}
	if err := file.Close(); err != nil {
		return info.VersionID, err
	}
	return fileName, nil
}

func (mc MC) RemoveFile(ctx context.Context, fileName string) error {
	//get object version
	objInfo, err := mc.mcp.StatObject(context.Background(), mc.BucketName, fileName, minio.StatObjectOptions{})
	if err != nil {
		return nil
	}
	opts := minio.RemoveObjectOptions{
		GovernanceBypass: true,
		VersionID:        objInfo.VersionID,
	}
	err = mc.mcp.RemoveObject(context.Background(), mc.BucketName, fileName, opts)
	if err != nil {
		return err
	}
	return nil
}
func (mc MC) GetFileURL(ctx context.Context, fileName string) string {
	reqParams := make(url.Values)
	presignedURL, err := mc.mcp.PresignedGetObject(ctx, mc.BucketName, fileName, time.Second*24*60*60, reqParams)
	if err != nil {
		//default image if can find image by uuid
		return ""
	}
	return presignedURL.String()
}
