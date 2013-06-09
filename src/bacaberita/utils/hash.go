package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
)

func Sha1(str string) string {
	h := sha1.New()
	io.WriteString(h, str)
	return hex.EncodeToString(h.Sum(nil))
}
