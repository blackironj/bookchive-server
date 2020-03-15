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
