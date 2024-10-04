package utils

import (
    "errors"
    "fmt"
    "os"
    "time"

    "github.com/golang-jwt/jwt/v4"
)

const AUTH_TOKEN_DURATION = time.Hour * 24 * 365

type AuthToken struct {
    jwt.StandardClaims
    User *string `json:"user"`
    Session *string `json:"session"`
}

func (a AuthToken) Expired() bool {
    return a.ExpiresAt < time.Now().Unix()
}

func NewAuthJWT(user *string, session *string) (string) {
    claims := AuthToken{
        User: user,
        Session: session,
    }
    claims.ExpiresAt = time.Now().Add(AUTH_TOKEN_DURATION).Unix()
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signed, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
    return signed
}

func ParseAuthJWT(token string) (*AuthToken, error) {
    claims := new(AuthToken)
    parsed, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(os.Getenv("JWT_SECRET")), nil
    })
    if err != nil {
        return nil, err
    }
    if !parsed.Valid {
        return nil, errors.New("invalid token")
    }
    return claims, nil
}
