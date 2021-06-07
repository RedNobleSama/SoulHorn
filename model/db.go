/**
* @Author: oreki
* @Date: 2021/6/6 22:03
* @Email: a912550157@gmail.com
 */

package model

import (
	"SoulHorn/utils"
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var db *gorm.DB
var err error

// InitDb 加载数据库
func InitDb() {
	db, err = gorm.Open(utils.Db, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassword,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	))
	if err != nil {
		fmt.Println("数据库连接失败,检查参数", err)
	}

	//禁用表名的复数形式，如果设置为true，则表名为user
	db.SingularTable(true)

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB := db.DB()

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	// 可以通过Set设置附加参数，下面设置表的存储引擎为InnoDB
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{}, &DouBanBook{})
}
