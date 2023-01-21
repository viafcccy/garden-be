package nsq

import (
	"encoding/json"

	"github.com/viafcccy/garden-be/global"
)

type Publish interface {
	PublishMsg(topic string, msg interface{}) (err error)
}

var _ Publish = (*publish)(nil)

type publish struct {
}

func NewPublishClient() Publish {
	return &publish{}
}

// PublishMsg 推送 nsq 消息
func (p *publish) PublishMsg(topic string, msg interface{}) (err error) {
	var reqBodyByte []byte
	reqBodyByte, err = json.Marshal(msg)
	if err != nil {
		global.GLog.Errorf("ERROR: nsqProducer json Marshal:%v\n", err)
	}

	// 发送 nsq
	client := NewNsqClient()
	cp, err := client.NsqProducer()
	if err != nil {
		global.GLog.Errorf("ERROR: NsqProducer:%v\n", err)
	}
	err = cp.Publish(topic, reqBodyByte)
	if err != nil {
		global.GLog.Errorf("ERROR: nsqProducer Publish:%v\n", err)
		return
	}

	return
}
