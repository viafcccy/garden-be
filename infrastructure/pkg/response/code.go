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
	OK  = response(0, "ok")        // 通用成功
	Err = response(500, "服务器内部错误") // 通用错误

	// 服务级错误码
	ErrParam     = response(10001, "参数有误")
	ErrSignParam = response(10002, "签名参数有误")

	// 模块级错误码 - 用户模块
	ErrUserService = response(20100, "用户服务异常")
	ErrUserPhone   = response(20101, "用户手机号不合法")
	ErrUserCaptcha = response(20102, "用户验证码有误")

	// 库存模块
	ErrOrderService = response(20200, "订单服务异常")
	ErrOrderOutTime = response(20201, "订单超时")
)
