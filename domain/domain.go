package domain

type UserRepo interface {
	GetByID(id int64) (*User, error)
	GetByUserName(username string) (*User, error)
	GetByEmail(email string) (*User, error)
	Create(email, username string, Password []byte) (*User, error)
}

type BookRepo interface {
	GetByEra(from, to int32) ([]Book, error)
	GetByPublisherName(PublisherName string) ([]Book, error)
	GetByAuthorName(authorName string) ([]Book, error)
	GetByAuthorID(AuthorID int32) ([]Book, error)
	GetByPublisherID(PublisherID int32) ([]Book, error)
	GetByYear(year int32) ([]Book, error)
	GetByISBN(isbn string) (*Book, error)
	GetByName(bookName string) (*Book, error)
	Create(book *Book) (*Book, error)
}

type AuthorRepo interface {
	//GetByID(id int64) (*Author, error)
	GetByName(AuthorName string) (*Author, error)
	Create(author *Author) (*Author, error)
}

type ReviewRepo interface {
	Create(review *Review) (*Review, error)
	GetByID(reviewID int32) (*Review, error)
	GetByBookID(bookID int32) ([]Review, error)
	GetByBookName(bookName string) ([]Review, error)
	GetByUserID(userID int32) ([]Review, error)
}

type StatisticsRepo interface {
	GetByMostDownloadable(count int) ([]Book, error)
}

type QuoteRepo interface {
	Create(quote *Quote) (*Quote, error)
	GetByID(quoteID uint32) (*Quote, error)
	GetByBookID(quoteID uint32) ([]Quote, error)
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
	DB DB
}
