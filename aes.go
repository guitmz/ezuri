package main

import (
	"crypto/aes"
	"crypto/cipher"
)

func aesEnc(srcBytes []byte, key string, iv string) []byte {
	block, err := aes.NewCipher([]byte(key))
	check(err)

	encrypter := cipher.NewCFBEncrypter(block, []byte(iv))
	encrypted := make([]byte, len(srcBytes))
	encrypter.XORKeyStream(encrypted, srcBytes)
	return encrypted
}
