package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CorsConfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTION")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding, Authorization, X-API-Key")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
		//fmt.Println(c.Request.Method)
		//fmt.Println(c.Request.Response)
		//fmt.Println(c.Request.WithContext(c))
	}
}

func CheckAPIKey(clientKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if clientKey != c.GetHeader("X-API-Key") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid api key")
		}

		c.Next()
	}
}

func ValidateAndSetToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		type TokenClaims struct {
			Email string `json:"email"`
			jwt.StandardClaims
		}

		accessToken := c.GetHeader("Authorization")
		//accessToken = accessToken[len("bearer "):]

		if accessToken == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, "validateToken error: missing token")
		}
		fmt.Println(accessToken)

		/*
			claims := &TokenClaims{}
			token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
				// Make sure that the token method conforms to "SigningMethodHMAC"
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("ValidateToken unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(internal.AppConfig.JwtKey()), nil
			})
			if err != nil {
				if err == jwt.ErrSignatureInvalid {
					c.AbortWithStatusJSON(http.StatusBadRequest, "validateToken error: invalid signature: "+err.Error())
				}

				c.AbortWithStatusJSON(http.StatusInternalServerError, "validateToken error: "+err.Error())
			}
			if !token.Valid {
				c.AbortWithStatusJSON(http.StatusBadRequest, "validateToken error: invalid token: "+err.Error())
			}

			//do claims validation here


			c.Set("token", token)
			c.Set("claims", claims)
		*/
		c.Next()
	}
}

func ErrorLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Continue processing the request

		// If errors are present after processing
		if len(c.Errors) > 0 {
			var errors []map[string]interface{}

			for _, err := range c.Errors {
				// Log each error
				log.Printf("Time: %s, URL: %s, Error: %s, HTTP Code: %d\n",
					time.Now().Format(time.RFC3339),
					c.Request.URL.String(),
					err.Error(),
					c.Writer.Status(),
				)
				// Add the error to the slice for JSON response
				errors = append(errors, map[string]interface{}{
					"time":       time.Now().Format(time.RFC3339),
					"url":        c.Request.URL.String(),
					"error":      err.Error(),
					"httpStatus": c.Writer.Status(),
				})
			}

			// Respond with all logged errors in JSON format
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"errors": errors,
			})
		}
	}
}
