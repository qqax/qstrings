package qrandom

import (
	"crypto/rand"
	"errors"
	"net/url"
)

func RandByteArray(n int) ([]byte, error) {
	randArray := make([]byte, n)
	if _, err := rand.Read(randArray); err != nil {
		return nil, errors.New("unable to create random byte array")
	}
	return randArray, nil
}

func RandUrl(n int) (string, error) {
	randArray := make([]byte, n)
	if _, err := rand.Read(randArray); err != nil {
		return "", errors.New("unable to create random url")
	}
	return url.QueryEscape(string(randArray)), nil
}

//func RandomString(n int) (string, error) {
//	randArray := make([]byte, n)
//	if _, err := rand.Read(randArray); err != nil {
//		return "", errors.New("unable to create random string")
//	}
//	return string(randArray), nil
//}
