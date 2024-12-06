package main

import (
	// Aca van todas las importaciones de archivos y dependencias
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
    "net/http"
    "time"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

func main() {
    r := gin.Default()

    r.POST("/login", login)
    r.GET("/welcome", authenticateJWT(), welcome)

    r.Run(":8080")
}

func login(c *gin.Context) {
    var creds struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    if err := c.BindJSON(&creds); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    // Aquí deberías validar las credenciales del usuario
    if creds.Username != "user" || creds.Password != "password" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    expirationTime := time.Now().Add(5 * time.Hour)
    claims := &Claims{
        Username: creds.Username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func authenticateJWT() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")

        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
            c.Abort()
            return
        }

        claims := &Claims{}
        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        c.Set("username", claims.Username)
        c.Next()
    }
}

func welcome(c *gin.Context) {
    username, _ := c.Get("username")
    c.JSON(http.StatusOK, gin.H{"message": "Welcome " + username.(string)})
}