package postgres

import (
	"errors"
	"fmt"
	"iqraa-api/domain"

	"github.com/go-pg/pg/v9"
)

type UserRepo struct {
	DB *pg.DB
}

func NewUserRepo(DB *pg.DB) *UserRepo {
	return &UserRepo{DB: DB}
}

func (u *UserRepo) GetByID(id int64) (*domain.User, error) {
	user := &domain.User{}
	err := u.DB.Model(user).Where("id = ?", id).First()
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, domain.ErrNoResult
		}

		return nil, err
	}

	fmt.Println(user)

	return user, nil
}

/*
func (u *UserRepo) GetByEmail(id int64) (*domain.User, error) {

}

func (u *UserRepo) GetByUserName(id int64) (*domain.User, error) {

}

func (u *UserRepo) Create(user *domain.User) (*domain.User, error) {

}
*/
