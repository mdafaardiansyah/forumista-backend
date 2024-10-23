package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mdafaardiansyah/forumista-backend/internal/configs"
	"github.com/mdafaardiansyah/forumista-backend/pkg/jwt"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	secretKey := configs.Get().Service.SecretJWT
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")

		header = strings.TrimSpace(header)
		if header == "" {
			c.AbortWithError(http.StatusUnauthorized, errors.New("missing token"))
			return
		}

		userID, username, err := jwt.ValidateToken(header, secretKey)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		c.Set("userID", userID)
		c.Set("username", username)
		c.Next()
	}
}
