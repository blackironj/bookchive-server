package service

import (
	"github.com/blackironj/bookchive-server/da"
	"github.com/blackironj/bookchive-server/model"
)

func GetBook(bookID string) (*model.Books, error) {
	book, err := da.GetBookByID(da.DB, bookID)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func GetBooksInLibrary(userUUID string) ([]*model.BookInLibrary, error) {
	books, err := da.GetBooksInLibrary(da.DB, userUUID)
	if err != nil {
		return nil, err
	}
	return books, nil
}
