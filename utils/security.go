package utils

import (
	"math/rand"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func HashString(password string) (string, error) {
    bytes := []byte(password)
    hashedPassword, err := bcrypt.GenerateFromPassword(bytes, bcrypt.DefaultCost)
    if err != nil { return "", err }
    return string(hashedPassword), nil
}

func CompareHash(normal, hashed string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(normal))
    return err == nil
}

func randomInt(min, max int) int {
    return rand.Intn(max - min) + min
}

func GenerateCode() (string, string) {
    code := strconv.Itoa(randomInt(100000, 999999))
    hashed, _ := HashString(code)
    return code, hashed
}
