package repositories

import (
	"main/internal/models"
	configs "main/internal/shared/utils/init"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"fmt"
	// "crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

type UserWalletRepository interface {
	Create(id *uuid.UUID, walletAddress *string, data *models.WalletData) error
	FindByID(id uuid.UUID) (*models.UserWallet, error)
	Update(userWallet *models.UserWallet) error
	Delete(id uuid.UUID) error
}

type userWalletRepository struct {
	db *gorm.DB
}

// NewUserWalletRepository 使用全局的資料庫引擎實例
func NewUserWalletRepository() UserWalletRepository {
	return &userWalletRepository{
		db: configs.Global.Sqlite.Engine, // 使用全局資料庫
	}
}

// Create 創建新的 UserWallet
func (r *userWalletRepository) Create(id *uuid.UUID, walletAddress *string, data *models.WalletData) error {
	
	privKeyPEM, err := privateKeyToPEM(data.PrivateKeyRSA)
	if err != nil {
			return fmt.Errorf("private key to pem err: %v", err)
	}

	userWallet := models.UserWallet{
		UserID: *id,
		WalletAddress: *walletAddress,
		WalletPublicKey: data.WalletPublicKey,
		PrivateKeyRSA: privKeyPEM,
		PrivateKeyChaCha: data.PrivateKeyChaCha,
		SaltArgon: data.SaltArgon,
		FirstToken: data.FirstToken,
		SecondToken: data.SecondToken,
	}
	return r.db.Create(&userWallet).Error
}

// FindByID 根據 ID 查找 UserWallet
func (r *userWalletRepository) FindByID(id uuid.UUID) (*models.UserWallet, error) {
	var userWallet models.UserWallet
	if err := r.db.First(&userWallet, "userID = ?", id).Error; err != nil {
		return nil, err
	}
	return &userWallet, nil
}

// Update 更新 UserWallet
func (r *userWalletRepository) Update(userWallet *models.UserWallet) error {
	return r.db.Save(userWallet).Error
}

// Delete 刪除 UserWallet
func (r *userWalletRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.UserWallet{}, "userID = ?", id).Error
}


func privateKeyToPEM(priv *rsa.PrivateKey) (string, error) {
	// 將私鑰轉換為 ASN.1 DER 編碼
	privDER := x509.MarshalPKCS1PrivateKey(priv)

	// 創建 PEM 區塊
	privPEM := pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privDER,
	}

	// 將 PEM 區塊編碼為字節數組，並返回字符串
	return string(pem.EncodeToMemory(&privPEM)), nil
}