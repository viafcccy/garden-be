package global

import (
	"github.com/jinzhu/gorm"

	"github.com/viafcccy/garden-be/infrastructure/config"
	"github.com/viafcccy/garden-be/infrastructure/pkg/log"
)

var (
	Gconfig *config.Config // 全局配置
	GLog    log.Logger     // 全局log
	GDB     *gorm.DB       // 全局 DB
)
