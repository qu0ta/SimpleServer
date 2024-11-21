package pkg

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
)

func IfExists(filePath string) (bool, string) {
	_, err := os.Stat(filePath)
	if err == nil {
		return true, ""
	} else if os.IsNotExist(err) {
		return false, ""
	} else {
		return false, err.Error()
	}
}
func GenerateSecret() string {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		fmt.Println("Error generating secret:", err)
		return ""
	}
	return base64.StdEncoding.EncodeToString(bytes)
}
