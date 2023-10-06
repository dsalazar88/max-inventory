package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

const key = "01234567890123456789012345678901"

var (
	ErrDecript = errors.New("ciphertext too short")
)

// The Encrypt function takes a plaintext byte array, encrypts it using AES encryption with a randomly
// generated nonce, and returns the encrypted ciphertext.
func Encrypt(plaintext []byte) ([]byte, error) {
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil

}

// The Decrypt function takes a ciphertext and uses AES encryption to decrypt it and return the
// plaintext.
func Decrypt(ciphertext []byte) ([]byte, error) {
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, ErrDecript
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	return gcm.Open(nil, nonce, ciphertext, nil)

}

// The function "ToBase64" encodes a byte array into a base64 string using the RawStdEncoding.
func ToBase64(plaintext []byte) string {
	return base64.RawStdEncoding.EncodeToString(plaintext)
}

// The function "FromBase64" decodes a base64 encoded string into a byte slice.
func FromBase64(ciphertext string) ([]byte, error) {
	return base64.RawStdEncoding.DecodeString(ciphertext)
}
