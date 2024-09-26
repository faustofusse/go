package utils

import (
    "errors"
    "fmt"
    "os"
    "time"

    "github.com/golang-jwt/jwt/v4"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

const (
    EMAIL_EXPIRY = time.Hour * 24
    PASSWORD_EXPIRY = time.Hour * 24
    SESSION_EXPIRY = time.Hour * 24 * 365
)

func createJWT(claims jwt.Claims) string {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
    if err != nil {
        panic(err)
    }
    return tokenString
}

func parseJWT[T jwt.Claims, PT interface{ *T; jwt.Claims }](tokenString string) (*T, error) {
    claims := PT(new(T))
    if token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(os.Getenv("JWT_SECRET")), nil
    }); err != nil {
        return nil, err
    } else if !token.Valid {
        return nil, errors.New("Invalid token")
    } else {
        result := *claims
        return &result, nil
    }
}

// ------------- TYPES

type EmailVerification struct {
    ID          primitive.ObjectID      `json:"id"`
    Name        string                  `json:"name"`
    Expiry      int64                   `json:"expiry"`
    jwt.StandardClaims
}

type AuthToken struct {
    User        primitive.ObjectID     `json:"user"`
    Session     primitive.ObjectID     `json:"session"`
    Expiry      int64                   `json:"expiry"`
    jwt.StandardClaims
}

type PasswordToken struct {
    User        primitive.ObjectID      `json:"user"`
    Expiry      int64                   `json:"expiry"`
    jwt.StandardClaims
}

// ------------- PARSE

func ParseAuthJWT(tokenString string) (*AuthToken, error) {
    return parseJWT[AuthToken](tokenString)
}

func ParsePasswordJWT(tokenString string) (*PasswordToken, error) {
    return parseJWT[PasswordToken](tokenString)
}

func ParseEmailJWT(tokenString string) (*EmailVerification, error) {
    return parseJWT[EmailVerification](tokenString)
}

// ------------- CREATE

func CreateAuthJWT(userId primitive.ObjectID, sessionId primitive.ObjectID) string {
    return createJWT(AuthToken{
        User: userId,
        Session: sessionId,
        Expiry: time.Now().Add(SESSION_EXPIRY).Unix(),
    })
}

func CreateEmailJWT(userId *primitive.ObjectID, name string) string {
    return createJWT(EmailVerification{
        ID: *userId,
        Name: name,
        Expiry: time.Now().Add(EMAIL_EXPIRY).Unix(),
    })
}

func CreatePasswordJWT(userId primitive.ObjectID) string {
    return createJWT(PasswordToken{
        User: userId,
        Expiry: time.Now().Add(PASSWORD_EXPIRY).Unix(),
    })
}

// ------------- UTILS

func AuthExpired(tokenString string) bool {
    result, err := ParseAuthJWT(tokenString)
    if err != nil {
        panic("Could not parse token, this should never happen")
    }
    return result.Expired()
}

func (a PasswordToken) Expired() bool {
    return a.Expiry < time.Now().Unix()
}

func (a AuthToken) Expired() bool {
    return a.Expiry < time.Now().Unix()
}

