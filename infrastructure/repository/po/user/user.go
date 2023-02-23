package po

import "time"

// 用户基本信息
type User struct {
	Id        uint64     `gorm:"column:id" json:"id"`                 //id
	UserName  string     `gorm:"column:user_name" json:"user_name"`   //用户名，全局唯一
	Password  string     `gorm:"column:password" json:"password"`     //密码 SHA3-512 加密
	NickName  string     `gorm:"column:nick_name" json:"nick_name"`   //昵称
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"` //创建时间
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"` //更新时间
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"` //删除时间
}
