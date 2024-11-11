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

// func main() {
// 	key := []byte("0xffbd5decd376374366e4ebefc7c5cae7d469358b1c7d8223e7f44b56c3f811eacb0338752e5dd439ff5f6ac2bc2783aa162c112a21427662ed10c408c6c230") // AES-128密鑰，長度必須是16、24或32字節之一
// 	if len(key) > 32 {
// 		key = key[:32]
// 	}
// 	plainText := "0x062090fe4c33de35a840d3638ec8610965cfc5d477039a1b84a7c6e899295d29"
	
// 	// 加密
// 	encryptedText, err := EncryptAES([]byte(plainText), key)
// 	if err != nil {
// 		fmt.Println("加密失敗:", err)
// 		return
// 	}
// 	fmt.Println("加密後的文本:", encryptedText)

// 	// 解密
// 	decryptedText, err := DecryptAES(encryptedText, key)
// 	if err != nil {
// 		fmt.Println("解密失敗:", err)
// 		return
// 	}
// 	fmt.Println("解密後的文本:", decryptedText)
// }
