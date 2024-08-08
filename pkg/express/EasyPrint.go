package express

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	// "io/ioutil"
	"mall-pkg/utils"
	"net/http"
	"strings"
	"time"
)

type EasyPrintOrder struct {
	Host         string
	ShopId       int64
	ClientId     string
	ClientSecret string
	Format       string
	ParentId     string
	V            string
}

func NewEasyPrintOrder(host string, shopid int64, clientId string, clientSecret string, format string, partner_id string, v string) *EasyPrintOrder {
	return &EasyPrintOrder{
		Host:         host,
		ShopId:       shopid,
		ClientId:     clientId,
		ClientSecret: clientSecret,
		Format:       format,
		ParentId:     partner_id,
		V:            v,
	}
}

func (e *EasyPrintOrder) SendOrder(params *SendEasyOrderReq) error {
	//fmt.Printf("%+v\n", params)
	var result SendEasyOrderResp
	method := "foonsu.ydd.order.upload"
	ts := time.Now().Unix()
	sign := e.Sign(ts, method)
	host := e.SetHost(method, ts, sign)
	var req *http.Request
	params.ShopID = e.ShopId
	bodymarshal, err := json.Marshal(params)
	if err != nil {
		return err
	}

	reqbody := strings.NewReader(string(bodymarshal))
	req, err = http.NewRequest(http.MethodPost, host, reqbody)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	c := http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return err
	}
	if result.Code != "200" || !result.Success {
		return nil
	}
	return nil
}

func (e *EasyPrintOrder) QueryRouter(logisticsNo string) (*RouterData, error) {
	var result = &QueryRouterResp{}
	method := "foonsu.ydd.router.query"
	ts := time.Now().Unix()
	sign := e.Sign(ts, method)
	host := e.SetHost(method, ts, sign)
	var req *http.Request
	params := &QueryRouterReq{
		ShopID:      e.ShopId,
		LogisticsNo: logisticsNo,
	}
	bodymarshal, err := json.Marshal(params)
	if err != nil {
		return &result.Data, err
	}

	reqbody := strings.NewReader(string(bodymarshal))
	req, err = http.NewRequest(http.MethodPost, host, reqbody)
	if err != nil {
		return &result.Data, err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	c := http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return &result.Data, err
	}
	defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &result.Data, err
	}
	err = json.Unmarshal(body, result)
	if err != nil {
		return &result.Data, err
	}
	if result.Code != 200 || !result.Success {
		return &result.Data, errors.New(result.Msg)
	}
	var code = CpCode(result.Data.CpCode)
	result.Data.CpCode = code.String()
	return &result.Data, nil
}

func (e *EasyPrintOrder) Sign(ts int64, method string) string {
	str := fmt.Sprintf("%sclient_id%sformat%smethod%spartner_id%stimestamp%dv%s%s", e.ClientSecret, e.ClientId, e.Format, method, e.ParentId, ts, e.V, e.ClientSecret)
	sign := utils.MD5V([]byte(str))
	sign = strings.ToLower(sign)
	return sign
}

func (e *EasyPrintOrder) SetHost(method string, ts int64, sign string) string {
	host := fmt.Sprintf("%s?client_id=%s&format=%s&method=%s&partner_id=%s&timestamp=%d&v=%s&sign=%s", e.Host, e.ClientId, e.Format, method, e.ParentId, ts, e.V, sign)
	return host
}
