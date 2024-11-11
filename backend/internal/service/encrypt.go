package service

import (
	"crypto/rsa"
	// "fmt"
	"main/internal/service/encrypt"
)

type Data struct {
	WalletPublicKey  string          `json:"walletPublicKey"` // save
	PrivateKeyRSA    *rsa.PrivateKey `json:"privateKeyRSA"`	  // save
	PrivateKeyChaCha []byte          `json:"privateKeyChaCha"` // save
	SaltArgon        []byte          `json:"saltArgon"` // save
	FirstToken       []byte          `json:"firstToken"`
	SecondToken      []byte          `json:"secondToken"`
}

// 加密函數
func Encrypt(walletPrivateKey, walletPublicKey string) (Data, error) {
	var data Data
	data.WalletPublicKey = walletPublicKey

	// AES
	privateKeyAES := []byte(walletPublicKey) // AES-128密鑰，長度必須是16、24或32字節之一
	if len(privateKeyAES) > 32 {
		privateKeyAES = privateKeyAES[:32]
	}
	plainText := []byte(walletPrivateKey)

	// AES加密
	encryptedTextAES, err := encrypt.EncryptAES(plainText, privateKeyAES)
	if err != nil {
		panic(err)
	}
	// fmt.Println("AES加密後的文本:", encryptedTextAES)

	// 複合式加密
	firstText := encryptedTextAES[:(len(encryptedTextAES)+1)/2]
	secondText := encryptedTextAES[(len(encryptedTextAES)+1)/2:]
	// fmt.Println("firstText:", firstText)
	// fmt.Println("secondText:", secondText)

	// RSA
	// 生成 RSA 私鑰和公鑰
	privateKeyRSA, publicKeyRSA, err := encrypt.GenerateRSAKeys(2048)
	if err != nil {
		panic(err)
	}
	data.PrivateKeyRSA = privateKeyRSA

	// fmt.Println("RSA privateKey:", privateKeyRSA)
	// fmt.Println("RSA publicKey:", publicKeyRSA)

	// 要加密的消息
	plainText = []byte(firstText)

	// 使用公鑰加密消息
	encryptedTextRSA, err := encrypt.EncryptRSA(publicKeyRSA, plainText)
	if err != nil {
		panic(err)
	}
	data.FirstToken = encryptedTextRSA
	// fmt.Printf("加密後的密文: %x\n", encryptedTextRSA)

	// ChaCha20
	// 生成ChaCha20密鑰 (32 字節)
	privateKeyChaCha, err := encrypt.GenerateChaChaKey()
	if err != nil {
		panic(err)
	}
	data.PrivateKeyChaCha = privateKeyChaCha
	// fmt.Printf("ChaCha20 Key: %x\n", privateKeyChaCha)

	// 使用 Argon2 生成 ChaCha20 隨機數 (12 字節)
	// 生成16字節長度的隨機鹽, 要存起來
	saltArgon, err := encrypt.GenerateArgonSalt(16)
	if err != nil {
		panic(err)
	}
	data.SaltArgon = saltArgon
	// fmt.Printf("Argon2 salt: %x\n", saltArgon)

	nonce := encrypt.GenerateChaChaNonce(saltArgon)
	// fmt.Printf("ChaCha20 nonce: %x\n", nonce)

	// 原始消息
	plainText = []byte(secondText)

	// 加密消息
	encryptedTextChaCha, err := encrypt.EncryptChaCha(privateKeyChaCha, nonce, plainText)
	if err != nil {
		panic(err)
	}
	data.SecondToken = encryptedTextChaCha
	// fmt.Printf("加密後的密文: %x\n", encryptedTextChaCha)

	return data, err
}

// 解密函數
func Decrypt(data Data) (string, error) {
	// ChaCha20
	nonce := encrypt.GenerateChaChaNonce(data.SaltArgon)
	// fmt.Printf("ChaCha20 nonce: %x\n", nonce)

	// 解密消息
	decrypted, err := encrypt.DecryptChaCha(data.PrivateKeyChaCha, nonce, data.SecondToken)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("解密後的明文: %s\n", string(decrypted))

	// RSA
	// 使用私鑰解密消息
	decryptedTextRSA, err := encrypt.DecryptRSA(data.PrivateKeyRSA, data.FirstToken)
	if err != nil {
		panic(err)
	}
	// fmt.Println("解密後的明文:", string(decryptedTextRSA))
	encryptedTextAES := string(decryptedTextRSA)+string(decrypted)
	// fmt.Println("解密後的明文:", encryptedTextAES)

	// AES解密
	privateKeyAES := []byte(data.WalletPublicKey) // AES-128密鑰，長度必須是16、24或32字節之一
	if len(privateKeyAES) > 32 {
		privateKeyAES = privateKeyAES[:32]
	}
	decryptedTextAES, err := encrypt.DecryptAES(encryptedTextAES, privateKeyAES)
	if err != nil {
		panic(err)
	}
	// fmt.Println("AES解密後的文本:", decryptedTextAES)

	return decryptedTextAES, err
}

// func main() {
// 	// AES
// 	privateKeyAES := []byte(walletPublicKey) // AES-128密鑰，長度必須是16、24或32字節之一
// 	if len(privateKeyAES) > 32 {
// 		privateKeyAES = privateKeyAES[:32]
// 	}
// 	plainText := []byte(walletPrivateKey)

// 	// AES加密
// 	encryptedTextAES, err := encrypt.EncryptAES(plainText, privateKeyAES)
// 	if err != nil {
// 		fmt.Println("AES加密失敗:", err)
// 		return
// 	}
// 	fmt.Println("AES加密後的文本:", encryptedTextAES)

// 	// AES解密
// 	decryptedTextAES, err := encrypt.DecryptAES(encryptedTextAES, privateKeyAES)
// 	if err != nil {
// 		fmt.Println("AES解密失敗:", err)
// 		return
// 	}
// 	fmt.Println("AES解密後的文本:", decryptedTextAES)

// 	// 複合式加密
// 	firstText := encryptedTextAES[:(len(encryptedTextAES)+1)/2]
// 	secondText := encryptedTextAES[(len(encryptedTextAES)+1)/2:]
// 	fmt.Println("firstText:", firstText)
// 	fmt.Println("secondText:", secondText)

// 	// RSA
// 	// 生成 RSA 私鑰和公鑰
// 	privateKeyRSA, publicKeyRSA, err := GenerateRSAKeys(2048)
// 	if err != nil {
// 		fmt.Println("密鑰生成失敗:", err)
// 		return
// 	}

// 	fmt.Println("RSA privateKey:", privateKeyRSA)
// 	fmt.Println("RSA publicKey:", publicKeyRSA)

// 	// 要加密的消息
// 	plainText = []byte(firstText)

// 	// 使用公鑰加密消息
// 	encryptedTextRSA, err := EncryptRSA(publicKeyRSA, plainText)
// 	if err != nil {
// 		fmt.Println("加密失敗:", err)
// 		return
// 	}
// 	// firstToken := firstText
// 	fmt.Printf("加密後的密文: %x\n", encryptedTextRSA)

// 	// 使用私鑰解密消息
// 	decryptedTextRSA, err := DecryptRSA(privateKeyRSA, encryptedTextRSA)
// 	if err != nil {
// 		fmt.Println("解密失敗:", err)
// 		return
// 	}
// 	fmt.Println("解密後的明文:", string(decryptedTextRSA))

// 	// ChaCha20
// 	// 生成ChaCha20密鑰 (32 字節)
// 	privateKeyChaCha, err := GenerateChaChaKey()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("ChaCha20 Key: %x\n", privateKeyChaCha)

// 	// 使用 Argon2 生成 ChaCha20 隨機數 (12 字節)
// 	// 生成16字節長度的隨機鹽
// 	saltArgon, err := GenerateArgonSalt(16)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Printf("Argon2 salt: %x\n", saltArgon)

// 	nonce := GenerateChaChaNonce(saltArgon)
// 	fmt.Printf("ChaCha20 nonce: %x\n", nonce)

// 	// 原始消息
// 	plainText = []byte(secondText)

// 	// 加密消息
// 	encrypted, err := EncryptChaCha(privateKeyChaCha, nonce, plainText)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("加密後的密文: %x\n", encrypted)

// 	nonce = GenerateChaChaNonce(saltArgon)
// 	fmt.Printf("ChaCha20 nonce: %x\n", nonce)

// 	// 解密消息
// 	decrypted, err := DecryptChaCha(privateKeyChaCha, nonce, encrypted)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("解密後的明文: %s\n", string(decrypted))
// }
