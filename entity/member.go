package entity

import (
	"time"
)

type Member struct {
	ID            int64     `json:"id" gorm:"column:id;primary_key"`
	Name          string    `json:"name"`
	Birthday      string    `json:"birthday"`
	Blood         string    `json:"blood"`
	Constellation string    `json:"constellation"`
	Height        string    `json:"height"`
	Status        string    `json:"status"`
	Description   string    `json:"description"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at`
}

type Members []Member