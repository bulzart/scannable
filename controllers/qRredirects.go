package controllers

import (
	"artwear/initializers"
	"artwear/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func CreateRedirect(c *gin.Context) {
	var body struct {
		Url       string    `binding:"required" json:"url"`
		StartDate time.Time `binding:"required" json:"start_date"`
		EndDate   time.Time `binding:"required" json:"end_date"`
	}

	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not read body or fields not correct",
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
	fmt.Println(qr_id)
	qr_redirect := models.QR_redirect{Url: body.Url, StartDate: body.StartDate, EndDate: body.EndDate, QrID: uint(qr_id)}
	if initializers.DB.Create(&qr_redirect).Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not add QR scheduled redirect",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "QR scheduled redirect created successfully!",
	})
}

func GetLatestRedirect(c *gin.Context) {
	id := c.Param("id")
	var qrRedirect models.QR_redirect
	initializers.DB.Where("start_date < ?", time.Now()).Where("end_date > ?", time.Now()).Find(&qrRedirect, "qr_id = ?", id)
	if qrRedirect.ID == 0 {
		var qr models.QR_
		initializers.DB.Find(&qr, "id = ?", id)
		c.JSON(http.StatusOK, gin.H{
			"redirects": &qr.Url,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"redirects": &qrRedirect.Url,
	})
}
