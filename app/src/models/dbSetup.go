package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"monitor/extend/conf"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Setup() {
	var err error

	var connectStr = fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.DBConf.User,
		conf.DBConf.Password,
		conf.DBConf.Host,
		conf.DBConf.DBName,
	)
	DB, err = gorm.Open(conf.DBConf.DBType, connectStr)

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
		time.Sleep(5 * time.Second)
		DB, err = gorm.Open(conf.DBConf.DBType, connectStr)
		if err != nil {
			panic(err.Error())
		}
	}
	if DB.Error != nil {
		fmt.Printf("database error %v", DB.Error)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return conf.DBConf.TablePrefix + defaultTableName
	}

	DB.LogMode(conf.DBConf.Debug)
	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)

	//migrate 迁移
	DB.Set(
		"gorm:table_options",
		"ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci",
	).AutoMigrate(&User{}, &Task{})
	DB.Model(&User{}).AddUniqueIndex("uk_email", "email")

}
