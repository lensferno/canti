package codecs

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
)

func HmacMd5(key, data string) string {
	h := hmac.New(md5.New, []byte(key))
	h.Write([]byte(data))

	return hex.EncodeToString(h.Sum(nil))
}
