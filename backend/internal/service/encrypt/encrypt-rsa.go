package encrypt

import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/sha256"
)

// 生成 RSA 密鑰對
func GenerateRSAKeys(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
    privateKey, err := rsa.GenerateKey(rand.Reader, bits)
    if err != nil {
        return nil, nil, err
    }
    publicKey := &privateKey.PublicKey
    return privateKey, publicKey, nil
}

// 使用公鑰加密消息
func EncryptRSA(publicKey *rsa.PublicKey, message []byte) ([]byte, error) {
    cipherText, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, message, nil)
    if err != nil {
        return nil, err
    }
    return cipherText, nil
}

// 使用私鑰解密消息
func DecryptRSA(privateKey *rsa.PrivateKey, cipherText []byte) ([]byte, error) {
    decryptedMessage, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, cipherText, nil)
    if err != nil {
        return nil, err
    }
    return decryptedMessage, nil
}

// func main() {
//     // 生成 RSA 私鑰和公鑰
//     privateKey, publicKey, err := GenerateRSAKeys(2048)
//     if err != nil {
//         fmt.Println("密鑰生成失敗:", err)
//         return
//     }

// 		fmt.Println("privateKey:", privateKey)
// 		fmt.Println("publicKey:", publicKey)

//     // 要加密的消息
//     message := []byte("0x1ee9348031d04e02648b60d99d25e7648f2928ad265bff35025afa2562997278")

//     // 使用公鑰加密消息
//     cipherText, err := EncryptRSA(publicKey, message)
//     if err != nil {
//         fmt.Println("加密失敗:", err)
//         return
//     }
//     fmt.Printf("加密後的密文: %x\n", cipherText)

//     // 使用私鑰解密消息
//     decryptedMessage, err := DecryptRSA(privateKey, cipherText)
//     if err != nil {
//         fmt.Println("解密失敗:", err)
//         return
//     }
//     fmt.Println("解密後的明文:", string(decryptedMessage))
// }
