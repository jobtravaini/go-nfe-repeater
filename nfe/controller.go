package nfe

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type NfeController struct{
	Service INfeService
}

func NewNfeController(service INfeService) NfeController {
	return NfeController{Service:service}
}

func (controller NfeController) RetrieveNfe(c *gin.Context) {
	accessKey, ok := c.Request.URL.Query()["key"]

	if !ok || "" == accessKey[0] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing query parameter key"})
		return
	}

	value, err := controller.Service.GetNfe(accessKey[0])

	if err != nil {
		handleServiceError(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_key": value.AccessKey, "xml": value.XmlValue})
}

func handleServiceError(err error, c *gin.Context) {
	if gorm.IsRecordNotFoundError(err) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The nfe requested was not found"})
		return
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred while retrieve your NFE, please try again later"})
		return
	}
}