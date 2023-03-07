package service

import (
	"context"

	assembler "github.com/viafcccy/garden-be/application/assembler/user"
	dto "github.com/viafcccy/garden-be/application/dto/user"
	"github.com/viafcccy/garden-be/domain/behavior/user"
)

type UserService struct {
	assUserRSP    *assembler.UserRSP
	assUserREQ    *assembler.UserREQ
	domainService user.UserService
}

func NewUserService() *UserService {
	return &UserService{
		assUserRSP:    assembler.NewUserRSP(),
		assUserREQ:    assembler.NewUserREQ(),
		domainService: user.NewUserService(),
	}
}

// GetSimpleUserInfo 获取用户信息给到 interfaces
func (u *UserService) GetSimpleUserInfo(ctx context.Context, req *dto.SimpleUserInfoReq) (*dto.SimpleUserInfoRsp, error) {

	userEntity := u.assUserREQ.D2ESimpleUserInfo(req)

	entUser, err := u.domainService.FindUserByID(ctx, userEntity.Id) // 业务复杂的话，这里应该调用 domain/aggregate聚合
	if err != nil {
		return nil, err
	}
	return u.assUserRSP.E2DSimpleUserInfo(entUser), nil
}

func (u *UserService) Login(ctx context.Context, req *dto.LoginReq) (*dto.LoginRsp, error) {

	userEntity := u.assUserREQ.D2ELogin(req)

	entUser, err := u.domainService.Login(userEntity.UserName, userEntity.RawPassword) // 业务复杂的话，这里应该调用 domain/aggregate聚合
	if err != nil {
		return nil, err
	}
	return u.assUserRSP.E2DLogin(entUser), nil
}
