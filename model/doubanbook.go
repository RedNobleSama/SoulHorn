/**
* @Author: oreki
* @Date: 2021/6/7 22:02
* @Email: a912550157@gmail.com
 */

package model

import (
	"SoulHorn/utils/errmsg"
	"SoulHorn/utils/gredis"
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
)

// DouBanBook 豆瓣图书模型
type DouBanBook struct {
	//GORM 定义一个 gorm.Model 结构体，其包括字段 ID、CreatedAt、UpdatedAt、DeletedAt
	ID        uint            `gorm:"primary_key AUTO_INCREMENT;"`
	Title     string          `gorm:"type:varchar(50);" json:"title" label:"书名"`
	Author    string          `gorm:"type:varchar(50);"  json:"author" label:"作者"`
	Press     string          `gorm:"type:varchar(100);"  json:"press" label:"出版社"`
	PressTime string          `gorm:"type:varchar(50);"  json:"presstime" label:"出版时间"`
	ISBM      string          `gorm:"type:varchar(20);"  json:"isbm" label:"ISBM"`
	Score     decimal.Decimal `gorm:"type:decimal(20,8);"  json:"score" label:"评分"`
}

//func (b DouBanBook)GetBook(book *DouBanBook, pageSize int, pageNum int) []DouBanBook{
//	var res []DouBanBook
//	if book.Title != "" {
//		db = db.Where("title = ?", book.Title)
//	}
//	if book.Author != "" {
//		db = db.Where("title = ?", book.Title)
//	}
//	db.Find(&res)
//	return res
//}

func (b DouBanBook) GetBooks(pageSize int, pageNum int) ([]DouBanBook, int) {
	var res []DouBanBook
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&res).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}
	gredis.Cache(res)
	return res, errmsg.SUCCESS
}

//func (b DouBanBook) Cache(data []DouBanBook) {
//	var ctx = context.Background()
//	redis := gredis.InitRedis()
//	marshal, err := json.Marshal(data)
//	if err != nil {
//		return
//	}
//	_, err = redis.Set(ctx, "moose-go", marshal, 10*time.Minute).Result()
//	if err != nil {
//		fmt.Println("失败")
//	}
//	fmt.Println("成功")
//}
