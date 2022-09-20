package encrypt

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

func Sha1(origin string) string {
	h := sha1.New()
	h.Write([]byte(origin))
	resStr := h.Sum(nil)
	return hex.EncodeToString(resStr)
}

func Sha256(origin string) string {
	h := sha256.New()
	h.Write([]byte(origin))
	resStr := h.Sum(nil)
	return hex.EncodeToString(resStr)
}
