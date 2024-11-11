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