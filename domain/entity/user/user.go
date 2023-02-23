package entity

import "time"

type User struct {
	Id        uint64     `json:"id"`
	UserName  string     `json:"userName"`
	Password  string     `json:"Password"`
	NickName  string     `json:"nickName"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
