package po

import "time"

// 用户日志
type UserLog struct {
	Id        uint64    `gorm:"column:id" json:"id"`                 //id
	UserId    int64     `gorm:"column:user_id" json:"user_id"`       //user_id
	Type      string    `gorm:"column:type" json:"type"`             //日志类型：info, err, warning, error
	Info      string    `gorm:"column:info" json:"info"`             //日志内容
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"` //创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"` //更新时间
}
