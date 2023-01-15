package boot

import (
	"github.com/viafcccy/garden-be/global"
	"github.com/viafcccy/garden-be/infrastructure/config"
	"github.com/viafcccy/garden-be/infrastructure/pkg/database/mysql"
	"github.com/viafcccy/garden-be/infrastructure/pkg/log"
)

// 初始化
func init() {
	global.Gconfig = config.NewConfig() // 初始化全局配置
	global.GLog = log.NewLogger()       // 初始化日志
	global.GDB = mysql.NewDB()          // 初始化 mysql
}
