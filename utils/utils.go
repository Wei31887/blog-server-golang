package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"math/rand"
)

// Encode string by MD5 algorithm
func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// turn json into string type
func ParseJsonString(value interface{}) (string) {
	newValue, _ := json.Marshal(value)
	key := string(newValue)
	return key
}

var alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString(n int) string {
	b := make([]byte, n)
    for i := range b {
        b[i] = alpha[rand.Intn(len(alpha))]
    }
    return string(b)
}