package user_login_test

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mahfuzon/temol/libraries"
	"github.com/mahfuzon/temol/response/api_response"
	"github.com/mahfuzon/temol/test"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestChecValidation(t *testing.T) {
	db := libraries.SetDbTest()

	test.TruncateTableUsers(db)
	userController := test.SetupUserController(db)

	requestJsonString := `{
}`

	router := libraries.SetRouter()
	router.POST("api/v1/auth/login", userController.Login)

	req := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/auth/login", strings.NewReader(requestJsonString))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	result := rec.Result()
	assert.Equal(t, 422, result.StatusCode)

	body := result.Body

	responseBody, _ := io.ReadAll(body)
	var apiResponse api_response.ApiResponse

	err := json.Unmarshal(responseBody, &apiResponse)
	assert.NoError(t, err)

	fmt.Println(apiResponse.Data)
}
