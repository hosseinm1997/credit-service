package models

type MicroserviceClient struct {
	ID     uint `gorm:"primary_key"`
	Title  string
	Secret string
}
