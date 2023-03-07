package assembler

import (
	userDto "github.com/viafcccy/garden-be/application/dto/user"
	entity "github.com/viafcccy/garden-be/domain/entity/user"
)

type UserRSP struct {
}

func NewUserRSP() *UserRSP {
	return &UserRSP{}
}

// E2DSimpleUserInfo 把用户实体映射到简单的实体 dto 中，返回给前端
func (o *UserRSP) E2DSimpleUserInfo(user *entity.User) *userDto.SimpleUserInfoRsp {
	var simpleUser *userDto.SimpleUserInfoRsp

	simpleUser.Id = user.Id
	simpleUser.NickName = user.NickName

	return simpleUser
}

// E2DLogin
func (u *UserRSP) E2DLogin(e *entity.User) *userDto.LoginRsp {
	login := &userDto.LoginRsp{}

	login.UserName = e.UserName
	login.NickName = e.NickName
	login.Token = e.Token

	return login
}
