package user

import (
	service "github.com/viafcccy/garden-be/application/service/user"
)

type UserServer struct {
	UserSrv *service.UserService // 注入 application/service
}
