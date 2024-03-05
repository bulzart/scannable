package controllers

import (
	"artwear/initializers"
	"artwear/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func createRedirect(c *gin.Context) {
	var body models.QR_redirect
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not read body",
		})
		return
	}
	qr_id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Qr code reference is not valid",
		})
		return
	}
	qr_redirect := models.QR_redirect{Url: body.Url, StartDate: body.StartDate, EndDate: body.EndDate, QRID: uint(qr_id)}
	if initializers.DB.Create(&qr_redirect).Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not add QR scheduled redirect",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "QR scheduled redirect deleted successfully!",
	})
}
