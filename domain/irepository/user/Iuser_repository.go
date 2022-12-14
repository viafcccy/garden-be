package irepository

import (
	"context"

	entity "github.com/viafcccy/garden-be/domain/entity/user"
)

//go:generate mockgen --source ./Iuser_repository.go --destination ./mock/mock_user.go --package mock
type IUserRepository interface {
	SaveUser(ctx context.Context, user *entity.User) (*entity.User, error)
	GetUser(ctx context.Context, id uint64) (*entity.User, error)
}
