package api

// context 里的uid字段
const (
	// 用户UID
	Context_Key_UID = "uid"
	//用户等级
	Context_Key_Grade = "grade"
	// 商户权限
	Context_Key_Privilege = "privilege"
	// 商户ID
	Context_Key_MerchantId = "merchant"
	// 商户是否登陆
	Context_Key_MerchantLogin = "merchant_login"
)

// 商户权限
const (
	// 无任何权限(需要填写信息及审核)
	MerchantPrivilege_Non int = 0
	// 商品权限
	MerchantPrivilege_Product int = 1 << 0
)

const (
	// http 头 x-token
	Req_Header_Token1 = "x-token"
	// http 头 token
	Req_Header_Token2 = "token"
	// 更新 token
	Resp_Header_NewToken = "new-token"
	// 新 token 超时时间
	Resp_Header_NewTokenExpire = "new-expires-at"
	// 中间件增加的头信息 - claims
	Middle_Header_Claims = "claims"
	//权限
	Middle_Header_AuthorityId = "AuthorityId"
)

// 通用回复
type BaseResp struct {
	// 错误码
	Code int `json:"code"`
	//错误信息
	Msg string `json:"msg,omitempty"`
	//是否需要登录
	Reload bool `json:"reload,omitempty"`
	//数据
	Data interface{} `json:"data,omitempty"`
}

// 分页请求
type PageQuery struct {
	//页数
	Page int `json:"page,optional" form:"page"`
	//每页数量
	PageSize int `json:"pageSize" form:"pageSize"`
}

// 分页回复
type PageResp struct {
	//总条目数量
	Total int64 `json:"total"`
	//数据列表
	List     interface{} `json:"list,omitempty"`
	Page     int         `json:"page,omitempty"`
	PageSize int         `json:"pageSize,omitempty"`
}

// 主键查询方式
type PageQueryById struct {
	// 参考ID(主要用于主键方式查询的简单查询方式)
	RefId int `json:"refId,optional" form:"refId,optional"`
	// 是否是向前查询
	IsPrev bool `json:"isPrev,optional" form:"isPrev,optional"`
	//每页数量
	PageSize int `json:"pageSize" form:"pageSize"`
}

// ID请求
type IDReq struct {
	ID int `json:"id" form:"id"`
}

// ID请求
type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

type IDStrReq struct {
	ID string `json:"id" form:"id"`
}
type ExpressStateReq struct {
	Id        int    `json:"id" form:"id"`
	ExpressNo string `json:"express_no" form:"express_no"`
}
