package domain

type UserRepo interface {
	GetByID(id int64) (*User, error)
	//GetByUserName(username string) (*User, error)
	//GetByEmail(email string) (*User, error)
	//Create(user *User) (*User, error)
}

type BookRepo interface {
	//GetByID(id int64) (*Book, error)
	//GetByName(bookName string) (*Book, error)
	//GetByISBN(isbn string) (*Book, error)
	//GetByAuthorID(AuthorID int64)([]Book, error)
	//Create(book *Book) (*Book, error)
}

type AuthorRepo interface {
	//GetByID(id int64) (*Author, error)
	//GetByName(AuthorName string) (*Author, error)
	//Create(author *Author) (*Author, error)
}

type DB struct {
	UserRepo   UserRepo
	BookRepo   BookRepo
	AuthorRepo AuthorRepo
}

type Domain struct {
	DB DB
}
