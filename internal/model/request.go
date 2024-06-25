package model

type Request struct {
	Model
	CreatedAt           string `gorm:"type:varchar(255);not null" json:"createdAt" immutable:"true"`
	UserAddress         string `gorm:"type:varchar(255);not null" json:"userAddress" binding:"required" immutable:"true"`
	Nonce               string `gorm:"type:varchar(255);not null" json:"nonce" binding:"required" immutable:"true"`
	Name                string `gorm:"type:varchar(255);not null" json:"name" binding:"required" immutable:"true"`
	InputCount          string `gorm:"type:varchar(255);not null" json:"inputCount" binding:"required" immutable:"true"`
	PreviousOutputCount string `gorm:"type:varchar(255);not null" json:"previousOutputCount" binding:"required" immutable:"true"`
	PreviousSignature   string `gorm:"type:varchar(255);not null" json:"previousSignature" binding:"required" immutable:"true"`
	Signature           string `gorm:"type:varchar(255);not null" json:"signature" binding:"required" immutable:"true"`
	Processed           *bool  `gorm:"type:tinyint(1);not null;default:1" json:"processed" default:"false"`
}
