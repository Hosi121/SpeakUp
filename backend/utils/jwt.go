package utils

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"time"

	"github.com/Hosi121/SpeakUp/config"

	"github.com/golang-jwt/jwt/v5"
)

// JWT トークンを生成する関数
func GenerateJWT(userID string) (string, error) {
	config.LoadEnv()
	// Get the private key
	_, secret := config.GetKey()
	block, _ := pem.Decode(secret)
	if block == nil {
		return "", fmt.Errorf("failed to parse PEM block containing the private key")
	}

	secretKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse private key: %v", err)
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"user_id":    userID,
		"created_at": time.Now().Unix(),
		"exp":        time.Now().Add(time.Hour * 24).Unix(), // Set expiration
	})

	// Sign the token
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %v", err)
	}

	return tokenString, nil
}

// JWT トークンを検証する関数
func ValidateJWT(tokenString string) (string, error) {
	config.LoadEnv()
	// Get the public key
	public, _ := config.GetKey()
	block, _ := pem.Decode(public)
	if block == nil {
		return "", fmt.Errorf("failed to parse PEM block containing the public key")
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse public key: %v", err)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is correct
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return publicKey, nil
	})
	if err != nil {
		return "", fmt.Errorf("failed to parse token: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Get user_id from claims
		if userID, ok := claims["user_id"].(string); ok {
			return userID, nil
		}
		return "", fmt.Errorf("user_id not found in token claims")
	}

	return "", fmt.Errorf("invalid token")
}
