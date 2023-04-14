package database

import (
	"prak/config"
	"prak/models"
)

// Get all book
func GetBook() (books []models.Book, err error) {
	err = config.DB.Find(&books).Error

	if err != nil {
		return []models.Book{}, err
	}

	return
}

// Create book
func CreateBook(books models.Book) (models.Book, error) {
	err := config.DB.Create(&books).Error

	if err != nil {
		return models.Book{}, err
	}

	return books, nil
}

// Get book by id
func GetBookById(id any) (models.Book, error) {
	var books models.Book

	err := config.DB.Where("id = ?", id).First(&books).Error

	if err != nil {
		return models.Book{}, err
	}

	return books, nil
}

// Update book by id
func UpdateBookById(books models.Book, id any) (models.Book, error) {
	err := config.DB.Where("id = ?", id).Updates(&books).Error

	if err != nil {
		return models.Book{}, err
	}

	return books, nil
}

// Delete book by id
func DeleteBook(id any) (interface{}, error) {
	err := config.DB.Where("id = ?", id).Delete(&models.Book{}).Error

	if err != nil {
		return nil, err
	}

	return "success delete book", nil
}
