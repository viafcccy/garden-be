package converter

import (
	entity "github.com/viafcccy/garden-be/domain/entity/user"
	po "github.com/viafcccy/garden-be/infrastructure/persistence/po/user"
)

type UserConv struct {
}

// Entity 2 PO
func E2PUser(user *entity.User) *po.User {
	var poUser po.User

	poUser.Id = user.Id
	poUser.UserName = user.UserName
	poUser.Password = user.Password
	poUser.NickName = user.NickName
	poUser.CreatedAt = user.CreatedAt

	return &poUser
}

// PO 2 Entity
func P2EUser(user *po.User) *entity.User {
	var enUser entity.User

	enUser.Id = user.Id
	enUser.UserName = user.UserName
	enUser.Password = user.Password
	enUser.NickName = user.NickName
	enUser.CreatedAt = user.CreatedAt
	enUser.UpdatedAt = user.UpdatedAt
	enUser.DeletedAt = user.DeletedAt

	return &enUser
}
