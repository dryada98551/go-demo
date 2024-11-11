package pkg

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func MintUSDT(contractAddress, privateKeyHex string, mintToAddress string, amount *big.Int) (string, error) {
	// 連接到以太坊節點
	client, err := ethclient.Dial(AlchemySepolia)
	if err != nil {
		return "", fmt.Errorf("failed to connect to Ethereum client: %v", err)
	}

	// 讀取 ABI 文件
	abiPath := "./internal/ABI/ERC20.json"
	file, err := os.ReadFile(abiPath)
	if err != nil {
		return "", fmt.Errorf("failed to read ABI file: %v", err)
	}

	// 解析 ABI
	contractAbi, err := abi.JSON(strings.NewReader(string(file)))
	if err != nil {
		return "", fmt.Errorf("failed to parse ABI: %v", err)
	}

	// 轉換合約地址和接收者地址
	contractAddr := common.HexToAddress(contractAddress)
	mintTo := common.HexToAddress(mintToAddress)

	// 將私鑰轉換為 ECDSA 格式
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return "", fmt.Errorf("failed to convert private key: %v", err)
	}

	// 從私鑰生成公鑰地址
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("failed to get public key from private key")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 獲取最新的 nonce 值（交易計數器）
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", fmt.Errorf("failed to get nonce: %v", err)
	}

	// 設置 gas 限制和 gas 價格
	gasLimit := uint64(200000) // 根據合約的 gas 限制設置
	gasPriceSuggest, err := client.SuggestGasPrice(context.Background())
	fmt.Println("Suggest gas price: ", gasPriceSuggest)
	if err != nil {
		return "", fmt.Errorf("failed to suggest gas price: %v", err)
	}

	gasPrice := new(big.Int)
	gasPrice.Mul(big.NewInt(350), big.NewInt(1e9)) // 固定的 gasPrice (GWWei)

	// 構建交易數據，調用合約的 mint 函數
	mintData, err := contractAbi.Pack("mint", mintTo, amount)
	if err != nil {
		return "", fmt.Errorf("failed to pack mint function data: %v", err)
	}

	// 構建交易對象
	tx := types.NewTransaction(nonce, contractAddr, big.NewInt(0), gasLimit, gasPrice, mintData)

	// 使用私鑰對交易進行簽名
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", fmt.Errorf("failed to get network ID: %v", err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign transaction: %v", err)
	}

	// 發送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", fmt.Errorf("failed to send transaction: %v", err)
	}

	txHash := signedTx.Hash().Hex()
	fmt.Printf("Transaction sent: %s\n", txHash)

	return txHash, nil
}
