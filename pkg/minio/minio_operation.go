package minio

import (
	"context"
	"log"
	"mime/multipart"
	"net/url"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"

	"github.com/jykuo-love-shiritori/twp/pkg/constants"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MC *minio.Client

func init() {

	endpoint := os.Getenv("MINIO_HOST") + ":" + os.Getenv("MINIO_API_PORT")
	accessKey := os.Getenv("MINIO_ACCESS_KEY")
	secretKey := os.Getenv("MINIO_SECRET_KEY")

	// Initialize minio client object.
	var err error
	MC, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalln(err)
	}
	if err = CheckBuckets(context.Background(), os.Getenv("MINIO_BUCKET_NAME")); err != nil {
		log.Fatalln(err)
	}

}

func CheckBuckets(ctx context.Context, bucketName string) error {
	err := MC.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: "sa-east-1"})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := MC.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("bucket '%s' already have\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	}
	return nil
}

func PutFile(ctx context.Context, logger *zap.SugaredLogger, file *multipart.FileHeader) (pgtype.UUID, error) {
	file_size := file.Size
	object, err := file.Open()
	if err != nil {
		return pgtype.UUID{}, err
	}

	id := uuid.New()
	info, err := MC.PutObject(ctx, constants.BUCKETNAME, id.String(), object, file_size, minio.PutObjectOptions{ContentType: "multipart/form-data"})
	if err != nil {
		logger.Error(err)
		return pgtype.UUID{}, err
	}
	if err := object.Close(); err != nil {
		logger.Error(err)
		return pgtype.UUID{}, err
	}
	logger.Info(info)

	return pgtype.UUID{Bytes: id, Valid: true}, nil
}
func GeneratePresignedURL(ctx context.Context, id string) (string, error) {
	reqParams := make(url.Values)
	presignedURL, err := MC.PresignedGetObject(ctx, constants.BUCKETNAME, id, time.Second*24*60*60, reqParams)
	if err != nil {
		return "", err
	}
	return presignedURL.String(), nil
}

func GetFileURL(c echo.Context, logger *zap.SugaredLogger, id uuid.UUID) string {
	url, err := GeneratePresignedURL(c.Request().Context(), id.String())
	if err != nil {
		logger.Error(err)
		//default image if can find image by uuid
		return "https://imgur.com/UniMfif.png"
	}
	return url
}
