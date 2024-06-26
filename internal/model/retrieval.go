package model

import "time"

type Retrieval struct {
	Model
	CreatedAt *time.Time `json:"createdAt" readonly:"true" gen:"-"`
	Nonce     string     `gorm:"type:varchar(255);not null" json:"nonce" binding:"required" immutable:"true"`
	Provider  string     `gorm:"type:varchar(255);not null" json:"provider" binding:"required" immutable:"true"`
	Refunds   Refunds    `gorm:"type:json;not null;default:('[]')" json:"refunds"`
}

type RetrievalList struct {
	Metadata ListMeta    `json:"metadata"`
	Items    []Retrieval `json:"items"`
}
