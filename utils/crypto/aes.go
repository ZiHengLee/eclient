package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

var (
	DEFAULT_IV = []byte("0102030405060708")
)

var (
	ErrIncorrectPadding = errors.New("padding incorrect")
)

type AESCoder struct {
	key []byte
	iv  []byte
}

func NewAESCoder(key []byte) *AESCoder {
	return &AESCoder{
		key: key,
		iv:  DEFAULT_IV,
	}
}

func (coder *AESCoder) AES128CBCEncrypt(raw []byte) ([]byte, error) {
	return coder.aes128CBCEncrypt(raw, coder.key, coder.iv)
}

func (coder *AESCoder) AES128CBCDecrypt(crypted []byte) ([]byte, error) {
	return coder.aes128CBCDecrypt(crypted, coder.key, coder.iv)
}

func (coder *AESCoder) aes128CBCEncrypt(raw, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err

	}

	blockSize := block.BlockSize()
	raw = coder.PKCS5Padding(raw, blockSize)
	if iv == nil {
		iv = key[:blockSize]
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(raw))
	mode.CryptBlocks(crypted, raw)
	return crypted, nil
}

func (coder *AESCoder) aes128CBCDecrypt(crypted, key, iv []byte) (raw []byte, err error) {

	defer func() {
		if err := recover(); err != nil {
			raw = nil
			err = errors.New("aes128CBCDecrypt error")
		}
	}()

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	if iv == nil {
		iv = key[:blockSize]
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	raw = make([]byte, len(crypted))
	mode.CryptBlocks(raw, crypted)
	raw, err = coder.PKCS5UnPadding(raw, blockSize)
	if err != nil {
		return nil, err
	}
	return raw, nil
}

func (coder *AESCoder) PKCS5Padding(raw []byte, blockSize int) []byte {
	padding := blockSize - len(raw)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(raw, padtext...)
}

func (coder *AESCoder) PKCS5UnPadding(raw []byte, blockSize int) ([]byte, error) {
	dlen := len(raw)
	if dlen == 0 || dlen%blockSize != 0 {
		return nil, ErrIncorrectPadding
	}
	last := int(raw[dlen-1])
	if dlen < last {
		return nil, ErrIncorrectPadding
	}
	if last == 0 || last > blockSize {
		return nil, ErrIncorrectPadding
	}
	for _, val := range raw[dlen-last:] {
		if int(val) != last {
			return nil, ErrIncorrectPadding
		}
	}
	return raw[:dlen-last], nil
}
