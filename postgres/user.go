package postgres

import (
	"errors"
	"fmt"
	"iqraa-api/domain"
	"time"

	"github.com/go-pg/pg/v9"
)

type UserRepo struct {
	DB *pg.DB
}

func NewUserRepo(DB *pg.DB) *UserRepo {
	return &UserRepo{DB: DB}
}

func (u *UserRepo) GetByEmail(email string) (*domain.User, error) {
	var user = new(domain.User)

	fmt.Println(11)
	// Check that email is taken
	fmt.Println(email)
	err := u.DB.Model(user).Where("email = ?", email).First()
	if err != nil {
		fmt.Println(12)
		if errors.Is(err, pg.ErrNoRows) {
			fmt.Println(13)
			return nil, domain.ErrNoResult
		}
		fmt.Println(err)
		return nil, err
	}
	return user, nil
}

func (u *UserRepo) GetByUserName(username string) (*domain.User, error) {
	var user = &domain.User{}

	// Check if username is taken
	err := u.DB.Model(user).Where("username = ?", username).First()
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, domain.ErrNoResult
		}

		return nil, err
	}
	return user, nil
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

func (u *UserRepo) Create(email, username string, password []byte) (*domain.User, error) {
	user := &domain.User{
		Email:          email,
		Username:       username,
		HashedPassword: password,
		UserTypeID:     3,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	_, err := u.DB.Model(user).Returning("*").Insert()
	if err != nil {
		return nil, err
	}
	return user, nil
}
