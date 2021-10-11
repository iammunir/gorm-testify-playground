package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"learn-gorm/controller"
	"learn-gorm/models"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var controllerMock = controller.ControllerMock{Mock: mock.Mock{}}
var handlerTest = NewHandler(&controllerMock)

func MockJsonPost(c *gin.Context, content interface{}) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")

	jsonbytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
}

func TestGetCustomerById(t *testing.T) {

	// Setup
	gin.SetMode(gin.TestMode)

	// set data mocking
	reqData := models.CustomerRequest{
		CustomerId: "10",
		RequestDetails: models.RequestDetails{
			Address: 1,
			Company: 1,
			Contact: 1,
		},
	}

	resp := models.Customer{
		FirstName: "John",
		LastName:  "Doe",
	}

	// set recorder and gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}

	MockJsonPost(c, reqData)

	controllerMock.Mock.On("GetCustomerById", reqData).Return(resp)
	handlerTest.GetCustomerById(c)

	response := w.Result()
	body, _ := io.ReadAll(response.Body)

	var bodyResp models.Customer
	json.Unmarshal(body, &bodyResp)

	assert.EqualValues(t, http.StatusOK, response.StatusCode, "status code must be 200")
	assert.EqualValues(t, resp.FirstName, bodyResp.FirstName, "first name must be equal")
}
