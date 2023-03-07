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

// D2ESimpleUserInfo 将传入的 dto 的 SimpleUserInfoReq 转成 entity.User
func (o *UserREQ) D2ESimpleUserInfo(dto *dto.SimpleUserInfoReq) *entity.User {
	var userEntity entity.User

	userEntity.Id = dto.Id

	return &userEntity
}

func (o *UserREQ) D2ELogin(dto *dto.LoginReq) *entity.User {
	var userEntity entity.User

	userEntity.UserName = dto.UserName
	userEntity.RawPassword = dto.Password

	return &userEntity
}
