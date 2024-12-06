package middlewares

import (
    "net/http"
    "strings"
    "time"

    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

type AuthMiddleware struct{}

func NewAuthMiddleware() *AuthMiddleware {
    return &AuthMiddleware{}
}

func (auth *AuthMiddleware) ValidateToken(c *gin.Context) {
    // Obtener el header "Authorization"
    authToken := c.GetHeader("Authorization")

    if authToken == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
        c.Abort()
        return
    }

    // Eliminar el prefijo "Bearer " si está presente
    tokenString := strings.TrimPrefix(authToken, "Bearer ")

    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if err != nil || !token.Valid {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
        c.Abort()
        return
    }

    // Verificar si el token ha expirado
    if claims.ExpiresAt < time.Now().Unix() {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
        c.Abort()
        return
    }

    // Establecer el usuario en el contexto
    c.Set("username", claims.Username)

    c.Next()
}