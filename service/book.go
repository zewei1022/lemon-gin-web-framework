package service

import (
	"github.com/zewei1022/lemon-gin-web-framework/global"
	"github.com/zewei1022/lemon-gin-web-framework/model"
	"github.com/zewei1022/lemon-gin-web-framework/model/request"
)

func FindBookById(idInfo request.IdInfo) (book model.Book, err error) {
	var findBook model.Book
	global.LGWF_DB.Model(&findBook).Where("id = ?", idInfo.ID).First(&findBook)
	if book.ID > 0 {
		return findBook, nil
	}
	return findBook, err
}

func GetBookList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.LGWF_DB.Model(&model.Book{})
	var bookList []model.Book
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&bookList).Error
	return bookList, total, err
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

func UpdateBook(bookInfo request.UpdateBookInfo) (err error) {
	book := model.Book{
		Name:      bookInfo.Name,
		Author:    bookInfo.Author,
		Price:     bookInfo.Price,
		Page:      bookInfo.Page,
		Isbn:      bookInfo.Isbn,
		Publisher: bookInfo.Publisher,
	}
	err = global.LGWF_DB.Model(&model.Book{}).Where("id = ?", bookInfo.ID).Updates(book).Error
	return err
}

func DeleteBook(id int) (err error) {
	err = global.LGWF_DB.Where("id = ?", id).Delete(&model.Book{}).Error
	return err
}
