package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"todo-hexagonal-arch/internal/core/interfaces"
)

type authMiddleware struct {
	tokenUtil interfaces.TokenUtil
}

func NewAuthMiddleware(tokenUtil interfaces.TokenUtil) *authMiddleware {
	return &authMiddleware{tokenUtil: tokenUtil}
}

func (am *authMiddleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			token := t[1]
			claims, err := am.tokenUtil.ValidateToken(token)
			if err != nil {
				c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": err.Error()})
				c.Abort()
				return
			}
			c.Set("x-user-id", claims.UserID)
			c.Next()
			return
		}
		c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Not authorized"})
		c.Abort()
		return
	}
}
