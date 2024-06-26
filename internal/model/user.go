package model

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

type Account struct {
	Model
	CreatedAt     *time.Time            `json:"createdAt" readonly:"true" gen:"-"`
	User          string                `gorm:"type:varchar(255);not null;index:deleted_user_provider" json:"user" binding:"required" immutable:"true"`
	Provider      string                `gorm:"type:varchar(255);not null;index:deleted_user_provider" json:"provider" binding:"required" immutable:"true"`
	Balance       string                `gorm:"type:varchar(255);not null" json:"balance"`
	PendingRefund string                `gorm:"type:varchar(255);not null" json:"pendingRefund"`
	Refunds       Refunds               `gorm:"type:json;not null;default:('[]')" json:"refunds"`
	DeletedAt     soft_delete.DeletedAt `gorm:"softDelete:nano;not null;default:0;index:deleted_user_provider" json:"-" readonly:"true"`
}

type Refunds []Refund

type Refund struct {
	Amount    string `json:"amount"`
	CreatedAt string `json:"createdAt" readonly:"true" gen:"-"`
	Processed bool   `json:"processed"`
}

type AccountList struct {
	Metadata ListMeta  `json:"metadata"`
	Items    []Account `json:"items"`
}
