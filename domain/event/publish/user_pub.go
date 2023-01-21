package publish

import (
	"github.com/viafcccy/garden-be/global"
	"github.com/viafcccy/garden-be/infrastructure/pkg/mq/nsq"
)

// UserNsqPubDemo 发送消息案例，这里模拟的是发送一个获取用户信息的 mq
func UserNsqPubDemo(id uint64) error {
	msg := make(map[string]uint64)
	msg["id"] = id

	client := nsq.NewPublishClient()
	err := client.PublishMsg("test_topic", msg)
	if err != nil {
		global.GLog.Errorf("ERROR: UserNsqDemo producer :")
	}

	return nil
}
