package Models

import (
	"time"

	"github.com/google/uuid"
)

type Provider struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey;" json:"id" validate:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name" validate:"required"`
	Email     string    `gorm:"type:varchar(255);not null" json:"email" validate:"required"`
	Phone     string    `gorm:"type:varchar(255);not null" json:"phone" validate:"required"`
	Address   string    `gorm:"type:varchar(255);not null" json:"address" validate:"required"`
	CompanyId uuid.UUID `gorm:"type:char(36);index" json:"companyId" validate:"required"`
	IsActive  bool      `gorm:"default:true" json:"isActive"`
}

func (b *Provider) TableName() string {
	return "providers"
}
