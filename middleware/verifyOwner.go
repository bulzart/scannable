package middleware

import (
	"artwear/initializers"
	"artwear/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func VerifyOwner(c *gin.Context) {
	id := c.Param("id")
	user, exists := c.Get("user")
	var qr models.QR_
	initializers.DB.First(&qr, id)

	if !exists || qr.UserID != user.(models.User).ID {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "You can't edit or delete this QR"})
		return
	}

	c.Next()
}
