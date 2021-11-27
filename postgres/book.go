package postgres

import (
	"errors"

	"iqraa-api/domain"

	"github.com/go-pg/pg/v9"
)

type BookRepo struct {
	DB *pg.DB
}

func NewBookRepo(DB *pg.DB) *BookRepo {
	return &BookRepo{DB: DB}
}

func (b *BookRepo) Create( /*name, isbn string, year int32*/ book *domain.Book) (*domain.Book, error) {
	/*
		book := &domain.Book{
			BookName:         name,
			ISBN:             isbn,
			PublishYear:      year,
			PublisherID:      1,
			BookTypeID:       1,
			BookTypeDetailID: 1,
			BookAuthorID:     1,
			CreateUserID:     1,
			UpdateUserID:     1,
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
		}
	*/
	_, err := b.DB.Model(book).Returning("*").Insert()
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (b *BookRepo) GetByName(BookName string) (*domain.Book, error) {
	book := &domain.Book{}
	err := b.DB.Model(book).Where("book_name = ?", BookName).First()
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, domain.ErrNoResult
		}
		return nil, err
	}

	return book, nil
}

func (b *BookRepo) GetByISBN(isbn string) (*domain.Book, error) {
	book := &domain.Book{}
	err := b.DB.Model(book).Where("isbn = ?", isbn).First()
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, domain.ErrNoResult
		}
		return nil, err
	}
	return book, nil
}

// Get Books By Author Name
func (b *BookRepo) GetByAuthorName(authorName string) ([]domain.Book, error) {
	books := []domain.Book{}
	err := b.DB.Model(&books).
		ColumnExpr("book.*").
		//ColumnExpr("a.id AS author__id, a.full_name AS author__name").
		Join("JOIN authors a"). 
		JoinOn("a.id = book.book_author_id").
		JoinOn("a.full_name = ?", authorName).
		Select()
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, domain.ErrNoResult
		}
		return nil, err
	}
	return books, nil
}

// Get Books By Author ID
func (b *BookRepo) GetByAuthorID(AuthorID int32) ([]domain.Book, error) {
	books := []domain.Book{}
	err := b.DB.Model(&books).Where("book_author_id = (?)", AuthorID).Select()
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, domain.ErrNoResult
		}
		return nil, err
	}
	return books, nil
}

// Get Books By Publisher ID
func (b *BookRepo) GetByPublisherID(PublisherID int32) ([]domain.Book, error) {
	books := []domain.Book{}
	err := b.DB.Model(&books).Where("publisher_id IN (?)", PublisherID).Select()
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, domain.ErrNoResult
		}
		return nil, err
	}
	return books, nil
}

// Get Books Published In Specific Year
func (b *BookRepo) GetByYear(year int32) ([]domain.Book, error) {
	books := []domain.Book{}
	err := b.DB.Model(&books).Where("publish_year = (?)", year).Select()
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, domain.ErrNoResult
		}
		return nil, err
	}
	return books, nil
}


// Get Books By Publisher Name
func (b *BookRepo) GetByPublisherName(PublisherName string) ([]domain.Book, error) {
	books := []domain.Book{}
	err := b.DB.Model(&books).
		ColumnExpr("book.*").
		//ColumnExpr("a.id AS author__id, a.full_name AS author__name").
		Join("JOIN publisher p"). 
		JoinOn("p.publisher_id = book.publisher_id").
		JoinOn("p.publishing_house_name = ?", PublisherName).
		Select()
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, domain.ErrNoResult
		}
		return nil, err
	}
	return books, nil
}



// Get Books Between two years
func (b *BookRepo) GetByEra(from, to int32) ([]domain.Book, error) {
	books := []domain.Book{}
	err := b.DB.Model(&books).
		Where("publish_year >= ?", from).
		Where("publish_year <= ?", to).
		Select()
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, domain.ErrNoResult
		}
		return nil, err
	}
	return books, nil
}

/*

// Get Books By Type ID
// Get Books By Type Detail ID
func (b *BookRepo) GetByTypeAndTypeDetail(typeID, typeDetailID int32) ([]*domain.Book, error) {
	query := ""
	if typeID != -1 {
		query += ""
	}
	if typeDetailID != -1 {
		query += ""
	}

}
*/

/*
//book := &domain.Book{}
	books := []domain.Book{}

	//YearBooks := b.DB.Model(book).ColumnExpr("publish_year").Where("publish_year = ?", 1)
	//err := b.DB.Model(&books).Where("publish_year IN (?)", YearBooks).Select()

*/

//https://pg.uptrace.dev/queries/
