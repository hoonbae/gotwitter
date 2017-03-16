package models

import (
	"time"
)

type Tweet struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`

	User   User
	UserId uint

	Likes int `json:"likes"`

	TweetMessage string `json:"tweet_message"`
}
