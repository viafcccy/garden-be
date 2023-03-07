package repositoryimpl

import (
	"context"
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"

	entity "github.com/viafcccy/garden-be/domain/entity/user"
	userRepository "github.com/viafcccy/garden-be/domain/repository/user"
	"github.com/viafcccy/garden-be/global"
	"github.com/viafcccy/garden-be/infrastructure/consts"
	converter "github.com/viafcccy/garden-be/infrastructure/persistence/converter/user"
	po "github.com/viafcccy/garden-be/infrastructure/persistence/po/user"
)

// 检查是否实现 userRepository.UserRepository 接口
var _ userRepository.UserRepository = (*UserDao)(nil)

type UserDao struct {
	db *gorm.DB
}

func NewUserRepo() *UserDao {
	return &UserDao{
		db: global.GDB,
	}
}

func (u *UserDao) getCacheKey(data string) string {
	return fmt.Sprintf("%s%s", consts.UserCacheKey, data)
}

func (u *UserDao) SaveUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	poUser := converter.E2PUser(user)
	err := u.db.Create(&poUser).Error
	if err != nil {
		global.GLog.Errorln(err.Error())
		return nil, err
	}

	enUser := converter.P2EUser(poUser)
	return enUser, nil
}

func (u *UserDao) GetUser(ctx context.Context, id uint64) (*entity.User, error) {
	var (
		poUser po.User
	)
	err := u.db.First(&poUser, id).Error
	if err != nil {
		global.GLog.Errorln(err.Error())
		return nil, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("food not found")
	}

	enUser := converter.P2EUser(&poUser)
	return enUser, nil
}

func (u *UserDao) GetUserByUserName(userName string) (*entity.User, error) {
	var (
		poUser po.User
	)
	err := u.db.First(&poUser, po.User{UserName: userName}).Error
	if err != nil {
		global.GLog.Errorln(err.Error())
		return nil, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("food not found")
	}

	enUser := converter.P2EUser(&poUser)
	return enUser, nil
}
