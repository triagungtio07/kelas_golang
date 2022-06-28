package utils

import (
	"errors"
	"fmt"

	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/iancoleman/strcase"
	"github.com/triagungtio07/kelas_golang/config/env"
	"golang.org/x/crypto/bcrypt"
)

func CreateJwtToken(claims jwt.MapClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()
	token.Claims = claims
	sToken, _ := token.SignedString([]byte(env.SecretKey))

	return sToken
}

func Validate(v interface{}) ([]string, error) {
	validate := validator.New()
	err := validate.Struct(v)
	if err == nil {
		return nil, err
	}

	v, ok := err.(validator.ValidationErrors)
	if !ok {
		return nil, errors.New("invalid type")
	}

	messages := []string{}
	for _, err := range err.(validator.ValidationErrors) {
		messages = append(messages, fmt.Sprintf("%s is %s", strcase.ToSnake(err.Field()), err.Tag()))
	}

	return messages, err
}

func EncodePassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hash)
}

func ValidatePassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
