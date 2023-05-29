package user_login_test

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mahfuzon/temol/helper"
	"github.com/mahfuzon/temol/libraries"
	"github.com/mahfuzon/temol/models"
	"github.com/mahfuzon/temol/response/api_response"
	"github.com/mahfuzon/temol/test"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUserLoginSuccess(t *testing.T) {
	db := libraries.SetDbTest()
	test.TruncateTableUsers(db)

	passwordHash, err := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.MinCost)
	assert.NoError(t, err)

	//create dummy user
	user := models.User{
		Name:        "mahfuzon",
		Email:       helper.ConverseToSqlNullString("mahfuzon0@gmail.com"),
		Password:    string(passwordHash),
		PhoneNumber: helper.ConverseToSqlNullString("081278160990"),
		Address:     "aasasa",
	}

	err = db.Create(&user).Error
	assert.NoError(t, err)

	userController := test.SetupUserController(db)

	requestJsonString := `{
  "password" : "12345678",
  "phone_number" : "081278160990"
}`

	router := libraries.SetRouter()
	router.POST("api/v1/auth/login", userController.Login)

	req := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/auth/login", strings.NewReader(requestJsonString))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	result := rec.Result()
	assert.Equal(t, 200, result.StatusCode)

	body := result.Body

	responseBody, _ := io.ReadAll(body)
	var apiResponse api_response.ApiResponse

	err = json.Unmarshal(responseBody, &apiResponse)
	assert.NoError(t, err)

	fmt.Println(apiResponse.Data)
}
