package entity

import "time"

type User struct {
	Id           uint64     `json:"id"`
	UserName     string     `json:"userName"`
	NickName     string     `json:"nickName"`
	Password     string     `json:"password"`
	RawPassword  string     `json:"rawPassword"`
	Token        string     `json:"token"`
	SuccessLogin bool       `json:"successLog"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	DeletedAt    *time.Time `json:"deletedAt"`
}
