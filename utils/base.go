package utils

import (
	"crypto/hmac"
	"crypto/rand"
	"encoding/hex"
	"hash"
	"strconv"
	"strings"
)

const ()

func SecureRandomHex(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}

	return hex.EncodeToString(bytes)

}

func HmacHexDigest(hashFunc func() hash.Hash, secret, message string) (string, error) {
	h := hmac.New(hashFunc, []byte(secret))

	if _, err := h.Write([]byte(message)); err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

func Map[T any, U any](list []T, mapFunc func(T) U) (mapped []U) {
	mapped = make([]U, len(list))
	for i, v := range list {
		mapped[i] = mapFunc(v)
	}

	return

}

func StringToUint(i string) (parseInt uint) {
	parsedInt, err := strconv.ParseUint(i, 10, 32)
	if err != nil {
		panic(err)
	}

	return uint(parsedInt)
}

func Blank(value string) bool {
	return strings.TrimSpace(value) == ""
}
