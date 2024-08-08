package express

type ShipAddress struct {
	Name     string `json:"name,omitempty"`
	Mobile   string `json:"mobile,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Province string `json:"province,omitempty"`
	District string `json:"district,omitempty"`
	City     string `json:"city,omitempty"`
	Detail   string `json:"detail,omitempty"`
}

type SearchExpressReq struct {
	ExpressCode string `json:"express_code"`
	ExpressNo   string `json:"express_no"`
}
type SearchExpressDataReq struct {
	WaybillCodes string `json:"waybill_codes"`
	CpCode       string `json:"cp_code"`
	ResultSort   string `json:"result_sort"`
}

type ExpressData struct {
	Time    string `json:"time"`
	Context string `json:"context"`
	Status  string `json:"status"`
}

type SearchExpressData struct {
	WaybillCode string        `json:"waybill_code"`
	CpCode      string        `json:"cp_code"`
	Status      string        `json:"status"`
	Data        []ExpressData `json:"data"`
	Order       string        `json:"order"`
	No          string        `json:"no"`
	Brand       string        `json:"brand"`
}

type SearchExpressResp struct {
	Code int                 `json:"code"`
	Msg  string              `json:"msg"`
	Data []SearchExpressData `json:"data"`
}
type ConsolidationData struct {
	Name string `json:"name"` //集包地名称
	Code string `json:"code"` //集包地代码
}
type SortationData struct {
	Name string `json:"name"` //大字与一段码
}
type OriginData struct {
	Code string `json:"code,omitempty"` //
	Name string `json:"name,omitempty"` //
}
type RoutingInfoData struct {
	Consolidation ConsolidationData `json:"consolidation,omitempty"` //本单号的集包地信息
	RouteCode     string            `json:"route_code,omitempty"`    //本单号的二三段码
	Sortation     SortationData     `json:"sortation,omitempty"`     //本单号的大字与一段码
	Origin        OriginData        `json:"origin,omitempty"`        //
}
type OrifinalData struct {
	EncryptedData string `json:"encryptedData"` //加密的打印数据
	Signature     string `json:"signature"`     //加密打印数据的签名
	TemplateURL   string `json:"templateURL"`   //模板文件url，在自行打印时可替换为所需的其他模板
	Ver           string `json:"ver"`           //菜鸟组件版本号，仅当第三方平台为菜鸟时返回
	Type          string `json:"type"`          //加密类型 cn_encrypted:菜鸟
}
type TaskInfoData struct {
	Tid          string          `json:"tid"`          //任务ID，为打印任务中输入的值
	WaybillCode  string          `json:"waybill_code"` //申请到的运单号
	TemplateId   string          `json:"template_id"`  //本次所用的打印模板ID
	RoutingInfo  RoutingInfoData `json:"routing_info"` //本单号的路由、分拣信息（品牌为非顺丰）
	OrifinalData OrifinalData    `json:"originalData"` //菜鸟、拼多多、抖音、快手、京东等第三方平台返回的原始加密打印数据，连接各第三方平台打印组件打印时使用
}
type PrintWayBillData struct {
	Status       string       `json:"status"`       //处理状态，success为成功，其他为失败
	TaskId       string       `json:"task_id"`      //本次打印任务编号，在查询打印任务状态时需要
	TaskInfo     TaskInfoData `json:"task_info"`    //此次打印任务详情，只有通过云打印机上配置的单号源获取单号成功后,才返回此字段
	PreviewImage string       `json:"PreviewImage"` //底单文件URL，仅当有底单文件（jpg或pdf格式）生成时才返回
}

type PrintWayBillResp struct {
	Code int              `json:"code"` //响应状态码。0-成功，非0-失败
	Msg  string           `json:"msg"`  //返回结果说明
	Uid  string           `json:"uid"`  //本次请求唯一业务流水号
	Data PrintWayBillData `json:"data"` //JSON格式响应数据
}
type AddressData struct {
	Province string `json:"province"` //发件人省份
	District string `json:"district"` //发件人县级名称
	City     string `json:"city"`     //发件人市级名称
	Detail   string `json:"detail"`   //发件人详细地址
}
type SenderData struct {
	Address AddressData `json:"address"` //发件人地址
	Mobile  string      `json:"mobile"`  //发件人手机
	Name    string      `json:"name"`    //发件人姓名
	Phone   string      `json:"phone"`   // 发件人固话
}
type RecipientData struct {
	SenderData
	Oaid string `json:"oaid,omitempty"` //淘宝加密订单获取单号时必传，收件人ID
}
type PackageData struct {
	Name   string `json:"name"`   //物品名称，最长50个字符
	Weight string `json:"weight"` //包裹重量，单位为公斤
}
type PrintData struct {
	Tid         string        `json:"tid"`            //任务ID，最长30个字符。建议用订单号之类的唯一标识，提交和回调时将返回该字段
	WaybillCode string        `json:"waybill_code"`   //运单号。若为空，则使用云打印机上绑定的单号源自动获取运单号，参见示例代码：结合小邮筒配置单号源打印
	CpCode      string        `json:"cp_code"`        //快递品牌，快递公司标识符，如：zt
	Sender      SenderData    `json:"sender"`         //发件人信息
	Recipient   RecipientData `json:"recipient"`      //收件人信息
	Package     PackageData   `json:"package"`        // 包裹信息
	Note        string        `json:"note,omitempty"` //备注
	// GoodsName   string          `json:"goods_name,omitempty"`   //
	// Weight      string          `json:"weight,omitempty"`       //
	// Sequence    string          `json:"sequence,omitempty"`     //
	// CodeAmount  int             `json:"code_amount,omitempty"`  //
	// RoutingInfo RoutingInfoData `json:"routing_info,omitempty"` //
}

type PrintWayBillReq struct {
	AgentId   string `json:"agent_id"`   // 目标云打印机的访问密钥
	PrintType int    `json:"print_type"` // 打印类型，1：仅生成底单，jpg格式； 2：仅打印； 3：打印并生成jpg格式底单； 7：仅生成底单，pdf格式；默认为3
	// Private     int         `json:"private"`       //是否隐藏收发件人联系电话 1 面单上隐藏 2 底单图片上隐藏 3 面单和底单图片都隐藏；若都不隐藏，不传值即可；
	// TemplateId  string      `json:"template_id"`   //打印模板ID
	// CustomTplId string      `json:"custom_tpl_id"` //自定义区域模板id，该id从微信小程序【标签打王】获取。当template_id指定的模板支持自定义区域时才有效
	// UserName    string      `json:"user_name"`     //打印人名称
	PrintData []PrintData `json:"print_data"` //打印数据，每条打印数据为按如下格式提供的一个数组元素，一次最多50条打印数据，若要获取单号，建议一次不要超过10条
}
type OrderData struct {
	OrderId   string        `json:"order_id"`
	Sender    SenderData    `json:"sender"`
	Recipient RecipientData `json:"recipient"`
	Package   PackageData   `json:"package"`
}
type WaybillCodeReq struct {
	CpCode          string      `json:"cp_code"`
	AccountNo       string      `json:"account_no"`
	AccountPassword string      `json:"account_password"`
	OrderData       []OrderData `json:"order_data"`
}
type WaybillCodeResp struct {
	Code int         `json:"code"` //响应状态码。0-成功，非0-失败
	Msg  string      `json:"msg"`  //返回结果说明
	Uid  string      `json:"uid"`  //本次请求唯一业务流水号
	Data interface{} `json:"data"` //JSON格式响应数据
}

type SendEasyOrderReq struct {
	BuyerMessage     string        `json:"buyerMessage"`     //买家留言
	BuyerNick        string        `json:"buyerNick"`        //买家昵称
	DiscountAmount   string        `json:"discountAmount"`   //优惠金额 单位：元
	OrdersGoods      []OrdersGoods `json:"ordersGoods"`      //商品信息
	PayAmount        string        `json:"payAmount"`        //实收金额（请传用户实际支付的金额，包含运费） 单位：元
	PayType          string        `json:"payType"`          // 01 在线支付 02 货到付款
	PostAmount       string        `json:"postAmount"`       //运费 单位：元
	ReceiverAddress  string        `json:"receiverAddress"`  //收件人地址
	ReceiverCity     string        `json:"receiverCity"`     //收件人城市
	ReceiverCounty   string        `json:"receiverCounty"`   //收件人区/县级市
	ReceiverMobile   string        `json:"receiverMobile"`   //收件人手机（与电话二选一入参即可）
	ReceiverName     string        `json:"receiverName"`     //收件人名称
	ReceiverProvince string        `json:"receiverProvince"` //收件人省份
	SellerMessage    string        `json:"sellerMessage"`    //卖家备注
	ShipperAddress   string        `json:"shipperAddress"`   //发件人地址
	ShipperCity      string        `json:"shipperCity"`      //发件人城市
	ShipperCounty    string        `json:"shipperCounty"`    //发件人区/县级市
	ShipperMobile    string        `json:"shipperMobile"`    //发件人手机（与电话二选一入参即可）
	ShipperName      string        `json:"shipperName"`      //发件人名称
	ShipperProvince  string        `json:"shipperProvince"`  //发件人省份
	ShopID           int64         `json:"shopId"`           //店铺ID
	TotalAmount      string        `json:"totalAmount"`      //合计金额（优惠前的金额，合计金额=商品单价*数量-优惠金额） 单位：元
	TradeNo          string        `json:"tradeNo"`          //平台交易单号
	TradeStatus      string        `json:"tradeStatus"`      //01(待付款);02(待发货);03(已发货);04(退款中);05（部分发货）;06(部分退款);10(交易成功);11(交易关闭)
}
type OrdersGoods struct {
	Oid            string `json:"oid"`            //子单号
	PlatformSkuID  string `json:"platformSkuId"`  //平台商品SkuId或平台商品Id
	GoodsName      string `json:"goodsName"`      //平台商品标题
	GoodsNo        string `json:"goodsNo"`        //平台商品Sku编码 (非必填)
	Quantity       int    `json:"quantity"`       //商品数量
	Price          string `json:"price"`          //商品单价
	RealPrice      string `json:"realPrice"`      //实际售价 单位：元
	PicPath        string `json:"picPath"`        //图片地址链接 (非必填)
	GoodsProps     string `json:"goodsProps"`     //商品属性 (非必填)
	ShortGoodsName string `json:"shortGoodsName"` //商品简称(非必填)
	Unit           string `json:"unit"`           //单位(非必填)
	Weight         int    `json:"weight"`         //商品重量(非必填)
}
type SendEasyOrderResp struct {
	Code    string      `json:"code"` //200
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
}
type FullTraceDetails struct {
	RouteAddress string `json:"routeAddress"` //当前地点
	RouteInfo    string `json:"routeInfo"`    //物流详情
	RouteDate    string `json:"routeDate"`    //路由发生的时间
}
type RouterData struct {
	LogisticsNo         string             `json:"logisticsNo"`         //运单号
	CpCode              string             `json:"cpCode"`              //物流公司
	LogisticsStatus     string             `json:"logisticsStatus"`     //物流状态
	LogisticsStatusDesc string             `json:"logisticsStatusDesc"` //物流状态描述
	FullTraceDetails    []FullTraceDetails `json:"fullTraceDetails"`    //详细路由数据
}
type QueryRouterReq struct {
	ShopID      int64  `json:"shopId"`      //店铺id
	LogisticsNo string `json:"logisticsNo"` //运单号
}
type QueryRouterResp struct {
	Code    int     `json:"code"`    //错误编码
	Success bool       `json:"success"` //成功为true,失败为false
	Msg     string     `json:"msg"`     //提示信息
	Data    RouterData `json:"data"`    //成功数据统一返回参数,失败为null
}
