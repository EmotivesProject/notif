package model

import "time"

type Notification struct {
	ID        int       `json:"id,omitempty" gorm:"primaryKey"`
	Username  string    `json:"username"`
	Title     string    `json:"title"`
	Message   string    `json:"message"`
	Link      string    `json:"link"`
	CreatedAt time.Time `json:"created_time" gorm:"autoCreateTime"`
	Seen      bool      `json:"seen"`
}
