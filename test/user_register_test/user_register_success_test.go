package user_register_test

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
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRegisterSuccess(t *testing.T) {
	db := libraries.SetDbTest()
	test.TruncateTableUsers(db)

	// buat dummy user
	user := models.User{
		Name:        "mahfuzon",
		Email:       helper.ConverseToSqlNullString("mahfuzon0@gmail.com"),
		Password:    "12345678",
		PhoneNumber: helper.ConverseToSqlNullString("08965456765"),
		Address:     "aasasa",
	}
	err := db.Create(&user).Error
	if err != nil {
		assert.NoError(t, err)
	}

	user2 := models.User{
		Name:        "mahfuzon",
		Email:       helper.ConverseToSqlNullString("mahfuzon0@gmail.co"),
		Password:    "12345678",
		PhoneNumber: helper.ConverseToSqlNullString("08965456764"),
		Address:     "aasasa",
	}
	err = db.Create(&user2).Error
	if err != nil {
		assert.NoError(t, err)
	}
	// end buat dummy user

	requestJsonString := `{
  "name" : "mahfuzon akhiar",
"email" : "mahfuzon0@gmail.coma",
  "password" : "12345678",
  "phone_number" : "089654567688",
  "address" : "indonesi"
}`

	userController := test.SetupUserController(db)

	router := libraries.SetRouter()
	router.POST("api/v1/auth/register", userController.Register)

	req := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/auth/register", strings.NewReader(requestJsonString))
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
