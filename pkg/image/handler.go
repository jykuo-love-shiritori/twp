package image

import (
	"net/http"

	minioWrapper "github.com/jykuo-love-shiritori/twp/minio"
	"github.com/jykuo-love-shiritori/twp/pkg/common"
	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"
)

func RegisterImage(e *echo.Echo, mc *minioWrapper.MC, logger *zap.SugaredLogger) {
	e.GET("/image/:fileName", func(c echo.Context) error {
		ctx := c.Request().Context()

		var fileName string
		err := echo.PathParamsBinder(c).String("fileName", &fileName).BindError()
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		file, err := mc.GetFile(ctx, fileName)
		if err != nil {
			logger.Error(err)
			errResp := minio.ToErrorResponse(err)
			if errResp.StatusCode == http.StatusNotFound {
				return echo.NewHTTPError(http.StatusNotFound)
			}
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.Blob(http.StatusOK, common.FileMimeFrom(fileName), file)
	})
}
