package log

import (
	"testing"

	log2 "github.com/viafcccy/garden-be/infrastructure/pkg/log"
	_ "github.com/viafcccy/garden-be/test"
)

func TestLog(t *testing.T) {
	log := log2.NewLogger()
	log.Errorf("ddddd")
}
