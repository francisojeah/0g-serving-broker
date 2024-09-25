package db

import (
	"time"

	"github.com/0glabs/0g-serving-broker/provider/model"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

func (d *DB) Migrate() error {
	d.db.Set("gorm:table_options", "ENGINE=InnoDB")

	m := gormigrate.New(d.db, &gormigrate.Options{UseTransaction: false}, []*gormigrate.Migration{
		{
			ID: "create-service",
			Migrate: func(tx *gorm.DB) error {
				type Service struct {
					model.Model
					Name        string                `gorm:"type:varchar(255);not null;uniqueIndex:deleted_name"`
					Type        string                `gorm:"type:varchar(255);not null"`
					URL         string                `gorm:"type:varchar(255);not null"`
					InputPrice  int64                 `gorm:"type:bigint;not null"`
					OutputPrice int64                 `gorm:"type:bigint;not null"`
					DeletedAt   soft_delete.DeletedAt `gorm:"softDelete:nano;not null;default:0;index:deleted_name"`
				}
				return tx.AutoMigrate(&Service{})
			},
		},
		{
			ID: "create-user",
			Migrate: func(tx *gorm.DB) error {
				type User struct {
					model.Model
					User                 string                `gorm:"type:varchar(255);not null;uniqueIndex:deleted_user"`
					LastRequestNonce     *int64                `gorm:"type:bigint;not null;default:0"`
					LastResponseFee      *int64                `gorm:"type:bigint;not null;default:0"`
					LockBalance          *int64                `gorm:"type:bigint;not null;default:0"`
					LastBalanceCheckTime *time.Time            `json:"lastBalanceCheckTime"`
					Signer               model.StringSlice     `gorm:"type:json;not null;default:('[]')"`
					UnsettledFee         *int64                `gorm:"type:bigint;not null;default:0"`
					DeletedAt            soft_delete.DeletedAt `gorm:"softDelete:nano;not null;default:0;index:deleted_user"`
				}
				return tx.AutoMigrate(&User{})
			},
		},
		{
			ID: "create-request",
			Migrate: func(tx *gorm.DB) error {
				type Request struct {
					model.Model
					UserAddress       string `gorm:"type:varchar(255);not null;uniqueIndex:processed_userAddress_nonce"`
					Nonce             int64  `gorm:"type:bigint;not null;index:processed_userAddress_nonce"`
					ServiceName       string `gorm:"type:varchar(255);not null"`
					InputFee          int64  `gorm:"type:bigint;not null"`
					PreviousOutputFee int64  `gorm:"type:bigint;not null"`
					Fee               int64  `gorm:"type:bigint;not null"`
					Signature         string `gorm:"type:varchar(255);not null"`
					Processed         *bool  `gorm:"type:tinyint(1);not null;default:0;index:processed_userAddress_nonce"`
				}
				return tx.AutoMigrate(&Request{})
			},
		},
	})

	return errors.Wrap(m.Migrate(), "migrate database")
}
