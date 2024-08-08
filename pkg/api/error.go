package api

const (
	// 成功
	Error_OK int = iota
	// 服务器内部错误
	Error_Server
	// 非法 Token
	Error_InvalidToken
	// token 超时
	Error_TokenExpire
	//  用户不存在
	Error_AccountNotExists
	// 用户名或者密码错误
	Error_AccountOrPassword
	// 账号被禁
	Error_AccountForbid
	// 手机号已经存在
	Error_PhoneExists
	// 邀请人不存在
	Error_InviterNotExists
	//手机号格式错误
	Error_PhoneInvalid
	//收货人格式错误
	Error_Address_NameInvalid
	//详细地址错误
	Error_Address_DetailInvalid
	//地址不存在
	Error_Address_NotExists
	//图形验证码错误
	Error_CaptchaInvalid
	//没有权限
	Error_NoPrivilege
	//参数错误
	Error_Parameter
	//账号审核中
	Err_AccountAuthing
	// 微信号已绑定
	Err_WechatBinded
	// 账号已经存在
	Err_AccountExists
)

const (
	// 商品丢失(无权限或者找不到)
	Error_User_ProductMiss = 1000
	//无权限
	Error_User_NoPrivilege = 1001
	//地址错误
	Error_User_AddressNotExists = 1002
	//购物商品为空
	Error_User_NoProduct = 1003
	// 无库存
	Error_User_NoSku = 1004
	//邀请人错误
	Error_User_Refer = 1005
	// 订单已经支付
	Error_User_OrderPay = 1006
	// 订单不存在
	Error_User_OrderNotExists = 1007
	// 不能删除未完成订单
	Error_User_DelOrder = 1008
	// 访问太频繁
	Error_User_TooFreq = 1009
	// 商品已经评论
	Error_User_OrderComment = 1010
	// 订单未完成，不能评论
	Error_User_OrderNotFinish = 1011
	// 订单未发货
	Error_User_OrderNotDelivery = 1012
)

// 管理员专用的错误码
const (
	Error_Admin_ApiNotExists = 10000
	// API 已经存在
	Error_Admin_ApiExists = 10001
	// 角色权限已经存在
	Error_Admin_AuthorityExists = 10002
	// 角色存在子角色
	Error_Admin_AuthorityHasChild = 10003
	// 用户不存在
	Error_Admin_UserNotExists = 10004
	//密码错误
	Error_Admin_Password = 10005
	//不能删除自己
	Error_Admin_DelSelf = 10006
	//用户已经存在
	Error_Admin_UserExists = 10007
	//菜单已经存在
	Error_Admin_MenuExists = 10008
	//存在子菜单
	Error_Admin_MenuHasChild = 10009
	//商品分类已经存在
	Error_Admin_CategoryExists = 10010
	//商品分类存在引用
	Error_Admin_CategoryHasRef = 10011
	//商品存在子类
	Error_Admin_CategoryHasChild = 10012

	// 订单已经支付
	Error_Admin_OrderPayed = 10013
	// 订单已经发货
	Error_Admin_Deliveryed = 10014
	//订单未设置收货地址
	Error_Admin_NoAddress = 10015
	// 存在库存，不更更改规格类型
	Error_Admin_HaveSku = 10016
	// 不是商户的商品
	Error_Admin_NotYourProduct = 10017
	//订单已取消
	Error_Admin_OrderCancel = 10018
	// 订单状态
	Error_Admin_OrderStatus = 10019
)

var InvalidParameterResp = &BaseResp{
	Code: Error_Parameter,
	Msg:  "参数错误",
}

func NewInvalidParameter(msg string) *BaseResp {
	return &BaseResp{
		Code: Error_Parameter,
		Msg:  "参数解析错误 : " + msg,
	}
}
