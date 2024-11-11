package models

import (
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"crypto/rsa"
)

// UserWallet 表示用戶的錢包數據結構
type UserWallet struct {
	gorm.Model                 // 包含 ID, CreatedAt, UpdatedAt, DeletedAt
	UserID           uuid.UUID `gorm:"column:user_id;type:uuid;unique;not null" json:"user_id" validate:"required"`
	WalletPublicKey  string    `gorm:"column:walletPublicKey;type:VARCHAR(512);not null" json:"walletPublicKey" validate:"required"` // 公鑰
	PrivateKeyRSA    string    `gorm:"column:privateKeyRSA;type:VARCHAR(512);not null" json:"privateKeyRSA" validate:"required"`     // RSA 私鑰
	WalletAddress    string    `gorm:"column:walletAddress;type:VARCHAR(256);not null" json:"walletAddress" validate:"required"`     // 錢包地址
	PrivateKeyChaCha []byte    `gorm:"column:private_key_chacha;type:blob;not null" json:"private_key_chacha" validate:"required"`   // ChaCha 私鑰
	SaltArgon        []byte    `gorm:"column:salt_argon;type:blob;not null" json:"salt_argon" validate:"required"`                   // Argon2 的 salt
	FirstToken       []byte    `gorm:"column:first_token;type:blob;not null" json:"first_token" validate:"required"`                 // 第一個 Token
	SecondToken      []byte    `gorm:"column:second_token;type:blob;not null" json:"second_token" validate:"required"`               // 第二個 Token
}

func ValidateUserWallet(table *UserWallet) error {
	validate := validator.New()
	return validate.Struct(table)
}

type WalletData struct {
	WalletPublicKey  string          `json:"walletPublicKey"`
	PrivateKeyRSA    *rsa.PrivateKey `json:"privateKeyRSA"`
	PrivateKeyChaCha []byte          `json:"privateKeyChaCha"`
	SaltArgon        []byte          `json:"saltArgon"`
	FirstToken       []byte          `json:"firstToken"`
	SecondToken      []byte          `json:"secondToken"`
}
