package pkg

import (
	"crypto/sha1"
	"fmt"
)

func Encrypt(plaintext string) (cryptext string) {
	// hash値の作成
	cryptext = fmt.Sprint("%x", sha1.Sum([]byte(plaintext)))
	return
}
