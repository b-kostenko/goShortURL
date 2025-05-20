package middleware

import (
	"github.com/gin-gonic/gin"
)

func RequireAuthMiddleware(c *gin.Context) {
	//var user User
	//authHeader := c.GetHeader("Authorization")
	//
	//if authHeader == "" {
	//	return
	//}
	//tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	//
	//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	//	return []byte(os.Getenv("JWT_SECRET")), nil
	//}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	//
	//if err != nil {
	//	return
	//}
	//
	//if claims, ok := token.Claims.(jwt.MapClaims); ok {
	//	postgres.DB.First(&user, claims["sub"])
	//	if user.ID == 0 {
	//		return
	//	}
	//}
	//
	//c.Set("user", user)
	c.Next()
}
