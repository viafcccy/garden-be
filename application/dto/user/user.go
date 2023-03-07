package dto

// 请求对象
type (
	SimpleUserInfoReq struct {
		Id uint64 `form:"id" binding:"required"`
	}
	LoginReq struct {
		UserName string `json:"userName"`
		Password string `json:"password"`
	}
)

// 输出对象
type (
	SimpleUserInfoRsp struct {
		Id       uint64 `json:"id"`
		NickName string `json:"nickName"`
	}
	LoginRsp struct {
		UserName string `json:"userName"`
		NickName string `json:"nickName"`
		Token    string `json:"token"`
	}
)
