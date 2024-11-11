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
	// fmt.Println("當前時間 (yyyymmddhhmm):", formattedTime)
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

// func main() {
// 	// 生成密鑰 (32 字節)
// 	key, err := GenerateChaChaKey()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("ChaCha20 Key: %x\n", key)

// 	// 使用 Argon2 生成 ChaCha20 隨機數 (12 字節)
// 	// 生成16字節長度的隨機鹽
// 	salt, err := GenerateArgonSalt(16)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Printf("Argon2 salt: %x\n", salt)

// 	nonce := GenerateChaChaNonce(salt)
// 	fmt.Printf("ChaCha20 nonce: %x\n", nonce)

// 	// 原始消息
// 	plainText := []byte("cwFWely7l60tkrbEm7Qm+hgSaNWJ5JExPB5U+VL8wYYkERsTtw9y7g==")

// 	// 加密消息
// 	encrypted, err := EncryptChaCha(key, nonce, plainText)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("加密後的密文: %x\n", encrypted)

// 	nonce = GenerateChaChaNonce(salt)
// 	fmt.Printf("ChaCha20 nonce: %x\n", nonce)

// 	// 解密消息
// 	decrypted, err := DecryptChaCha(key, nonce, encrypted)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("解密後的明文: %s\n", string(decrypted))
// }
