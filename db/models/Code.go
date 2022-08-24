package models

import (
	"time"
)

type Code struct {
	ID               uint   `gorm:"primary_key"`
	Text             string `gorm:"unique"`
	MaxUsableCount   uint
	CurrentUsedCount uint
	Amount           uint
	CreatedAt        time.Time
	CodeUsageLogs    []CodeUsageLog
}
