package encrypt

import (
	"crypto/rand"
	"fmt"
	"time"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/chacha20"
)

// 生成密鑰 (32 字節)
func GenerateChaChaKey() ([]byte, error) {
	key := make([]byte, chacha20.KeySize)
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}
	return key, nil
}

// 生成指定長度的隨機鹽
func GenerateArgonSalt(length int) ([]byte, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, fmt.Errorf("生成鹽失敗: %w", err)
	}
	return salt, nil
}

// 使用 Argon2 生成 ChaCha20 隨機數 (12 字節)
func GenerateChaChaNonce(salt []byte) []byte {
	// 抓取當下時間
	currentTime := time.Now()

	// 格式化為 yyyymmddhhmm 格式
	// Go 的時間格式化是基於特定的參考時間 "Mon Jan 2 15:04:05 MST 2006"
	formattedTime := currentTime.Format("200601021504")[:11]

	// 打印格式化後的時間
	return argon2.Key([]byte(formattedTime), salt, 1, 64*1024, 4, 12) // 12字節的密鑰用於ChaCha20 隨機數
}

// 使用 ChaCha20 加密
func EncryptChaCha(key, nonce, plainText []byte) ([]byte, error) {
	encrypted := make([]byte, len(plainText))
	stream, err := chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		return nil, err
	}
	stream.XORKeyStream(encrypted, plainText)
	return encrypted, nil
}

// 使用 ChaCha20 解密
func DecryptChaCha(key, nonce, cipherText []byte) ([]byte, error) {
	decrypted := make([]byte, len(cipherText))
	stream, err := chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		return nil, err
	}
	stream.XORKeyStream(decrypted, cipherText)
	return decrypted, nil
}