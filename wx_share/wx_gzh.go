package wx_share

import (
	"strconv"
	"strings"
	"time"
)

type WxSignature struct {
	Noncestr  string `json:"noncestr"`
	Timestamp string `json:"timestamp"`
	Signature string `json:"signature"`
	AppID     string `json:"appId"`
}

func GetWxSign(url string) (WxSignature, error) {
	app_id := "wxb79d4e2eee8c83ab"
	app_secret := "725216c759f923751a5573aefa9f425c"

	token, err := GetAccessToken(app_id, app_secret)
	if err != nil {
		return WxSignature{}, err

	}
	jsapi_ticket, err := GetTicket(app_id, token)
	if err != nil {
		return WxSignature{}, err
	}
	jsapi_ticket = strings.Trim(jsapi_ticket, "\n")
	jsapi_ticket = strings.Trim(jsapi_ticket, "\"")

	var Signature WxSignature
	noncestr := RandStringBytes(16)
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	signatureStr := "jsapi_ticket=" + jsapi_ticket + "&noncestr=" + noncestr + "&timestamp=" + timestamp + "&url=" + url

	Signature.Signature = GetSha1(signatureStr)
	Signature.Noncestr = noncestr
	Signature.Timestamp = timestamp
	Signature.AppID = app_id
	return Signature, nil
}
