package service

import (
	// "main/internal/models"
	"main/internal/repositories"

	// "github.com/google/uuid"
)

type UserWalletService struct {
	repo repositories.UserWalletRepository
}

// NewUserWalletService 創建新的 UserWalletService
func NewUserWalletService() *UserWalletService {
	return &UserWalletService{
		repo: repositories.NewUserWalletRepository(),
	}
}

// // CreateUserWallet 處理創建 UserWallet 的業務邏輯
// func (s *UserWalletService) CreateUserWallet(id *uuid.UUID, walletAddress *string, data *WalletData) error {
// 	var userWallet models.UserWallet
// 	userWallet.UserID = *id
// 	userWallet.WalletAddress = *walletAddress
// 	userWallet.WalletPublicKey = data.WalletPublicKey
// 	userWallet.PrivateKeyRSA = data.PrivateKeyRSA
// 	userWallet.WalletPublicKey = data.WalletPublicKey
// 	userWallet.WalletPublicKey = data.WalletPublicKey
// 	return s.repo.Create(&userWallet)
// }

// // GetUserWalletByID 根據 ID 獲取用戶錢包
// func (s *UserWalletService) GetUserWalletByID(id uuid.UUID) (*models.UserWallet, error) {
// 	return s.repo.FindByID(id)
// }
