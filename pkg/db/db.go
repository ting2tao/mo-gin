package db

import (
	"errors"
	"fmt"
	"github.com/slovty/mo-gin/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBSet(cfg *config.Config) (*gorm.DB, error) {
	dsn := cfg.GetString("MYSQL_URI")
	if dsn == "" {
		fmt.Println("empty dsn")
		return nil, errors.New("empty dsn")
	}
	return NewDB(dsn)
}

func NewDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		//Logger:                                   logger.New(),
	})
	if err != nil {
		_ = fmt.Errorf(err.Error())
		return nil, err
	}

	return db.Debug(), nil
}
