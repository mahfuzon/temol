package user_register_test

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

func TestValidationRequired(t *testing.T) {
	db := libraries.SetDbTest()
	test.TruncateTableUsers(db)

	requestJsonString := `{
}`

	userController := test.SetupUserController(db)

	router := libraries.SetRouter()
	router.POST("api/v1/auth/register", userController.Register)

	req := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/auth/register", strings.NewReader(requestJsonString))
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

	fmt.Println(apiResponse)
}

func TestValidationEmail(t *testing.T) {
	db := libraries.SetDbTest()
	test.TruncateTableUsers(db)

	requestJsonString := `{
"email" : "email"
}`

	userController := test.SetupUserController(db)

	router := libraries.SetRouter()
	router.POST("api/v1/auth/register", userController.Register)

	req := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/auth/register", strings.NewReader(requestJsonString))
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

	fmt.Println(apiResponse)
}

func TestValidationMin(t *testing.T) {
	db := libraries.SetDbTest()
	test.TruncateTableUsers(db)

	requestJsonString := `{
"email" : "email",
"password" :"1234567"
}`

	userController := test.SetupUserController(db)

	router := libraries.SetRouter()
	router.POST("api/v1/auth/register", userController.Register)

	req := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/auth/register", strings.NewReader(requestJsonString))
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

	fmt.Println(apiResponse)
}
