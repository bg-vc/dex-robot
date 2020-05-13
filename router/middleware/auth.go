package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vincecfl/dex-robot/handler"
	"github.com/vincecfl/dex-robot/pkg"
	"github.com/vincecfl/dex-robot/pkg/errno"
)

func AuthMiddleware(accountType int) gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := pkg.ParseRequest(c, accountType); err != nil {
			fmt.Printf("AuthMiddleware error: %s\n", err)
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
