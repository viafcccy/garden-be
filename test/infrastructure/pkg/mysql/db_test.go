package mysql

import (
	"testing"

	"github.com/viafcccy/garden-be/global"
	_ "github.com/viafcccy/garden-be/test"
)

func TestInitDB(t *testing.T) {
	err := global.GDB.DB().Ping()
	if err != nil {
		t.Fatal("mysql conn err :", err)
	}
}
