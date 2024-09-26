package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"os"
	"strings"

	"github.com/google/uuid"
)

func HashKey(key string) (hashed string) {
    signingKey := []byte(os.Getenv("APIKEY_HASHING_KEY"))
    hash := hmac.New(sha256.New, signingKey)
    hash.Write([]byte(key))
    hashed = hex.EncodeToString(hash.Sum(nil))
    return hashed
}

func GenerateApiKey() (key string, hashed string) {
    // genero un uuid
    key = uuid.New().String()
    key = strings.ReplaceAll(key, "-", "")
    src := []byte(key)
    // lo paso a base64
    e := base64.RawStdEncoding
    dst := make([]byte, e.EncodedLen(len(src)))
    e.Encode(dst, src)
    encoded := string(dst)
    // lo hasheo
    hashed = HashKey(encoded)
    return encoded, hashed
}
