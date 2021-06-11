package service

import (
	"github.com/zewei1022/lemon-gin-web-framework/global"
	"github.com/zewei1022/lemon-gin-web-framework/model"
	"github.com/zewei1022/lemon-gin-web-framework/model/request"
)

func GetBookList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.LGWF_DB.Model(&model.Book{})
	var bookList []model.Book
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&bookList).Error
	return err, bookList, total
}
