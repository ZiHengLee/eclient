package crypto

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestAes(t *testing.T) {
	key := []byte("0123456789abcdef")
	c := NewAESCoder(key)
	raw := []byte("abc")
	dat, err := c.AES128CBCEncrypt(raw)
	if err != nil {
		fmt.Printf("aes encrypt err:%v\n", err)
		return
	}
	fmt.Printf("dat:%v\n", hex.EncodeToString(dat))
}
