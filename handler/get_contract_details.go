package handler

import (
	"learn-gorm/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *handler) GetContractDetails(c *gin.Context) {

	var reqData models.ContractRequest
	err := c.ShouldBindJSON(&reqData)
	if err != nil {
		log.Println(err.Error())
		errResp := models.BuildError("error when binding request body", err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	resultData, errCtrl := handler.contoller.GetContractDetails(reqData)
	if errCtrl != nil {
		c.JSON(http.StatusInternalServerError, errCtrl)
		return
	}

	c.JSON(http.StatusOK, resultData)
}
