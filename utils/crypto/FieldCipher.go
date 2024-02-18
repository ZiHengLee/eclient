package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"math"
	"math/big"
	"strings"
	"time"
)

type FieldCipher struct {
	keys map[string]*rsa.PrivateKey
}

func NewFieldCipher() (c *FieldCipher) {
	c = &FieldCipher{
		keys: make(map[string]*rsa.PrivateKey),
	}
	return
}

func (c *FieldCipher) AddKey(name string, key []byte) (err error) {
	if _, ok := c.keys[name]; ok {
		err = fmt.Errorf("private key name:%v already exists", name)
		return
	}
	block, _ := pem.Decode(key)
	if block == nil {
		err = fmt.Errorf("invalid rsa private key name:%v", name)
		return
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		err = fmt.Errorf("invalid rsa private key name:%v", name)
		return
	}
	c.keys[name] = priv
	return
}

func (c FieldCipher) GenKey(keynames ...string) (aesKeySrc string, aesKey []byte, ckey string, err error) {
	var k string
	rv, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt32))
	if err != nil {
		return
	}
	v := rv.Int64()
	if len(keynames) > 0 {
		k = keynames[int(v)%len(keynames)]
	} else if len(c.keys) > 0 {
		idx := int(v) % len(c.keys)
		for n := range c.keys {
			if idx == 0 {
				k = n
				break
			}
			idx -= 1
		}
	}
	p := c.keys[k]
	if p == nil {
		err = fmt.Errorf("cant't found private key:%v", k)
		return
	}

	now := time.Now()
	aesKeySrc = fmt.Sprintf("%v-%v", now.Unix(), v)
	rkey, err := rsa.EncryptPKCS1v15(rand.Reader, &p.PublicKey, []byte(aesKeySrc))
	if err != nil {
		return
	}
	ckey = k + "-" + base64.StdEncoding.EncodeToString(rkey)
	aesKeySum := sha256.Sum256([]byte(aesKeySrc))
	aesKey = aesKeySum[:16]
	return
}

func (c FieldCipher) ParseKey(ckey string) (aesKeySrc string, aesKey []byte, err error) {
	idx := strings.Index(ckey, "-")
	if idx <= 0 {
		err = fmt.Errorf("invalid key:%v", ckey)
		return
	}
	name := ckey[:idx]
	p := c.keys[name]
	if p == nil {
		err = fmt.Errorf("cant't found private key:%v", name)
		return
	}
	rkey, err := base64.StdEncoding.DecodeString(ckey[idx+1:])
	if err != nil {
		return
	}
	aesKeySrcBytes, err := rsa.DecryptPKCS1v15(rand.Reader, p, rkey)
	if err != nil {
		return
	}
	aesKeySrc = string(aesKeySrcBytes)
	aesKeySum := sha256.Sum256([]byte(aesKeySrc))
	aesKey = aesKeySum[:16]
	return
}

func (c FieldCipher) Encrypt(dat []byte, aesKey []byte) (res string, err error) {
	coder := NewAESCoder(aesKey)
	rres, err := coder.AES128CBCEncrypt([]byte(dat))
	if err != nil {
		return
	}
	res = base64.StdEncoding.EncodeToString(rres)
	return
}

func (c FieldCipher) Decrypt(res string, aesKey []byte) (dat []byte, err error) {
	rres, err := base64.StdEncoding.DecodeString(res)
	if err != nil {
		return
	}
	coder := NewAESCoder(aesKey)
	dat, err = coder.AES128CBCDecrypt(rres)
	return
}
