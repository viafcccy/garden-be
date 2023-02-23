package service

import (
	"context"
	"fmt"

	dto "github.com/viafcccy/garden-be/application/dto/user"
)

// GetSimpleUserInfo 获取用户信息给到 interfaces
func (u *UserService) Login(ctx context.Context, req *dto.SimpleUserInfoReq) (*dto.SimpleUserInfo, error) {

	// 测试专用接口，保护内部安全
	if req.Id > 1 {
		return nil, fmt.Errorf("测试接口，非法参数")
	}

	userEntity := u.assUserREQ.D2ESimpleUserInfo(req)

	entUser, err := u.domainService.FindUserByID(ctx, userEntity.Id) // 业务复杂的话，这里应该调用 domain/aggregate聚合
	if err != nil {
		return nil, err
	}
	return u.assUserRSP.E2DSimpleUserInfo(entUser), nil
}
