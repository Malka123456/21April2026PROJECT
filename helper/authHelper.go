package helper

import (
	"crypto/rand"
	"errors"
	"fmt"
	"learning-backend/config"
	"learning-backend/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthHelper struct {

	Secret string
}


func (h AuthHelper) GenerateHashedPassword(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", errors.New("Could not hash password")
	}
	return string(hashedPassword), nil
}

func (h AuthHelper) VerifyPassword(pP string, hP string) error {

	if len(pP) < 6 {
		return errors.New("password length should be at least 6 characters long")
	}

	err := bcrypt.CompareHashAndPassword([]byte(hP), []byte(pP))

	if err != nil {
		return errors.New("password does not match")
	}

	return nil
}



func (h AuthHelper) GenerateToken(userID uint) (string, error) {

		if userID == 0 {
			return "", errors.New("invalid user data for token generation") 
		}	

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": userID,

			"exp":     jwt.TimeFunc().Add(24 * time.Hour * 7).Unix(), // Token expires in a week
		})

		tokenString, err := token.SignedString([]byte(h.Secret))

		if err != nil {
			return "", fmt.Errorf("unable to sign token: %w", err)
		}

		return tokenString, nil
}

func (h AuthHelper) GetCurrentUser(ctx *fiber.Ctx) models.User {
	 user := ctx.Locals("user") 
	 return user.(models.User)
}

func (h AuthHelper) GenerateCode() (string, error) {
	return  RandomNumbers(6)
}

func RandomNumbers(length int) (string, error) {

	const numbers = "1234567890"

	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	numLength := len(numbers)

	for i := 0; i < length; i++ {
		buffer[i] = numbers[int(buffer[i])%numLength]
	}

	return string(buffer), nil
}


func NewAuthHelper(cfg config.AppConfig) AuthHelper {
	return AuthHelper{
		Secret: cfg.JWTSecret,
		
	}
}