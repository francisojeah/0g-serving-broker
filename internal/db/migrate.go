package db

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"

	"github.com/0glabs/0g-data-retrieve-agent/internal/model"
)

func Migrate(database *gorm.DB) error {
	database.Set("gorm:table_options", "ENGINE=InnoDB")

	m := gormigrate.New(database, &gormigrate.Options{UseTransaction: false}, []*gormigrate.Migration{
		{
			ID: "create-service",
			Migrate: func(tx *gorm.DB) error {
				type Service struct {
					model.Model
					CreatedAt   *time.Time            `json:"createdAt" readonly:"true" gen:"-"`
					Name        string                `gorm:"type:varchar(255);not null;uniqueIndex:deleted_name" json:"name" binding:"required" immutable:"true"`
					Type        string                `gorm:"type:varchar(255);not null" json:"type" binding:"required"`
					URL         string                `gorm:"type:varchar(255);not null" json:"url" binding:"required"`
					InputPrice  int64                 `gorm:"type:bigint;not null" json:"inputPrice" binding:"required"`
					OutputPrice int64                 `gorm:"type:bigint;not null" json:"outputPrice" binding:"required"`
					DeletedAt   soft_delete.DeletedAt `gorm:"softDelete:nano;not null;default:0;index:deleted_name" json:"-" readonly:"true"`
				}
				return tx.AutoMigrate(&Service{})
			},
		},
		{
			ID: "create-account",
			Migrate: func(tx *gorm.DB) error {
				type Account struct {
					model.Model
					CreatedAt     string                `gorm:"type:varchar(255);not null" json:"createdAt" immutable:"true"`
					User          string                `gorm:"type:varchar(255);not null;index:deleted_user_provider" json:"user" binding:"required" immutable:"true"`
					Provider      string                `gorm:"type:varchar(255);not null;index:deleted_user_provider" json:"provider" binding:"required" immutable:"true"`
					Balance       string                `gorm:"type:varchar(255);not null" json:"balance"`
					PendingRefund string                `gorm:"type:varchar(255);not null" json:"pendingRefund"`
					Refunds       model.Refunds         `gorm:"type:json;not null;default:('[]')" json:"refunds"`
					DeletedAt     soft_delete.DeletedAt `gorm:"softDelete:nano;not null;default:0;index:deleted_user_provider" json:"-" readonly:"true"`
				}
				return tx.AutoMigrate(&Account{})
			},
		},
		{
			ID: "create-request",
			Migrate: func(tx *gorm.DB) error {
				type Request struct {
					model.Model
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
				return tx.AutoMigrate(&Request{})
			},
		},
	})

	return errors.Wrap(m.Migrate(), "migrate database")
}
