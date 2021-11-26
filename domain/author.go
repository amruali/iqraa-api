package domain

import "time"

type Author struct {
	Id           int32     `json:"id"`
	FullName     string    `json:"full_name"`
	Dob          time.Time `json:"dob"`
	Nationality  string    `json:"nationality"`
	DateOfDeath  string    `json:"date_of_birth"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	CreateUserID int32     `json:"create_user_id"`
	UpdateUserID int32     `json:"update_user_id"`
	Image        string    `json:"image"`
	Books        []*Book   `pg:"rel:has-many"`
}
