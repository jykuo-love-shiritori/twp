package minio

import (
	"context"
	"fmt"
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

	endpoint := "localhost:" + os.Getenv("MINIO_API_PORT")
	accessKey := os.Getenv("MINIO_ACCESS_KEY")
	secretKey := os.Getenv("MINIO_SECRET_KEY")

	// log.Printf("Endpoint: %s", endpoint)
	// log.Printf("AccessKeyID: %s", accessKey)
	// log.Printf("SecretAccessKey: %s", secretKey)

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

func PutFile(ctx context.Context, logger *zap.SugaredLogger, file *multipart.FileHeader) (pgtype.UUID, error) {
	file_size := file.Size
	object, err := file.Open()
	if err != nil {
		return pgtype.UUID{}, err
	}
	defer object.Close()

	id := uuid.New()
	info, err := MC.PutObject(ctx, constants.BUCKETNAME, id.String(), object, file_size, minio.PutObjectOptions{ContentType: "multipart/form-data"})
	if err != nil {
		fmt.Println(err)
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
