package pkg

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ChainlinkPriceFeed(contractAddress string) (string, error) {
	// 連接到以太坊節點
	client, err := ethclient.Dial(AlchemySepolia)
	if err != nil {
		return "", fmt.Errorf("failed to connect to Ethereum client: %v", err)
	}

	// 讀取 ABI 文件
	abiPath := "./internal/ABI/chainlinkPriceFeed.json"
	file, err := os.ReadFile(abiPath)
	if err != nil {
		return "", fmt.Errorf("failed to read ABI file: %v", err)
	}

	// 解析 ABI
	contractAbi, err := abi.JSON(strings.NewReader(string(file)))
	if err != nil {
		return "", fmt.Errorf("failed to parse ABI: %v", err)
	}

	// 列出所有可讀取的 view 函數
	fmt.Println("可讀取的函數:")
	for _, method := range contractAbi.Methods {
		if method.StateMutability == "view" {
			fmt.Printf("函數名稱: %s, 輸入參數: %v, 輸出參數: %v\n", method.Name, method.Inputs, method.Outputs)
		}
	}

	// 呼叫 `latestAnswer` 函數, 取得 BTC/USD 最新價格
	answer, err := callLatestAnswer(client, contractAbi, contractAddress)
	if err != nil {
		return "", fmt.Errorf("failed to call latestAnswer: %v", err)
	}

	fmt.Printf("最新的價格回應: %s\n", answer)
	return answer, nil
}

// 調用 latestAnswer 函數取得最新價格
func callLatestAnswer(client *ethclient.Client, contractAbi abi.ABI, contractAddress string) (string, error) {
	// 轉換合約地址
	contractAddr := common.HexToAddress(contractAddress)

	// 準備 latestAnswer 函數的 data
	latestAnswer, err := contractAbi.Pack("latestAnswer")
	if err != nil {
		return "", fmt.Errorf("failed to pack latestAnswer call data: %v", err)
	}

	// 使用 CallContract 調用智能合約
	msg := ethereum.CallMsg{
		To:   &contractAddr,
		Data: latestAnswer,
	}

	// 呼叫合約，讀取結果
	result, err := client.CallContract(context.Background(), msg, nil) // nil 表示最新區塊
	if err != nil {
		return "", fmt.Errorf("failed to call contract: %v", err)
	}

	// 解析回傳結果
	var answer *big.Int
	err = contractAbi.UnpackIntoInterface(&answer, "latestAnswer", result)
	if err != nil {
		return "", fmt.Errorf("failed to unpack result: %v", err)
	}

	return answer.String(), nil
}
