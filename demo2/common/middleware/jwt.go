package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/coldchain-traceability-system/common/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// JWTAuthMiddleware 提供 JWT 认证中间件
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 Authorization 头部
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "必须提供 Authorization 头部",
			})
			c.Abort()
			return
		}

		// 检查头部格式是否正确
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization 头部格式必须为 Bearer {token}",
			})
			c.Abort()
			return
		}

		// 解析并验证令牌
		tokenString := parts[1]
		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			// 验证签名方法
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("意外的签名方法: %v", token.Header["alg"])
			}
			return []byte(config.AppConfig.JWT.Secret), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "无效的令牌",
			})
			c.Abort()
			return
		}

		// 将声明设置到上下文中
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		c.Next()
	}
}

// Claims 表示 JWT 声明结构
type Claims struct {
	UserID   uint64 `json:"user_id"`  // 用户 ID
	Username string `json:"username"` // 用户名

	jwt.RegisteredClaims
}

// GenerateToken 生成新的 JWT 令牌
func GenerateToken(userID uint64, username, role string) (string, error) {
	expirationTime := time.Now().Add(time.Duration(config.AppConfig.JWT.ExpireTime) * time.Hour)

	claims := &Claims{
		UserID:   userID,
		Username: username,

		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "coldchain-traceability-system",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.AppConfig.JWT.Secret))
}
