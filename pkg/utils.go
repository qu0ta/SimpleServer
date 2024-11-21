package pkg

import (
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
