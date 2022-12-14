package cmd

import service "github.com/viafcccy/garden-be/application/service/user"

type App struct {
	UserSrv *service.UserService
}

func NewApp(userSrv *service.UserService) *App {
	return &App{UserSrv: userSrv}
}
