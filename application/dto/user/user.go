package dto

// 输出对象
type (
	SimpleUserInfo struct {
		Id       uint64 `json:"id"`
		NickName string `json:"nickName"`
	}
)

// 请求对象
type (
	SimpleUserInfoReq struct {
		Id uint64 `form:"id" binding:"required"`
	}
)
