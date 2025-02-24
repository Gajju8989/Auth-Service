package jwtkey

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func GetJWTKey() ([]byte, error) {
	err := godotenv.Load("local.env")
	fmt.Printf("yes")
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return nil, fmt.Errorf("JWT_SECRET_KEY not set in environment")
	}

	return []byte(secretKey), nil
}
