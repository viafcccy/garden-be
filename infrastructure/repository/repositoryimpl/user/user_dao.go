package repositoryimpl

import (
	"context"
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"

	entity "github.com/viafcccy/garden-be/domain/entity/user"
	userRepository "github.com/viafcccy/garden-be/domain/irepository/user"
	"github.com/viafcccy/garden-be/global"
	"github.com/viafcccy/garden-be/infrastructure/consts"
	converter "github.com/viafcccy/garden-be/infrastructure/repository/converter/user"
	po "github.com/viafcccy/garden-be/infrastructure/repository/po/user"
)

var _ userRepository.IUserRepository = (*UserDao)(nil)

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
