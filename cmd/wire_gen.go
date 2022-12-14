// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package cmd

import (
	"github.com/viafcccy/garden-be/application/service/user"
)

// Injectors from wire.go:

//go:generate wire
func InitApp() (*App, error) {
	userService := service.NewUserService()
	app := NewApp(userService)
	return app, nil
}
