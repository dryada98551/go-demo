package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

// 加密函數
func EncryptAES(plainText, key []byte) (string, error) {
	// 創建AES區塊
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 初始化向量 (IV)
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// 使用CTR模式進行加密
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	// 將加密結果轉換為Base64字串
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// 解密函數
func DecryptAES(cipherTextBase64 string, key []byte) (string, error) {
	cipherText, err := base64.StdEncoding.DecodeString(cipherTextBase64)
	if err != nil {
		return "", err
	}

	// 創建AES區塊
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 讀取初始化向量 (IV)
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	// 使用CTR模式進行解密
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), nil
}
