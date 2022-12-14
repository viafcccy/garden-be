package publish

import (
	"testing"

	_ "github.com/viafcccy/garden-be/boot"
)

// UserNsqPubDemo 测试发送消息案例
func TestUser_UserNsqPubDemo(t *testing.T) {
	err := UserNsqPubDemo(1)
	if err != nil {
		t.Fatal(err)
	}
}
