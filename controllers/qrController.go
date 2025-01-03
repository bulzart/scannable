package controllers

import (
	"artwear/initializers"
	"artwear/models"
	"fmt"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/gin-gonic/gin"
	"image/png"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func CreateQR(ctx *gin.Context) {
	var body struct {
		Title string
		Url   string
	}
	rand.Seed(time.Now().UnixNano())

	// Generate a random integer between 0 and 100
	randomNumber := rand.Intn(999999999999)
	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not read the body",
		})
		return
	}

	if !strings.HasPrefix(body.Url, "https://") {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "URL must start with 'https://'",
		})
		return
	}

	var QR_ models.QR_
	url := strings.Split(body.Url, "://")
	initializers.DB.First(&QR_, "url like ?", "%"+url[1]+"%")
	if QR_.ID == 0 {
		qrCode, _ := qr.Encode(body.Url, qr.M, qr.Auto)
		qrCode, _ = barcode.Scale(qrCode, 365, 365)
		err := os.MkdirAll("qrcodes", os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}
		file, _ := os.Create("qrcodes/qrcode" + strconv.Itoa(randomNumber) + ".png")

		user, exists := ctx.Get("user")
		if !exists {
			// Handle the case where user is not retrieved
			ctx.AbortWithStatusJSON(400, gin.H{"message": "User not retrieved"})
			return
		}

		qr_ := models.QR_{Url: body.Url, UserID: user.(models.User).ID, ImageURL: "qrcode" + strconv.Itoa(randomNumber)}
		if err := initializers.DB.Create(&qr_).Error; err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Error while saving the QR",
			})
			return
		}
		defer file.Close()
		png.Encode(file, qrCode)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "QR created successfully!",
		})
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"message": "The ulr " + url[1] + " is already in use",
	})

}

func GetUserQRs(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.AbortWithStatusJSON(400, gin.H{"message": "User not retrieved"})
		return
	}

	fmt.Println(user.(models.User).ID)
	var _user models.User
	// Use Preload to eagerly load the associated QR codes for the user
	initializers.DB.Preload("QR").First(&_user, user.(models.User).ID)
	c.JSON(http.StatusOK, gin.H{
		"qr": _user.QR,
	})
}
func UpdateQR(c *gin.Context) {
	var body models.QR_
	var QR models.QR_
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not read request body",
		})
		return
	}
	id := c.Param("id")
	initializers.DB.First(&QR, id)
	if initializers.DB.Model(&QR).Updates(body).Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not update the QR",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "QR updated successfully!",
	})
}

func GetQRbyId(c *gin.Context) {
	id := c.Param("id")
	var qrRedirect models.QR_
	initializers.DB.Find(&qrRedirect, id)
	if qrRedirect.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "QR not found!",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"redirects": &qrRedirect,
	})
}
func DeleteQR(c *gin.Context) {
	id := c.Param("id")
	if initializers.DB.Delete(&models.QR_{}, id).Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not delete the QR",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "QR deleted successfully!",
	})
}
