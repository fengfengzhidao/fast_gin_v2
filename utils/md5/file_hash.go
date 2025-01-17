package md5

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func MD5WithFile(file io.Reader) string {
	m := md5.New()
	io.Copy(m, file)
	sum := m.Sum(nil)
	return hex.EncodeToString(sum)
}
