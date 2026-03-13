/*
*  Copyright (c) WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
*
*  WSO2 Inc. licenses this file to you under the Apache License,
*  Version 2.0 (the "License"); you may not use this file except
*  in compliance with the License.
*  You may obtain a copy of the License at
*
*    http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing,
* software distributed under the License is distributed on an
* "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
* KIND, either express or implied.  See the License for the
* specific language governing permissions and limitations
* under the License.
 */

package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
)

const (
	AES256KeySize    = 32
	AES256HexKeySize = 64
	GCMIVSize        = 128
	GCMTagSize       = 128
	hexCharacters    = "0123456789abcdefABCDEF"
)

type gcmCipherText struct {
	CipherText string `json:"cipherText"`
	IV         string `json:"iv"`
}

// Returns md5 hash of a given string
func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

// Encrypt string to base64 crypto using AES
func Encrypt(key []byte, text string) string {
	// key := []byte(keyText)
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		HandleErrorAndExit("Error in encryption", err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		HandleErrorAndExit("Error in encryption", err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(ciphertext)
}

// Decrypt from base64 to decrypted string
func Decrypt(key []byte, cryptoText string) string {
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher(key)
	if err != nil {
		HandleErrorAndExit("Error in decryption", err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		HandleErrorAndExit("Error in Decryption: Ciphertext too short", nil)
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext)
}

// ResolveAES256Key validates and converts a direct encryption key into raw key bytes.
func ResolveAES256Key(encryptionKey string) ([]byte, error) {
	trimmedKey := strings.TrimSpace(encryptionKey)
	if trimmedKey == "" {
		return nil, errors.New("Encryption key cannot be empty")
	}
	if isHexAES256Key(trimmedKey) {
		keyBytes, err := hex.DecodeString(trimmedKey)
		if err != nil {
			return nil, errors.New("Invalid hexadecimal characters found in encryption key")
		}
		return keyBytes, nil
	}
	keyBytes := []byte(trimmedKey)
	if len(keyBytes) != AES256KeySize {
		return nil, fmt.Errorf("Invalid AES key length: %d bytes. AES-256 requires a 32-byte (256-bit) key", len(keyBytes))
	}
	return keyBytes, nil
}

// EncryptAES256 encrypts plain text using AES-256 GCM and returns a self-contained ciphertext.
func EncryptAES256(key []byte, text string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	aesGCM, err := cipher.NewGCMWithNonceSize(block, GCMIVSize)
	if err != nil {
		return "", err
	}

	iv := make([]byte, GCMIVSize)
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	ciphertext := aesGCM.Seal(nil, iv, []byte(text), nil)
	payload, err := json.Marshal(gcmCipherText{
		CipherText: base64.StdEncoding.EncodeToString(ciphertext),
		IV:         base64.StdEncoding.EncodeToString(iv),
	})
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(payload), nil
}

// DecryptAES256 decrypts self-contained AES-256 GCM ciphertext into plain text.
func DecryptAES256(key []byte, cryptoText string) (string, error) {
	payload, err := base64.StdEncoding.DecodeString(strings.TrimSpace(cryptoText))
	if err != nil {
		return "", err
	}

	var encryptedValue gcmCipherText
	if err = json.Unmarshal(payload, &encryptedValue); err != nil {
		return "", err
	}

	ciphertext, err := base64.StdEncoding.DecodeString(encryptedValue.CipherText)
	if err != nil {
		return "", err
	}
	iv, err := base64.StdEncoding.DecodeString(encryptedValue.IV)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	aesGCM, err := cipher.NewGCMWithNonceSize(block, GCMIVSize)
	if err != nil {
		return "", err
	}

	plainText, err := aesGCM.Open(nil, iv, ciphertext, nil)
	if err != nil {
		return "", err
	}
	return string(plainText), nil
}

func isHexAES256Key(encryptionKey string) bool {
	if len(encryptionKey) != AES256HexKeySize {
		return false
	}
	for _, char := range encryptionKey {
		if !strings.ContainsRune(hexCharacters, char) {
			return false
		}
	}
	return true
}
