package models

import (
	"github.com/huyrun/go-admin/modules/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(c db.Connection) (*gorm.DB, error) {
	return gorm.Open(postgres.New(postgres.Config{Conn: c.GetDB("default")}), &gorm.Config{})
}
