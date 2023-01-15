package utils

import (
	"encoding/base64"
)

func Encode(str string) string {
	result := base64.StdEncoding.EncodeToString([]byte(str))
	return result
}

func Decode(s string) (string, error) {
	result, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(result), nil
}
