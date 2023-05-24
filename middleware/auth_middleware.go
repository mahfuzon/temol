package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/mahfuzon/temol/libraries"
	"github.com/mahfuzon/temol/response/api_response"
	"github.com/mahfuzon/temol/service"
	"net/http"
	"strings"
	"time"
)

func AuthMiddleware(userService service.UserService, authService service.AuthService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")

			if !strings.Contains(authHeader, "Bearer") {
				response := api_response.ConverseToErrorResponse("Unauthorized", "invalid token")
				return c.JSON(http.StatusUnauthorized, response)
			}

			tokenString := ""
			arrayToken := strings.Split(authHeader, " ")
			if len(arrayToken) == 2 {
				tokenString = arrayToken[1]
			}

			token, err := authService.ValidateToken(tokenString)
			if err != nil {
				response := api_response.ConverseToErrorResponse("Unauthorized", err.Error())
				return c.JSON(http.StatusUnauthorized, response)
			}

			claim, err := libraries.DecodeEncodedTokenToMapClaim(token)
			if err != nil {
				response := api_response.ConverseToErrorResponse("Unauthorized", err.Error())
				return c.JSON(http.StatusUnauthorized, response)
			}

			if time.Now().Unix() > claim.ExpiredAt {
				response := api_response.ConverseToErrorResponse("Unauthorized", "expired token")
				return c.JSON(http.StatusUnauthorized, response)
			}

			userID := claim.UserId

			userResponse, err := userService.FindById(userID)
			if err != nil {
				response := api_response.ConverseToErrorResponse("Unauthorized", err.Error())
				return c.JSON(http.StatusUnauthorized, response)
			}

			c.Set("user", userResponse)
			return next(c)
		}
	}
}
