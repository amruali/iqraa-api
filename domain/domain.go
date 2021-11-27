package domain

type UserRepo interface {
	GetByID(id int64) (*User, error)
	GetByUserName(username string) (*User, error)
	GetByEmail(email string) (*User, error)
	Create(email, username string, Password []byte) (*User, error)
}

type BookRepo interface {
	GetByMostDownloadable(count int) ([]Book, error)
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
}
type DB struct {
	UserRepo   UserRepo
	BookRepo   BookRepo
	AuthorRepo AuthorRepo
	ReviewRepo ReviewRepo
}

type Domain struct {
	DB DB
}
