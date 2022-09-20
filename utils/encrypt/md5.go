package encrypt

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(origin string) string {
	h := md5.New()
	h.Write([]byte(origin))
	resStr := h.Sum(nil)
	return hex.EncodeToString(resStr)
}
