package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {

	tokenString := c.Request.URL.Query().Get("jwt")
	jwtSecret := os.Getenv("QN_SSO_SECRET")

	// Parse the token string
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// Define the secret key used to sign the token
		return []byte(jwtSecret), nil
	})

	// Handle any parsing errors
	if err != nil {
		fmt.Println("[1] Error parsing token:", err)
		c.String(http.StatusUnauthorized, "Unauthorized - Error parsing token: %v", err)
		return
	}

	// Verify the token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user := claims["name"].(string)
		organization := claims["organization_name"].(string)
		email := claims["email"].(string)

		fmt.Printf("%v", claims)

		fmt.Printf("Token is valid and verified for user: %s\n", user)

		// FILL ME: add your dashboard logic and view here
		c.String(http.StatusOK, "Welcome to the dashboard, %s from %s with email: %s\n\n", user, organization, email)
		return
	} else {
		c.String(http.StatusUnauthorized, "Unauthorized - Failed JWT verification")
		return
	}
}
