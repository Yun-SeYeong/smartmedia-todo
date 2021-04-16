package database

import (
	"time"
	"todo/src/models"
	"todo/src/setting"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseManager struct {
	Name string
	DBMS *gorm.DB
}

var isInstance bool = false
var instance *DatabaseManager

func New(n string) (*DatabaseManager, error) {
	if !isInstance {
		instance = &DatabaseManager{
			Name: n,
		}

		db, err := gorm.Open(mysql.Open(setting.MYSQL_INFO), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		})
		if err != nil {
			return nil, err
		}
		mysqlDB, err := db.DB()
		if err != nil {
			return nil, err
		}
		mysqlDB.SetMaxIdleConns(setting.MAX_OPEN_CONNECTION)
		mysqlDB.SetMaxOpenConns(setting.MAX_IDLE_CONNECTION)
		mysqlDB.SetConnMaxLifetime(time.Hour)

		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Todo{})

		instance.DBMS = db

		isInstance = true
	}
	return instance, nil
}
