package models

import (
	"time"
)

type CodeUsageLog struct {
	ID          uint       `gorm:"primary_key"`
	CodeID      uint       `gorm:"uniqueIndex:code_usage_logs_code_id_reference_id_client_id_uindex"`
	ReferenceID uint       `gorm:"uniqueIndex:code_usage_logs_code_id_reference_id_client_id_uindex"`
	ClientID    uint       `gorm:"uniqueIndex:code_usage_logs_code_id_reference_id_client_id_uindex"`
	UsedAt      *time.Time `gorm:"default:(-)"`
}
