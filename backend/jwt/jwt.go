package jwt_auth

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
	// 秘密鍵の取得
	_, secret := config.GetKey()
	block, _ := pem.Decode(secret)
	if block == nil {
		fmt.Println("failed to parse PEM block containing the private key")
		return "", nil
	}

	secretKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("failed to parse private key: %v", err)
		return "", nil
	}

	// トークン生成
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"user_id":    userID,
		"created_at": time.Now(),
	})

	// tokenを秘密鍵で暗号化
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// JWT トークンを検証する関数
func ValidateJWT(tokenString string) (*jwt.Token, error) {
	config.LoadEnv()
	// 公開鍵の取得
	public, _ := config.GetKey()
	block, _ := pem.Decode(public)
	if block == nil {
		fmt.Println("failed to parse PEM block containing the public key")
		return nil, nil
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println("failed to parse public key: %v", err)
		return nil, nil
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// RS256で署名されているかを検査
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// 公開鍵で検証するように
		return publicKey, nil
	})

	if err != nil {
		return nil, err
	}

	if token.Valid {
		return token, nil
	}

	return nil, fmt.Errorf("invalid token")
}
