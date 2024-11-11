package service

import (
	"fmt"
	// "crypto/rsa"
	"main/internal/shared/utils/encrypt"
	"main/internal/models"
)

// type WalletData struct {
// 	WalletPublicKey  string          `json:"walletPublicKey"`
// 	PrivateKeyRSA    *rsa.PrivateKey `json:"privateKeyRSA"`
// 	PrivateKeyChaCha []byte          `json:"privateKeyChaCha"`
// 	SaltArgon        []byte          `json:"saltArgon"`
// 	FirstToken       []byte          `json:"firstToken"`
// 	SecondToken      []byte          `json:"secondToken"`
// }

// 加密函數
func Encrypt(walletPrivateKey, walletPublicKey *string) (models.WalletData, error) {
	var data models.WalletData
	data.WalletPublicKey = *walletPublicKey

	// AES
	privateKeyAES := []byte(*walletPublicKey) // AES-128密鑰，長度必須是16、24或32字節之一
	if len(privateKeyAES) > 32 {
		privateKeyAES = privateKeyAES[:32]
	}
	plainText := []byte(*walletPrivateKey)

	// AES加密
	encryptedTextAES, err := encrypt.EncryptAES(plainText, privateKeyAES)
	if err != nil {
		return data, fmt.Errorf("error encryp private key: %v", err)
	}

	// 複合式加密
	firstText := encryptedTextAES[:(len(encryptedTextAES)+1)/2]
	secondText := encryptedTextAES[(len(encryptedTextAES)+1)/2:]

	// RSA
	// 生成 RSA 私鑰和公鑰
	privateKeyRSA, publicKeyRSA, err := encrypt.GenerateRSAKeys(2048)
	if err != nil {
		return data, fmt.Errorf("error encryp private key: %v", err)
	}
	data.PrivateKeyRSA = privateKeyRSA

	// 要加密的消息
	plainText = []byte(firstText)

	// 使用公鑰加密消息
	encryptedTextRSA, err := encrypt.EncryptRSA(publicKeyRSA, plainText)
	if err != nil {
		return data, fmt.Errorf("error encryp private key: %v", err)
	}
	data.FirstToken = encryptedTextRSA

	// ChaCha20
	// 生成ChaCha20密鑰 (32 字節)
	privateKeyChaCha, err := encrypt.GenerateChaChaKey()
	if err != nil {
		return data, fmt.Errorf("error encryp private key: %v", err)
	}
	data.PrivateKeyChaCha = privateKeyChaCha

	// 使用 Argon2 生成 ChaCha20 隨機數 (12 字節)
	// 生成16字節長度的隨機鹽
	saltArgon, err := encrypt.GenerateArgonSalt(16)
	if err != nil {
		return data, fmt.Errorf("error encryp private key: %v", err)
	}
	data.SaltArgon = saltArgon

	nonce := encrypt.GenerateChaChaNonce(saltArgon)

	// 原始消息
	plainText = []byte(secondText)

	// 加密消息
	encryptedTextChaCha, err := encrypt.EncryptChaCha(privateKeyChaCha, nonce, plainText)
	if err != nil {
		return data, fmt.Errorf("error encryp private key: %v", err)
	}
	data.SecondToken = encryptedTextChaCha

	return data, err
}

// 解密函數
func Decrypt(data *models.WalletData) (string, error) {
	// ChaCha20
	nonce := encrypt.GenerateChaChaNonce(data.SaltArgon)

	// 解密消息
	decrypted, err := encrypt.DecryptChaCha(data.PrivateKeyChaCha, nonce, data.SecondToken)
	if err != nil {
		return "", fmt.Errorf("error decrypt private key: %v", err)
	}

	// RSA
	// 使用私鑰解密消息
	decryptedTextRSA, err := encrypt.DecryptRSA(data.PrivateKeyRSA, data.FirstToken)
	if err != nil {
		return "", fmt.Errorf("error decrypt private key: %v", err)
	}
	encryptedTextAES := string(decryptedTextRSA)+string(decrypted)

	// AES解密
	privateKeyAES := []byte(data.WalletPublicKey) // AES-128密鑰，長度必須是16、24或32字節之一
	if len(privateKeyAES) > 32 {
		privateKeyAES = privateKeyAES[:32]
	}
	decryptedTextAES, err := encrypt.DecryptAES(encryptedTextAES, privateKeyAES)
	if err != nil {
		return "", fmt.Errorf("error decrypt private key: %v", err)
	}

	return decryptedTextAES, err
}