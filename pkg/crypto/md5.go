package crypto

import (
	"crypto/md5"
	"fmt"
)

func ToMD5(data []byte) string {
	return fmt.Sprintf("%x", md5.Sum(data))
}
