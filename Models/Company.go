package Models

import (
	"time"

	"github.com/google/uuid"
)

type Company struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	ApiKey    string    `gorm:"type:varchar(255);not null" json:"apiKey"`
}

func (b *Company) TableName() string {
	return "companies"
}
