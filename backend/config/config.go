package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func GetDSN() string {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", dbUser, dbPassword, dbHost, dbPort, dbName)
}

func GetSupabaseAPI() (string, string) {
	supabaseRef := os.Getenv("SUPABASE_API_REF")
	supabaseKey := os.Getenv("SUPABASE_API_KEY")

	return supabaseRef, supabaseKey
}

func GetKey() ([]byte, []byte) {
	publicKey := []byte(os.Getenv("PUBLIC_KEY"))
	secretKey := []byte(os.Getenv("SECRET_KEY"))
	return publicKey, secretKey
}
