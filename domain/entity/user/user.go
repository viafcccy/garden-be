package entity

import "time"

type User struct {
	// db 字段
	Id          uint64
	UserName    string
	NickName    string
	ShaPassword string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time

	// 非 db 字段
	RawPassword  string `json:"rawPassword"`
	Token        string `json:"token"`
	SuccessLogin bool   `json:"successLog"`
}
