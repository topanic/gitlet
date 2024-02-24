package utils

import (
	"crypto/sha1"
	"fmt"
	"time"
)



func GetArgsNum(args []string) int {
	return len(args)
}

/* generate a SHA-1 code */
func GenerateID() string {
	timestamp := time.Now().String()

	hasher := sha1.New()
	hasher.Write([]byte(timestamp))
	hash := hasher.Sum(nil)
	return fmt.Sprintf("%x", hash)
}


