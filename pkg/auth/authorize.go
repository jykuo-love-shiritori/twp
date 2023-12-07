package auth

import (
	"net/http"

	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type loginParams struct {
	ClientId            string              `json:"client_id"`
	CodeChallenge       string              `json:"code_challenge"`
	CodeChallengeMethod codeChallengeMethod `json:"code_challenge_method"`
	RedirectUri         string              `json:"redirect_uri"`
	ResponseType        responseType        `json:"response_type"`
	Scope               string              `json:"scope"`
	State               string              `json:"state"`
	Email               string              `json:"email"`
	Password            string              `json:"password"`
}

func Authorize(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
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
			return echo.NewHTTPError(http.StatusBadRequest, "Unsupported code challenge method")
		}

		// check valid response type
		// only "code" is supported
		switch params.ResponseType {
		case code:
		default:
			return echo.NewHTTPError(http.StatusBadRequest, "Unsupported response type")
		}

		result, err := pg.Queries.FindUserInfoAndPassword(c.Request().Context(), params.Email)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(params.Password))
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Wrong username, email, or password")
		}

		// generate OTP
		code, err := generateRandomString(32)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		mu.Lock()
		codeChallengePairs[code] = challengeUser{
			CodeChallenge:   params.CodeChallenge,
			ChallengeMethod: params.CodeChallengeMethod,
			Username:        result.Username,
			Role:            result.Role,
		}
		mu.Unlock()

		return c.JSON(http.StatusOK, echo.Map{
			"code":  code,
			"state": params.State,
		})
	}
}
