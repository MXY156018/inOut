package express

//https://open.kuaidihelp.com/api/1063
// 快宝

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"mall-pkg/utils"
)

type KBSearchExpressReq struct {
	ExpressCode string `json:"express_code"`
	ExpressNo   string `json:"express_no"`
}
type KBSearchExpressDataReq struct {
	WaybillCodes string `json:"waybill_codes"`
	CpCode       string `json:"cp_code"`
	ResultSort   string `json:"result_sort"`
	Phone        string `json:"phone"`
}

type KBExpressData struct {
	Time    string `json:"time"`
	Context string `json:"context"`
	Status  string `json:"status"`
}

type KBSearchExpressData struct {
	WaybillCode string          `json:"waybill_code"`
	CpCode      string          `json:"cp_code"`
	Status      string          `json:"status"`
	Data        []KBExpressData `json:"data"`
	Order       string          `json:"order"`
	No          string          `json:"no"`
	Brand       string          `json:"brand"`
}

type KBSearchExpressResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type KuaiBao struct {
	Host    string
	AppId   string
	ApiKey  string
	Method  string
	AgentId string
}

type KuaiBaoErrtype int

const (
	Unkown KuaiBaoErrtype = 1000 + iota
	ServerErr
	Illegal
	NotExit
	AuthParamsErr
	SignExpired
	AuthErr
	AuthLimit
	FrequencyLimit
	ParamsErr
	InvalidStatus
	UnableRequest
	InvalidService
	Recognition KuaiBaoErrtype = 90003
)

func (et KuaiBaoErrtype) String() string {
	switch et {
	case Unkown:
		return "未知错误"
	case ServerErr:
		return "系统错误"
	case Illegal:
		return "非法请求"
	case NotExit:
		return "接口不存在或已废弃"
	case AuthParamsErr:
		return "授权参数错误"
	case SignExpired:
		return "授权签名已过期"
	case AuthErr:
		return "授权认证失败"
	case AuthLimit:
		return "权限受限"
	case FrequencyLimit:
		return "请求频率超过限制"
	case ParamsErr:
		return "请求参数错误"
	case InvalidStatus:
		return "无效的状态"
	case UnableRequest:
		return "无法完成请求"
	case InvalidService:
		return "无效的服务"
	case Recognition:
		return "识别失败"
	default:
		return "未知错误"
	}
}
func NewKuaiBao(Host, AppId, ApiKey, AgentId string) KuaiBao {
	return KuaiBao{
		Host:    Host,
		AppId:   AppId,
		ApiKey:  ApiKey,
		AgentId: AgentId,
	}
}
func (e *KuaiBao) SearchExpress(express_no string, express_code, phone string) (*KBSearchExpressResp, int, error) {
	var result KBSearchExpressResp
	var r http.Request
	method := "express.info.get"
	ts := time.Now().Unix()
	str := fmt.Sprintf("%s%s%s%s", e.AppId, method, fmt.Sprintf("%d", ts)[0:10], e.ApiKey)
	data := KBSearchExpressDataReq{
		WaybillCodes: express_no,
		// CpCode:       express_code,
		ResultSort: "0",
		Phone:      phone,
	}
	md5str := utils.MD5V([]byte(str))
	datastr, err := json.Marshal(data)
	if err != nil {
		return &result, 1, err
	}
	r.ParseForm()
	r.Form.Add("app_id", e.AppId)
	r.Form.Add("method", method)
	r.Form.Add("sign", md5str)
	r.Form.Add("data", string(datastr))
	r.Form.Add("ts", fmt.Sprintf("%d", ts)[0:10])
	bodystr := strings.TrimSpace(r.Form.Encode())
	req, err := http.NewRequest(http.MethodPost, e.Host, strings.NewReader(bodystr))
	if err != nil {
		return &result, 1, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	c := http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return &result, 1, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &result, 1, err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return &result, 1, err
	}
	if result.Code != 0 {
		var code = KuaiBaoErrtype(result.Code)
		msg := code.String()
		err = errors.New(msg)
	}
	return &result, result.Code, err
}
func (e *KuaiBao) PrintWayBill(content PrintWayBillReq) (PrintWayBillResp, error) {
	content.AgentId = e.AgentId
	var result PrintWayBillResp
	var r http.Request
	method := "cloud.print.waybill"
	ts := time.Now().Unix()
	str := fmt.Sprintf("%s%s%s%s", e.AppId, method, fmt.Sprintf("%d", ts)[0:10], e.ApiKey)
	md5str := utils.MD5V([]byte(str))

	datastr, err := json.Marshal(content)
	fmt.Printf("%+v\n", string(datastr))
	if err != nil {
		return result, err
	}
	r.ParseForm()
	r.Form.Add("app_id", e.AppId)
	r.Form.Add("method", method)
	r.Form.Add("sign", md5str)
	r.Form.Add("data", string(datastr))
	r.Form.Add("ts", fmt.Sprintf("%d", ts)[0:10])
	bodystr := strings.TrimSpace(r.Form.Encode())
	req, err := http.NewRequest(http.MethodPost, e.Host+"/api", strings.NewReader(bodystr))
	if err != nil {
		return result, err
	}
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	c := http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}
	fmt.Printf("%+v\n", result)
	if result.Code != 0 {
		var code = KuaiBaoErrtype(result.Code)
		msg := code.String()
		err = errors.New(msg)
	}
	return result, err
}

func (e *KuaiBao) GetWayBillCode(content WaybillCodeReq) (WaybillCodeResp, error) {
	content.AccountPassword = e.AgentId
	var result WaybillCodeResp
	var r http.Request
	method := "get.waybill.number"
	ts := time.Now().Unix()
	str := fmt.Sprintf("%s%s%s%s", e.AppId, method, fmt.Sprintf("%d", ts)[0:10], e.ApiKey)
	md5str := utils.MD5V([]byte(str))

	datastr, err := json.Marshal(content)
	fmt.Printf("%+v\n", string(datastr))
	if err != nil {
		return result, err
	}
	r.ParseForm()
	r.Form.Add("app_id", e.AppId)
	r.Form.Add("method", method)
	r.Form.Add("sign", md5str)
	r.Form.Add("data", string(datastr))
	r.Form.Add("ts", fmt.Sprintf("%d", ts)[0:10])
	bodystr := strings.TrimSpace(r.Form.Encode())
	req, err := http.NewRequest(http.MethodPost, e.Host+"/test", strings.NewReader(bodystr))
	if err != nil {
		return result, err
	}
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	c := http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}
	fmt.Printf("%+v\n", result)
	if result.Code != 0 {
		var code = KuaiBaoErrtype(result.Code)
		msg := code.String()
		err = errors.New(msg)
	}
	return result, err
}
