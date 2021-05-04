package datasource

import (
	"dora/config"
	"dora/pkg/logger"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var onceGorm sync.Once
var db *gorm.DB

func GormInstance() *gorm.DB {
	onceGorm.Do(func() {
		conf := config.GetConf()
		db = newGorm(conf.Gorm)
	})
	return db
}

func newGorm(config config.GormConfig) *gorm.DB {
	var dialect gorm.Dialector

	switch config.Driver {
	//case "sqlite":
	//	dialect = sqlite.Open(config.DSN)
	case "mysql":
		dialect = mysql.Open(config.DSN)
	default:
		panic("not found database config")
	}

	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func Migrate(db *gorm.DB, models []interface{}) error {
	err := db.AutoMigrate(models...)
	return err
}

func StopDataBase() {
	logger.Println("stop gorm database")
	s, err := GormInstance().DB()
	if err != nil {
		logger.Error(err)
		return
	}
	err = s.Close()
	if err != nil {
		logger.Error(err)
		return
	}
}
