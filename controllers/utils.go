package controllers

import (
	"errors"
	"log"
	"myGram/config"
	"regexp"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	if !IsPasswordValid(password) {
		return "", errors.New(`password minimum 6 characters`)
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func IsEmailValid(address string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return re.MatchString(address)
}

func IsPasswordValid(password string) bool {
	return len([]rune(password)) >= 6
}

func GetSecretKey() (string, error) {
	conf, err := config.LoadConfig("../")
	if err != nil {
		return "", err
	}
	return conf.SecretKey, nil
}

func GetContentType(c *gin.Context) string {
	return c.Request.Header.Get("Content-Type")
}

func GenerateToken(id uint, email string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	conf, err := GetSecretKey()
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}
	signedToken, _ := parseToken.SignedString([]byte(conf))
	return signedToken
}

func VerifyToken(c *gin.Context) (interface{}, error) {
	errResponse := errors.New("sign in proceed")
	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errResponse
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		conf, err := GetSecretKey()
		if err != nil {
			log.Fatal("Could not load environment variables", err)
		}

		return []byte(conf), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errResponse
	}

	return token.Claims.(jwt.MapClaims), nil
}

func GetUserId(c* gin.Context) uint {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	return userId
}