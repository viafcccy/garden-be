package errno

import (
	"errors"
)

// 常用错误类型

// 密码错误导致登录失败
var ErrLogFailWrongPwd = errors.New("ErrLogFailWrongPwd")
