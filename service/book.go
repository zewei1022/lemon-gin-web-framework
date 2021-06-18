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

func CreateBook(bookInfo request.CreateBookInfo) (err error) {
	book := model.Book{
		Name:      bookInfo.Name,
		Author:    bookInfo.Author,
		Price:     bookInfo.Price,
		Page:      bookInfo.Page,
		Isbn:      bookInfo.Isbn,
		Publisher: bookInfo.Publisher,
	}
	err = global.LGWF_DB.Create(&book).Error
	return err
}
