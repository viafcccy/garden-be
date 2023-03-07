package user

import (
	"context"

	"github.com/golang-module/dongle"
	entity "github.com/viafcccy/garden-be/domain/entity/user"
	repository "github.com/viafcccy/garden-be/domain/repository/user"
	"github.com/viafcccy/garden-be/global"
	dao "github.com/viafcccy/garden-be/infrastructure/persistence/repositoryimpl/user"
	"github.com/viafcccy/garden-be/infrastructure/pkg/errno"
)

// UserService user service interface
type UserService interface {
	FindUserByID(ctx context.Context, id uint64) (*entity.User, error)
	Login(userName string, rawPwd string) (userE *entity.User, err error)
}

// 检查接口实现的正确性
// // 使用声明方式
// var _ User = &Student{}
// 类型转换方式
var _ UserService = (*userService)(nil)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService() *userService {
	return &userService{
		userRepo: dao.NewUserRepo(),
	}
}

func (u *userService) FindUserByID(ctx context.Context, id uint64) (*entity.User, error) {
	return u.userRepo.GetUser(ctx, id)
}

// 用户登录
func (u *userService) Login(userName string, rawPwd string) (userE *entity.User, err error) {
	userE, err = u.userRepo.GetUserByUserName(userName)
	if err != nil {
		return nil, err
	}

	// 校验密码 加盐规则
	sha3Pwd := dongle.Encrypt.FromString(global.Gconfig.UserPassword.Salt + rawPwd).BySha3(512).ToHexString()
	if sha3Pwd == userE.Password {
		// 登录态 用户拼装
		userE.SuccessLogin = true
		userE.Token = "testToken"

		return userE, nil
	}

	// 密码不一致 登录失败
	userE.SuccessLogin = false
	return userE, errno.ErrLogFailWrongPwd
}
