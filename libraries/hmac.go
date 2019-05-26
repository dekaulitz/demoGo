package libraries

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func generateSha256Hex(key string) (expectedSha256String string) {
	hashKey := sha256.New()
	hashKey.Write([]byte(key))
	expectedKey := hashKey.Sum(nil)
	return hex.EncodeToString(expectedKey)
}

//GenerateHmac generates the hmac
func GenerateHmac(message string, key string) (hmacString string) {
	hashKey := generateSha256Hex(key)
	mac := hmac.New(sha256.New, []byte(hashKey))
	mac.Write([]byte(message))
	expectedMAC := mac.Sum(nil)
	return hex.EncodeToString(expectedMAC)
}

//GenerateHmac generates the hmac
func GenerateSimpleHmac(message string, key string) (hmacString string) {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(message))
	expectedMAC := mac.Sum(nil)
	return hex.EncodeToString(expectedMAC)
}

// CheckMAC verifies hash checksum
func CheckMAC(message, messageMAC, key string) bool {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(message))
	hmacExpected := hex.EncodeToString(mac.Sum(nil))
	return hmacExpected == messageMAC
}

//Sign a message using sha256
func Sign(message, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(message))
	signed := mac.Sum(nil)
	return hex.EncodeToString(signed)
}

//SignShaOne sign a message using sha1
func SignShaOne(message, key string) string {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(message))
	signed := mac.Sum(nil)
	return hex.EncodeToString(signed)
}

func ValidateHmac(hmacMessage string, payload string, key string) bool {
	fmt.Println(GenerateHmac(payload, key))
	return hmacMessage == GenerateHmac(payload, key)
}
