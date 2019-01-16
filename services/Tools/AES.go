package Tools

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

type AESlib struct {
	Key []byte
}

func NewAESlib() *AESlib {
	AESkey := make([]byte, 16)
	rand.Read(AESkey)
	return &AESlib{Key: AESkey}
}
func NewAESlibWithKey(key string) *AESlib {
	AESkey := make([]byte, 16)
	AESkey = []byte(key)
	return &AESlib{Key: AESkey}
}
func Encrypt(key []byte, text string) (string, error) {
	plaintext := []byte(text)
	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	out := gcm.Seal(nonce, nonce, plaintext, nil)
	r := base64.RawStdEncoding.EncodeToString(out)
	return r, nil
}

func Decrypt(key []byte, encriptedbase64Text string) (string, error) {
	ciphertext, err := base64.RawStdEncoding.DecodeString(encriptedbase64Text)
	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	r, er := gcm.Open(nil, nonce, ciphertext, nil)
	return BytesToString(r), er
}
func BytesToString(data []byte) string {
	return string(data[:])
}
