package handler

import (
	"learn-gorm/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *handler) GetUserByName(c *gin.Context) {

	var customerData models.CustomerRequest
	err := c.ShouldBindJSON(&customerData)
	if err != nil {
		log.Println(err.Error())
		errResp := models.BuildError("error when binding request body", err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	resultData, errCtrl := handler.contoller.GetUserByName(customerData)
	if errCtrl != nil {
		c.JSON(http.StatusInternalServerError, errCtrl)
		return
	}

	c.JSON(http.StatusOK, resultData)
}
