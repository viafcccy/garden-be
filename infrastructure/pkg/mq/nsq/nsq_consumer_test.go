package nsq

import (
	"log"
	"testing"
	"time"
)

func TestNsq_Consumer(t *testing.T) {
	client := NewConsumerClient()
	err := client.SubscribeMsg("test_topic", func(bytes []byte) error {
		// 输出 nsq 中的消息
		log.Println(string(bytes))

		return nil
	})

	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(time.Duration(3) * time.Second)
}
