package response

// 错误码规则：
// (1) 错误码需为 > 0 的数;
//
// (2) 错误码为 5 位数：
//              ----------------------------------------------------------
//                  第 1 位               2、3 位                  4、5 位
//              ----------------------------------------------------------
//                服务级错误码          模块级错误码	         具体错误码
//              ----------------------------------------------------------

var (
	// OK
	OK               = response(0, "ok")     // 通用成功
	Err              = response(500, "fail") // 通用错误
	ErrInvalidParams = response(10101, "请求参数错误")

	// 1 服务级错误码
	// 01 token
	ErrAuthCheckTokenFail    = response(10101, "Token 鉴权失败")
	ErrAuthCheckTokenTimeout = response(10102, "Token 已超时")
	ErrAuthToken             = response(10103, "Token 生成失败")
	ErrAuth                  = response(10104, "Token 错误")
	ErrNotCarryToken         = response(10105, "未携带 Token 请求")

	// 2 模块级错误码
	// 01 用户模块
	ErrExistUserName     = response(20101, "用户名已存在")
	ErrLoginFailWrongPwd = response(20102, "用户密码错误")
	// 02 文章模块
	ErrNotExistArticle = response(20201, "该文章不存在")
	ErrExistTag        = response(20202, "已存在该标签名称")
	ErrNotExistTag     = response(20203, "该标签不存在")
)
