package response

import "github.com/viafcccy/garden-be/infrastructure/pkg/errno"

// 错误类型固定返回映射字典
var ERR_RSP_MAP = map[error]*Response{
	errno.ErrLogFailWrongPwd: ErrLoginFailWrongPwd,
}
