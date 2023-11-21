package auth

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"net/url"

	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type loginParams struct {
	ClientId            string              `query:"client_id"`
	CodeChallenge       string              `query:"code_challenge"`
	CodeChallengeMethod codeChallengeMethod `query:"code_challenge_method"`
	RedirectUri         string              `query:"redirect_uri"`
	ResponseType        responseType        `query:"response_type"`
	Scope               string              `query:"scope"`
	State               string              `query:"state"`
}

func Login(db *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var params loginParams
		err := c.Bind(&params)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest, "Failed to parse data")
		}

		// check valid code challenge method
		// only "S256" and "plain" are supported
		switch params.CodeChallengeMethod {
		case s256:
		case plain:
		default:
			return echo.NewHTTPError(http.StatusBadRequest, "Unsupported code_challenge_method")
		}

		// check valid response type
		// only "code" is supported
		switch params.ResponseType {
		case code:
		default:
			return echo.NewHTTPError(http.StatusBadRequest, "Unsupported response_type")
		}

		buf := make([]byte, 32)
		_, err = rand.Read(buf)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Unexpected error")
		}
		code := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(buf)

		mu.Lock()
		codeChallengePairs[code] = challenge{params.CodeChallenge, params.CodeChallengeMethod}
		mu.Unlock()

		redirectUri, err := url.Parse(params.RedirectUri)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid redirect URI")
		}
		values := redirectUri.Query()
		values.Add("code", code)
		redirectUri.RawQuery = values.Encode()

		return c.Redirect(http.StatusFound, redirectUri.String())
	}
}
