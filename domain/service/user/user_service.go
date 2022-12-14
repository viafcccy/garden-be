package user

import (
	"context"

	entity "github.com/viafcccy/garden-be/domain/entity/user"
	repository "github.com/viafcccy/garden-be/domain/irepository/user"
	dao "github.com/viafcccy/garden-be/infrastructure/repository/repositoryimpl/user"
)

// UserService user service interface
type UserService interface {
	FindUserByID(ctx context.Context, id uint64) (*entity.User, error)
}

var _ UserService = (*userService)(nil)

type userService struct {
	userRepo repository.IUserRepository
}

func NewUserService() *userService {
	return &userService{
		userRepo: dao.NewUserRepo(),
	}
}

func (u *userService) FindUserByID(ctx context.Context, id uint64) (*entity.User, error) {
	return u.userRepo.GetUser(ctx, id)
}
