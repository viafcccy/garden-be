package event

import (
	"github.com/viafcccy/garden-be/cmd"
	nsq "github.com/viafcccy/garden-be/interfaces/event/subscribe"
)

// InitEvent 消息事件
func InitEvent(app *cmd.App) {
	nsq.UserSubDemo(app)
}
