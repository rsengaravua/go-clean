package models

import "time"

type Video struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Url         string    `json:"url"`
	Description string    `json:"thumbnail"`
	UpdatedAt   time.Time `db:"updated_at json:"updated_at`
	CreatedAt   time.Time `db:"created_at json:"created_at`
}
