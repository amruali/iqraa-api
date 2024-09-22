package domain

import (
	"iqraa-api/models"
)

type UserRepo interface {
	GetByID(id int64) (*models.User, error)
	GetByUserName(username string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Create(user *models.User) (*models.User, error)
}

type BookRepo interface {
	GetByEra(from, to int32) ([]models.Book, error)
	GetByPublisherName(PublisherName string) ([]models.Book, error)
	GetByAuthorName(authorName string) ([]models.Book, error)
	GetByAuthorID(AuthorID int32) ([]models.Book, error)
	GetByPublisherID(PublisherID int32) ([]models.Book, error)
	GetByYear(year int32) ([]models.Book, error)
	GetByISBN(isbn string) (*models.Book, error)
	GetByName(bookName string) (*models.Book, error)
	Delete(bookID int64) error
	UpdateByID(book *models.Book) error
	GetByID(bookID int64) (*models.Book, error)
	Create(book *models.Book) (*models.Book, error)
}

type AuthorRepo interface {
	//GetByID(id int64) (*Author, error)
	GetByName(AuthorName string) (*models.Author, error)
	Create(author *models.Author) (*models.Author, error)
}

type ReviewRepo interface {
	Create(review *models.Review) (*models.Review, error)
	GetByID(reviewID int32) (*models.Review, error)
	GetByBookID(bookID int32) ([]models.Review, error)
	GetByBookName(bookName string) ([]models.Review, error)
	GetByUserID(userID int32) ([]models.Review, error)
}

type StatisticsRepo interface {
	GetByTopDownloaded(count int) ([]models.Book, error)
}

type QuoteRepo interface {
	Create(quote *models.Quote) (*models.Quote, error)
	GetByID(quoteID uint32) (*models.Quote, error)
	GetByBookID(quoteID uint32) ([]models.Quote, error)
}

// Next Redis DataInterfaces

type RedisStringsRepo interface {
	GetStrings(key string) (string, error)
	SetStrings(key, value string) error
}

type RedisDB struct {
	RedisStringsRepo RedisStringsRepo
}

type DB struct {
	UserRepo       UserRepo
	BookRepo       BookRepo
	AuthorRepo     AuthorRepo
	ReviewRepo     ReviewRepo
	StatisticsRepo StatisticsRepo
	QuoteRepo      QuoteRepo
}

type Domain struct {
	DB      DB
	RedisDB RedisDB
}
