// Code generated by gen; DO NOT EDIT.

package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ================================= Account =================================
func (d *Account) BeforeCreate(tx *gorm.DB) error {
	if d.ID == nil {
		d.ID = PtrOf(uuid.New())
	}
	return nil
}

func (d *Account) Bind(ctx *gin.Context) error {
	var r Account
	if err := ctx.ShouldBindJSON(&r); err != nil {
		return err
	}
	d.User = r.User
	d.Provider = r.Provider
	d.Balance = r.Balance
	d.PendingRefund = r.PendingRefund
	d.Refunds = r.Refunds

	return nil
}

func (d *Account) BindWithReadonly(ctx *gin.Context, old Account) error {
	if err := d.Bind(ctx); err != nil {
		return err
	}
	if d.ID == nil {
		d.ID = old.ID
	}

	return nil
}

// ================================= Request =================================
func (d *Request) BeforeCreate(tx *gorm.DB) error {
	if d.ID == nil {
		d.ID = PtrOf(uuid.New())
	}
	return nil
}

func (d *Request) Bind(ctx *gin.Context) error {
	var r Request
	if err := ctx.ShouldBindJSON(&r); err != nil {
		return err
	}
	d.CreatedAt = r.CreatedAt
	d.UserAddress = r.UserAddress
	d.Nonce = r.Nonce
	d.Name = r.Name
	d.InputCount = r.InputCount
	d.PreviousOutputCount = r.PreviousOutputCount
	d.PreviousSignature = r.PreviousSignature
	d.Signature = r.Signature

	return nil
}

func (d *Request) BindWithReadonly(ctx *gin.Context, old Request) error {
	if err := d.Bind(ctx); err != nil {
		return err
	}
	if d.ID == nil {
		d.ID = old.ID
	}

	return nil
}

// ================================= Retrieval =================================
func (d *Retrieval) BeforeCreate(tx *gorm.DB) error {
	if d.ID == nil {
		d.ID = PtrOf(uuid.New())
	}
	return nil
}

func (d *Retrieval) Bind(ctx *gin.Context) error {
	var r Retrieval
	if err := ctx.ShouldBindJSON(&r); err != nil {
		return err
	}
	d.Nonce = r.Nonce
	d.Provider = r.Provider
	d.Refunds = r.Refunds

	return nil
}

func (d *Retrieval) BindWithReadonly(ctx *gin.Context, old Retrieval) error {
	if err := d.Bind(ctx); err != nil {
		return err
	}
	if d.ID == nil {
		d.ID = old.ID
	}

	return nil
}

// ================================= Service =================================
func (d *Service) BeforeCreate(tx *gorm.DB) error {
	if d.ID == nil {
		d.ID = PtrOf(uuid.New())
	}
	return nil
}

func (d *Service) Bind(ctx *gin.Context) error {
	var r Service
	if err := ctx.ShouldBindJSON(&r); err != nil {
		return err
	}
	d.Name = r.Name
	d.Type = r.Type
	d.URL = r.URL
	d.InputPrice = r.InputPrice
	d.OutputPrice = r.OutputPrice

	return nil
}

func (d *Service) BindWithReadonly(ctx *gin.Context, old Service) error {
	if err := d.Bind(ctx); err != nil {
		return err
	}
	if d.ID == nil {
		d.ID = old.ID
	}

	return nil
}

//=============== implementation of sql.scanner and sql.valuer  ===============
func (m Refunds) Value() (driver.Value, error) {
	return json.Marshal(m)
}
func (m *Refunds) Scan(value interface{}) error {
	if v, ok := value.([]byte); ok {
		return json.Unmarshal(v, m)
	}
	return fmt.Errorf("can't convert %T to Refunds", value)
}	
