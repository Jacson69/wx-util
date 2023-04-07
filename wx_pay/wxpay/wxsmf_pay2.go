package wxpay

import (
	"crypto/md5"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	//. "yzsmz/base"
	//"yzsmz/gowf"
	log "github.com/haming123/wego/dlog"
)

// 微信扫码付请求参数
type WxSmfCreateReq2 struct {
	XMLName          xml.Name `xml:"xml"`
	Appid            string   `xml:"appid"`
	Mch_id           string   `xml:"mch_id"`
	Sub_mch_id       string   `xml:"sub_mch_id,omitempty"`
	Nonce_str        string   `xml:"nonce_str"`
	Body             string   `xml:"body"`
	Out_trade_no     string   `xml:"out_trade_no"`
	Total_fee        int      `xml:"total_fee"`
	Spbill_create_ip string   `xml:"spbill_create_ip"`
	Auth_code        string   `xml:"auth_code"`
	Sign             string   `xml:"sign"`
}

// 微信扫码付返回参数
type WxSmfCreateRet2 struct {
	Return_code    string `xml:"return_code"`
	Return_msg     string `xml:"return_msg"`
	Appid          string `xml:"appid"`
	Mch_id         string `xml:"mch_id"`
	Sub_appid      string `xml:"sub_appid"`
	Sub_mch_id     string `xml:"sub_mch_id"`
	Nonce_str      string `xml:"nonce_str"`
	Sign           string `xml:"sign"`
	Result_code    string `xml:"result_code"`
	Openid         string `xml:"openid"`
	Is_subscribe   string `xml:"is_subscribe"`
	Trade_type     string `xml:"trade_type"`
	Bank_type      string `xml:"bank_type"`
	Total_fee      int    `xml:"total_fee"`
	Cash_fee       int    `xml:"cash_fee"`
	Transaction_id string `xml:"transaction_id"`
	Out_trade_no   string `xml:"out_trade_no"`
	Time_end       string `xml:"time_end"`
}

// 微信扫码付md5签名
func (ent *WxSmfCreateReq2) GenWithdrawSign(keystr string) string {
	kvs := url.Values{}
	kvs.Add("appid", ent.Appid)
	kvs.Add("mch_id", ent.Mch_id)
	if len(ent.Sub_mch_id) > 0 {
		kvs.Add("sub_mch_id", ent.Sub_mch_id)
	}
	kvs.Add("nonce_str", ent.Nonce_str)
	kvs.Add("body", ent.Body)
	kvs.Add("out_trade_no", ent.Out_trade_no)
	kvs.Add("total_fee", strconv.Itoa(ent.Total_fee))
	kvs.Add("spbill_create_ip", ent.Spbill_create_ip)
	kvs.Add("auth_code", ent.Auth_code)

	tmpstr, _ := url.QueryUnescape(kvs.Encode())
	tmpstr = tmpstr + "&key=" + keystr
	//gowf.LogI(tmpstr)
	log.Info(tmpstr)
	unifiedorderSign := fmt.Sprintf("%x", md5.Sum([]byte(tmpstr)))
	unifiedorderSign = strings.ToUpper(unifiedorderSign)
	return unifiedorderSign
}

// 微信扫码付下单
func WxSmfCreate2(order WxSmfCreateReq2, strkey string) WxSmfCreateRet2 {

	var ret WxSmfCreateRet2
	order.Sign = order.GenWithdrawSign(strkey)
	xmlBody, _ := xml.MarshalIndent(order, " ", " ")
	post_data := string(xmlBody)
	//gowf.LogD(post_data)
	log.Info(post_data)
	wx_addr := "https://api.mch.weixin.qq.com/pay/micropay"
	res, err := http.Post(wx_addr, "charset=UTF-8", strings.NewReader(post_data))
	if err != nil {
		//gowf.LogD(err)
		log.Error(err)
		return ret
	}

	raw, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		//gowf.LogD(err)
		log.Error(err)
		return ret
	}
	//gowf.LogD(string(raw))
	log.Info(string(raw))
	if err := xml.Unmarshal(raw, &ret); err != nil {
		//gowf.LogD(err)
		log.Error(err)
		return ret
	}

	return ret
}

// 微信扫码付查询请求参数
type WxSmfQueryReq2 struct {
	XMLName      xml.Name `xml:"xml"`
	Appid        string   `xml:"appid"`
	Mch_id       string   `xml:"mch_id"`
	Sub_mch_id   string   `xml:"sub_mch_id,omitempty"`
	Out_trade_no string   `xml:"out_trade_no"`
	Nonce_str    string   `xml:"nonce_str"`
	Sign         string   `xml:"sign"`
}

// 微信扫码付查询返回参数
type WxSmfQueryRet2 struct {
	Return_code      string `xml:"return_code"`
	Return_msg       string `xml:"return_msg"`
	Appid            string `xml:"appid"`
	Mch_id           string `xml:"mch_id"`
	Sub_mch_id       string `xml:"sub_mch_id"`
	Nonce_str        string `xml:"nonce_str"`
	Sign             string `xml:"sign"`
	Result_code      string `xml:"result_code"`
	Err_code_des     string `xml:"err_code_des"`
	Openid           string `xml:"openid"`
	Is_subscribe     string `xml:"is_subscribe"`
	Trade_type       string `xml:"trade_type"`
	Trade_state      string `xml:"trade_state"`
	Bank_type        string `xml:"bank_type"`
	Total_fee        int    `xml:"total_fee"`
	Cash_fee         int    `xml:"cash_fee"`
	Transaction_id   string `xml:"transaction_id"`
	Out_trade_no     string `xml:"out_trade_no"`
	Time_end         string `xml:"time_end"`
	Trade_state_desc string `xml:"trade_state_desc"`
}

// 微信扫码付查询md5签名
func (ent *WxSmfQueryReq2) GenWithdrawSign(keystr string) string {
	kvs := url.Values{}
	kvs.Add("appid", ent.Appid)
	kvs.Add("mch_id", ent.Mch_id)
	if len(ent.Sub_mch_id) > 0 {
		kvs.Add("sub_mch_id", ent.Sub_mch_id)
	}
	kvs.Add("nonce_str", ent.Nonce_str)
	kvs.Add("out_trade_no", ent.Out_trade_no)

	tmpstr, _ := url.QueryUnescape(kvs.Encode())
	tmpstr = tmpstr + "&key=" + keystr
	//gowf.LogI(tmpstr)
	log.Info(tmpstr)

	unifiedorderSign := fmt.Sprintf("%x", md5.Sum([]byte(tmpstr)))
	unifiedorderSign = strings.ToUpper(unifiedorderSign)
	return unifiedorderSign
}

// 微信扫码付支付状态查询
func WxSmfQuery2(param WxSmfQueryReq2, strkey string) WxSmfQueryRet2 {

	param.Sign = param.GenWithdrawSign(strkey)
	xmlBody, _ := xml.MarshalIndent(param, " ", " ")
	//gowf.LogD(string(xmlBody))
	log.Info(string(xmlBody))

	var retparam WxSmfQueryRet2

	wx_addr := "https://api.mch.weixin.qq.com/pay/orderquery"
	res, err := http.Post(wx_addr, "charset=UTF-8", strings.NewReader(string(xmlBody)))
	if err != nil {
		//gowf.LogD(err)
		log.Error(err)
		return retparam
	}

	raw, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		//gowf.LogD(err)
		log.Error(err)
		return retparam
	}
	//gowf.LogD(string(raw))
	log.Info(string(raw))
	if err := xml.Unmarshal(raw, &retparam); err != nil {
		//gowf.LogD(err)
		log.Error(err)
		return retparam
	}

	return retparam
}

// 微信扫码付撤销请求参数
type WxSmfReverseReq2 struct {
	XMLName      xml.Name `xml:"xml"`
	Appid        string   `xml:"appid"`
	Mch_id       string   `xml:"mch_id"`
	Sub_mch_id   string   `xml:"sub_mch_id,omitempty"`
	Out_trade_no string   `xml:"out_trade_no"`
	Nonce_str    string   `xml:"nonce_str"`
	Sign         string   `xml:"sign"`
}

// 微信扫码付撤销返回参数
type WxSmfReverseRet2 struct {
	Return_code string `xml:"return_code"`
	Return_msg  string `xml:"return_msg"`
	Appid       string `xml:"appid"`
	Mch_id      string `xml:"mch_id"`
	Sub_appid   string `xml:"sub_appid"`
	Sub_mch_id  string `xml:"sub_mch_id"`
	Nonce_str   string `xml:"nonce_str"`
	Sign        string `xml:"sign"`
	Result_code string `xml:"result_code"`
	Recall      string `xml:"recall"`
}

// 微信扫码付撤销md5签名
func (ent *WxSmfReverseReq2) GenWithdrawSign(keystr string) string {
	kvs := url.Values{}
	kvs.Add("appid", ent.Appid)
	kvs.Add("mch_id", ent.Mch_id)
	if len(ent.Sub_mch_id) > 0 {
		kvs.Add("sub_mch_id", ent.Sub_mch_id)
	}
	kvs.Add("nonce_str", ent.Nonce_str)
	kvs.Add("out_trade_no", ent.Out_trade_no)

	tmpstr, _ := url.QueryUnescape(kvs.Encode())
	tmpstr = tmpstr + "&key=" + keystr
	//gowf.LogI(tmpstr)
	log.Info(tmpstr)
	unifiedorderSign := fmt.Sprintf("%x", md5.Sum([]byte(tmpstr)))
	unifiedorderSign = strings.ToUpper(unifiedorderSign)
	return unifiedorderSign
}

// 微信扫码付撤销
func WxSmfReverse2(param WxSmfReverseReq2, strkey string, pem_name string) WxSmfReverseRet2 {

	var retparam WxSmfReverseRet2

	param.Sign = param.GenWithdrawSign(strkey)
	xmlBody, _ := xml.MarshalIndent(param, " ", " ")
	//gowf.LogD(string(xmlBody))
	log.Info(string(xmlBody))
	wx_addr := "https://api.mch.weixin.qq.com/secapi/pay/reverse"
	res, err := SecurePost(wx_addr, xmlBody, pem_name)
	if err != nil {
		//gowf.LogD(err.Error())
		log.Error(err)
		return retparam
	}
	defer res.Body.Close()

	raw, err := ioutil.ReadAll(res.Body)
	if err != nil {
		//gowf.LogD(err)
		log.Error(err)
		return retparam
	}
	//gowf.LogD(string(raw))
	log.Info(string(raw))
	if err := xml.Unmarshal(raw, &retparam); err != nil {
		//gowf.LogD(err)
		log.Error(err)
		return retparam
	}

	return retparam
}
