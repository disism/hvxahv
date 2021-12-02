package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/pkg/security"
	"log"
	"strings"
)

// Auth authentication middleware for gin network framework,
// Check whether the Token is carried in the request, and verify whether the Token is correct,
// Will obtain the username by parsing the Token and add the username in the context and set the key to loginUser.
func Auth(c *gin.Context) {
	ht := c.Request.Header.Get("Authorization")
	t := strings.Split(ht, "Bearer ")[1]
	if ht == "" {
		c.JSON(500, gin.H{
			"state":   "500",
			"message": "TOKEN IS NOT CARRIED IN THE REQUEST.",
		})
		c.Abort()
		return
	}

	_, err := security.VerifyToken(t)
	if err != nil {
		c.JSON(500, gin.H{
			"state":   "500",
			"message": "LOGIN FAILED. TOKEN IS INCORRECT!",
		})
		c.Abort()
	} else {
		u, err1 := security.ParseToken(t)
		if err1 != nil {
			log.Println("failed to obtain user through token.")
		}
		c.Set("devices", u.DevicesID)
		c.Set("username", u.User)
		c.Next()
	}

}
