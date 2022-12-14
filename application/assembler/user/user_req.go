package assembler

import (
	dto "github.com/viafcccy/garden-be/application/dto/user"
	entity "github.com/viafcccy/garden-be/domain/entity/user"
)

// UserREQ 请求参数
type UserREQ struct {
}

// NewUserREQ NewUserREQ
func NewUserREQ() *UserREQ {
	return &UserREQ{}
}

// D2ESimpleUserInfo 将传入的 dto的SimpleUserInfoReq转成 entity.User
func (u *UserREQ) D2ESimpleUserInfo(dto *dto.SimpleUserInfoReq) *entity.User {
	var userEntity entity.User

	userEntity.ID = dto.Id

	return &userEntity
}
