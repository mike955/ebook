package dao

import (
	"ebook/ebook-privilege/conf"
	"fmt"
	"log"
	
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	
	"time"
)

var DB *gorm.DB

type CommonModel struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreateTime  string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

func SetUp()  {
	var err error
	DB, err = gorm.Open(conf.MYSQL_TYPE, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.MYSQL_USER,
		conf.MYSQL_PASSWORD,
		conf.MYSQL_HOST,
		conf.MYSQL_PORT,
		conf.MYSQL_DATABASE,
		))
	if err != nil {
		log.Fatalf("dao.init.Setup err: %v", err)
	}
	
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return conf.MYSQL_TABLE_Prefix + defaultTableName
	}
	DB.LogMode(true)
	DB.SingularTable(true)
	DB.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	DB.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	//DB.Callback().Delete().Replace("gorm:delete", deleteCallback)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer DB.Close()
}

// updateTimeStampForCreateCallback will set `CreatedOn`, `ModifiedOn` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Format("2006-01-02 15:04:05")
		if createTimeField, ok := scope.FieldByName("CreateTime"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}
		
		if modifyTimeField, ok := scope.FieldByName("UpdateTime"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

// updateTimeStampForUpdateCallback will set `ModifiedOn` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("UpdateTime", time.Now().Format("2006-01-02 15:04:05"))
	}
}


// addExtraSpaceIfExist adds a separator
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
