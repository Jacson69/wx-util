package wxpay

import (
	"crypto/md5"
	"encoding/xml"
	"fmt"
	log "github.com/haming123/wego/dlog"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// 退款申请请求参数
type WxSmfRefundReq2 struct {
	XMLName       xml.Name `xml:"xml"`
	Appid         string   `xml:"appid"`
	Mch_id        string   `xml:"mch_id"`
	Sub_mch_id    string   `xml:"sub_mch_id,omitempty"`
	Nonce_str     string   `xml:"nonce_str"`
	Sign          string   `xml:"sign"`
	Out_trade_no  string   `xml:"out_trade_no"`
	Out_refund_no string   `xml:"out_refund_no"`
	Total_fee     int      `xml:"total_fee"`
	Refund_fee    int      `xml:"refund_fee"`
}

// 退款申请返回参数
type WxSmfRefundRet2 struct {
	Return_code    string `xml:"return_code"`
	Return_msg     string `xml:"return_msg"`
	Result_code    string `xml:"result_code"`
	Err_code_des   string `xml:"err_code_des"`
	Appid          string `xml:"appid"`
	Mch_id         string `xml:"mch_id"`
	Sub_mch_id     string `xml:"sub_mch_id"`
	Nonce_str      string `xml:"nonce_str"`
	Sign           string `xml:"sign"`
	Transaction_id string `xml:"transaction_id"`
	Out_trade_no   string `xml:"out_trade_no"`
	Out_refund_no  string `xml:"out_refund_no"`
	Refund_id      string `xml:"refund_id"`
	Refund_fee     int    `xml:"refund_fee"`
	Total_fee      int    `xml:"total_fee"`
	Cash_fee       int    `xml:"cash_fee"`
}

// 退款申请md5签名:服务商版
func (ent *WxSmfRefundReq2) GenWithdrawSign(keystr string) string {
	kvs := url.Values{}
	kvs.Add("appid", ent.Appid)
	kvs.Add("mch_id", ent.Mch_id)
	if len(ent.Sub_mch_id) > 0 {
		kvs.Add("sub_mch_id", ent.Sub_mch_id)
	}
	kvs.Add("nonce_str", ent.Nonce_str)
	kvs.Add("out_trade_no", ent.Out_trade_no)
	kvs.Add("out_refund_no", ent.Out_refund_no)
	kvs.Add("total_fee", strconv.Itoa(ent.Total_fee))
	kvs.Add("refund_fee", strconv.Itoa(ent.Refund_fee))

	tmpstr, _ := url.QueryUnescape(kvs.Encode())
	tmpstr = tmpstr + "&key=" + keystr
	//gowf.LogI(tmpstr)
	log.Info(tmpstr)
	unifiedorderSign := fmt.Sprintf("%x", md5.Sum([]byte(tmpstr)))
	unifiedorderSign = strings.ToUpper(unifiedorderSign)
	return unifiedorderSign
}

// 微信退款申请
func WxSmfRefund2(order WxSmfRefundReq2, strkey string, pem_name string) WxSmfRefundRet2 {

	order.Sign = order.GenWithdrawSign(strkey)
	xmlBody, _ := xml.MarshalIndent(order, " ", " ")
	//gowf.LogD(string(xmlBody))
	log.Info(string(xmlBody))
	var ret WxSmfRefundRet2
	wx_addr := "https://api.mch.weixin.qq.com/secapi/pay/refund"
	res, err := SecurePost(wx_addr, xmlBody, pem_name)
	if err != nil {
		//gowf.LogD(err.Error())
		log.Error(err.Error())
		return ret
	}
	defer res.Body.Close()

	raw, err := ioutil.ReadAll(res.Body)
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

// 微信退款申请查询请求参数
type WxSmfRefundQueryReq2 struct {
	XMLName       xml.Name `xml:"xml"`
	Appid         string   `xml:"appid"`
	Mch_id        string   `xml:"mch_id"`
	Sub_mch_id    string   `xml:"sub_mch_id,omitempty"`
	Nonce_str     string   `xml:"nonce_str"`
	Sign          string   `xml:"sign"`
	Out_refund_no string   `xml:"out_refund_no"`
}

// 微信退款申请请求参数md5签名
func (ent *WxSmfRefundQueryReq2) GenWithdrawSign(keystr string) string {
	kvs := url.Values{}
	kvs.Add("appid", ent.Appid)
	kvs.Add("mch_id", ent.Mch_id)
	if len(ent.Sub_mch_id) > 0 {
		kvs.Add("sub_mch_id", ent.Sub_mch_id)
	}
	kvs.Add("nonce_str", ent.Nonce_str)
	kvs.Add("out_refund_no", ent.Out_refund_no)

	tmpstr, _ := url.QueryUnescape(kvs.Encode())
	tmpstr = tmpstr + "&key=" + keystr
	//gowf.LogI(tmpstr)
	log.Info(tmpstr)
	unifiedorderSign := fmt.Sprintf("%x", md5.Sum([]byte(tmpstr)))
	unifiedorderSign = strings.ToUpper(unifiedorderSign)
	return unifiedorderSign
}

// 微信退款申请查询:服务商版
func WxSmfRefundQuery2(param WxSmfRefundQueryReq2, strkey string) Params {

	var retparam Params

	param.Sign = param.GenWithdrawSign(strkey)
	xmlBody, _ := xml.MarshalIndent(param, " ", " ")
	//gowf.LogD(string(xmlBody))
	log.Info(string(xmlBody))

	wx_addr := "https://api.mch.weixin.qq.com/pay/refundquery"
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

	//LogD(string(raw))
	retparam = XmlToMap(string(raw))
	return retparam
}
