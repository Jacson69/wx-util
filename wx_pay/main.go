package wxpay

import (
	"github.com/haming123/wego"
	log "github.com/haming123/wego/dlog"
	"github.com/haming123/wego/worm"
	"strings"
	"wechat/wx_util/wx_pay/dao"
	"wechat/wx_util/wx_pay/wxpay"
)

type WxMchParam struct {
	MchAppid  string `ini:"mchapp_id"`
	MchId     string `ini:"mch_id"`
	MchSecret string `ini:"mch_secret"`
	SubAppid  string `ini:"subapp_id"`
}

var WxMch WxMchParam
var shop dao.Shop
var order dao.JzOrder

// 提交微信支付
func WxPay(c *wego.WebContext) {
	dbs := worm.NewSession()
	type InParam struct {
		JzOrderId int64
		QrCode    string
		Pay_price float64
	}

	var pin InParam
	err := c.ReadJSON(&pin)
	if err != nil {
		log.Error(err)
		c.AbortWithError(510, err)
		return
	}
	dbs.TxBegin()
	order.Paying = int(pin.Pay_price * 100)
	//修改订单为：支付中
	order.Paying_no = dao.GenPayOrderTradeNo("CF", pin.JzOrderId)
	err = dao.SetJzddPaying(dbs, pin.JzOrderId, order.Paying, "sfs", order.Paying_no)
	if err != nil {
		log.Error("修改订单支付状态错误", err)
		c.AbortWithError(510, err)
		dbs.TxRollback()
		return
	}
	var preq wxpay.WxSmfCreateReq2
	mch_secret := ""
	if len(shop.Wx_mchkey) > 0 {
		//普通商户模式支付
		preq.Appid = WxMch.SubAppid
		preq.Mch_id = shop.Wx_mchid
		mch_secret = shop.Wx_mchkey
	} else {
		//服务商模式模式支付
		preq.Appid = WxMch.MchAppid
		preq.Mch_id = WxMch.MchId
		preq.Sub_mch_id = shop.Wx_mchid
		mch_secret = WxMch.MchSecret
	}

	preq.Auth_code = pin.QrCode
	preq.Body = "就诊订单支付"
	preq.Nonce_str = wxpay.GetNonceStr(32)
	preq.Out_trade_no = order.Paying_no
	preq.Spbill_create_ip = c.Input.ClientIP()
	preq.Total_fee = order.Paying
	pret := wxpay.WxSmfCreate2(preq, mch_secret)
	if pret.Return_code != "SUCCESS" {
		msg := "微信支付请求错误"
		log.Error(msg)
		c.AbortWithText(510, msg)
		dbs.TxRollback()
		return
	}
	dbs.TxCommit()
}

// 微信订单查询
func HandlerJzOrderWxPayQuery(c *wego.WebContext) {
	dbs := worm.NewSession()

	type InParam struct {
		JzOrderId int64
	}

	var pin InParam
	err := c.ReadJSON(&pin)
	if err != nil {
		log.Error(err)
		c.AbortWithError(510, err)
		return
	}

	trade_state := wxpay.WxSmfQueryRet2{}
	trade_state.Trade_state = "SUCCESS"
	if order.Paying > 0 {
		mch_secret := ""
		var preq wxpay.WxSmfQueryReq2
		if len(shop.Wx_mchkey) > 0 {
			//普通商户模式支付
			preq.Appid = WxMch.SubAppid
			preq.Mch_id = shop.Wx_mchid
			mch_secret = shop.Wx_mchkey
		} else {
			//服务商模式模式支付
			preq.Appid = WxMch.MchAppid
			preq.Mch_id = WxMch.MchId
			preq.Sub_mch_id = shop.Wx_mchid
			mch_secret = WxMch.MchSecret
		}
		preq.Nonce_str = wxpay.GetNonceStr(32)
		preq.Out_trade_no = order.Paying_no
		trade_state = wxpay.WxSmfQuery2(preq, mch_secret)

		if strings.EqualFold(trade_state.Trade_state, "SUCCESS") {
			//支付成功后生成支付流水、修改订单支付状态、生成其他订单
			dbs.TxBegin()

			//数据库操作的逻辑
			dbs.TxCommit()
		} else if strings.EqualFold(trade_state.Trade_state, "NOTPAY") {
			order.Paying = 0
		} else if strings.EqualFold(trade_state.Trade_state, "PAYERROR") {
			order.Paying = 0
		}
	}
	type outParam struct {
		Paying    int
		Price_fee float64
		Payed_fee float64
	}
	pout := outParam{order.Paying, order.Price_fee, order.Payed_fee}
	c.WriteJSON(200, pout)
}

// 微信订单撤回
func HandlerJzOrderWxReverse(c *wego.WebContext) {
	dbs := worm.NewSession()

	type InParam struct {
		JzOrderId int64
	}

	var pin InParam
	err := c.ReadJSON(&pin)
	if err != nil {
		log.Error(err)
		c.AbortWithError(510, err)
		return
	}

	trade_state := wxpay.WxSmfReverseRet2{}
	if order.Paying > 0 {
		mch_secret := ""
		mch_pemid := ""
		var preq wxpay.WxSmfReverseReq2
		if len(shop.Wx_mchkey) > 0 {
			//普通商户模式支付
			preq.Appid = WxMch.SubAppid
			preq.Mch_id = shop.Wx_mchid
			mch_pemid = shop.Wx_mchid
			mch_secret = shop.Wx_mchkey
		} else {
			//服务商模式模式支付
			preq.Appid = WxMch.MchAppid
			preq.Mch_id = WxMch.MchId
			preq.Sub_mch_id = shop.Wx_mchid
			mch_secret = WxMch.MchSecret
		}
		preq.Nonce_str = wxpay.GetNonceStr(32)
		preq.Out_trade_no = order.Paying_no

		trade_state = wxpay.WxSmfReverse2(preq, mch_secret, mch_pemid)

		if strings.EqualFold(trade_state.Result_code, "SUCCESS") {
			//清理就诊订单支付状态
			dao.ClearJzddPaying(dbs, order.Id)
		}
	}
	c.WriteJSON(200, true)
}

// 微信退款
func HandlerGhOrderWxRefund(c *wego.WebContext) {
	dbs := worm.NewSession()

	type InParam struct {
		PayId int64
	}

	var pin InParam
	err := c.ReadJSON(&pin)
	if err != nil {
		log.Error(err)
		c.AbortWithError(510, err)
		return
	}
	var payent dao.PayStream
	_, err = dbs.Model(&payent).ID(pin.PayId).Get()
	if err != nil {
		log.Error(err)
		c.AbortWithError(510, err)
		return
	}

	mch_secret := ""
	mch_pemid := ""
	var preq wxpay.WxSmfRefundReq2

	if len(shop.Wx_mchkey) > 0 {
		//普通商户模式支付
		preq.Appid = WxMch.SubAppid
		preq.Mch_id = shop.Wx_mchid
		mch_pemid = shop.Wx_mchid
		mch_secret = shop.Wx_mchkey
	} else {
		//服务商模式模式支付
		preq.Appid = WxMch.MchAppid
		preq.Mch_id = WxMch.MchId
		preq.Sub_mch_id = shop.Wx_mchid
		mch_secret = WxMch.MchSecret
	}
	preq.Nonce_str = wxpay.GetNonceStr(32)
	preq.Out_trade_no = payent.Trade_no
	preq.Out_refund_no = payent.Trade_no + "_TK"
	preq.Total_fee = int(payent.Amount)
	preq.Refund_fee = int(payent.Amount)

	trade_state := wxpay.WxSmfRefund2(preq, mch_secret, mch_pemid)

	if trade_state.Result_code != "SUCCESS" {
		msg := "退款申请失败；"
		c.AbortWithText(510, msg)
		dbs.TxRollback()
		return
	}

	dbs.TxBegin()
	//数据库操作的逻辑
	dbs.TxCommit()
	c.WriteJSON(200, true)
}
