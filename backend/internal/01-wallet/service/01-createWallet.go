package service

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// 此為範例
func CreateWallet() (string, string, string, error) {

	// 產出私鑰
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		// log.Fatalf("error generate private key: %v", err)
		return "", "", "", fmt.Errorf("error generate private key: %v", err)
	}

	// 從私鑰產出公鑰
	publicKey := privateKey.Public()
	// fmt.Println("publicKey : ", publicKey)
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		// log.Fatalf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return "", "", "", fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	// fmt.Println("publicKeyECDSA : ", publicKeyECDSA)
	// 從公鑰產生錢包地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	// fmt.Println("address :", address)
	// // 私鑰轉成 hex格式 ( 0x 開頭)
	// privateKeyByte := crypto.FromECDSA(privateKey)
	// privateKeyHex := hexutil.Encode(privateKeyByte[2:])

	// 私鑰轉成 ethereum 私鑰標準長度,  64位的16進制string 可直接 import metamask
	privateKeyByte := crypto.FromECDSA(privateKey)
	privateKeyHex := hexutil.Encode(privateKeyByte)
	// 公鑰轉成 hex格式
	publicKeyByte := crypto.FromECDSAPub(publicKeyECDSA)
	publicKeyHex := hexutil.Encode(publicKeyByte[2:])

	return privateKeyHex, publicKeyHex, address, nil
}
