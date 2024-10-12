package callback

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

var Key []byte

// PKCS7Padding implements padding for AES block size (16 bytes)
func PKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// AES-CBC 加密
func AesEncrypt(plainText, key []byte) (string, error) {
	// Create a new AES cipher block with the given key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Generate an IV (initialization vector)
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// Add padding to the plaintext to make it a multiple of the block size
	plainText = PKCS7Padding(plainText, aes.BlockSize)

	// Create the AES-CBC encrypter
	mode := cipher.NewCBCEncrypter(block, iv)

	// Allocate memory for the ciphertext
	cipherText := make([]byte, len(plainText))

	// Perform the encryption
	mode.CryptBlocks(cipherText, plainText)

	// Prepend the IV to the ciphertext
	cipherText = append(iv, cipherText...)

	// Encode the ciphertext in Base64 for easier transport/storage
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// PKCS7UnPadding removes padding after decryption
func PKCS7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("invalid ciphertext size")
	}

	// Get the value of the last byte which is the padding size
	padding := int(data[length-1])
	if padding > length {
		return nil, errors.New("invalid padding size")
	}

	// Remove the padding bytes
	return data[:length-padding], nil
}

// AES-CBC 解密
func AesDecrypt(cipherText string, key []byte) (string, error) {
	// Base64 decode the cipher text
	cipherData, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	// The IV is the first 16 bytes
	if len(cipherData) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := cipherData[:aes.BlockSize]
	cipherData = cipherData[aes.BlockSize:]

	// Create the AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Create the AES-CBC decrypter
	mode := cipher.NewCBCDecrypter(block, iv)

	// Allocate memory for the decrypted data
	plainText := make([]byte, len(cipherData))

	// Perform the decryption
	mode.CryptBlocks(plainText, cipherData)

	// Remove padding from the decrypted data
	plainText, err = PKCS7UnPadding(plainText)
	if err != nil {
		return "", err
	}

	// Convert the plaintext to a string and return it
	return string(plainText), nil
}
