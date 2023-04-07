package dao

type PayStream struct {
	//bigint(20):自增id
	Id int64 `db:"id;autoincr"`
	//bigint(20):机构id
	Shop_id int64 `db:"shop_id"`
	//varchar(128):支付方式：xj,wx,zfb,yhk,ybk,hyk,bpt,wpt
	Pay_mode string `db:"pay_mode"`
	//tinyint(4):订单类型：1预约挂号 2处方 3商品 4会员
	Order_type int `db:"order_type"`
	//varchar(64):订单编号
	Order_no string `db:"order_no"`
	//varchar(32):操作人员
	Oper_name string `db:"oper_name"`
	//varchar(32):客户姓名
	User_name string `db:"user_name"`
	//varchar(24):客户电话号码
	User_mobile string `db:"user_mobile"`
	//int(11):交易金额, 单位分
	Amount int64 `db:"amount"`
	//bigint(20):会员id
	Huiyuan_id int64 `db:"huiyuan_id"`
	//int(11):奖励积分
	Award_score int64 `db:"award_score"`
	//varchar(64):商户号
	Mch_id string `db:"mch_id"`
	//varchar(64):用户ID
	Mch_user string `db:"mch_user"`
	//varchar(128):商户交易号
	Trade_no_mch string `db:"trade_no_mch"`
	//varchar(128):交易号
	Trade_no string `db:"trade_no"`
	//bigint(20):交易时间
	Trade_time int64 `db:"trade_time"`
	//varchar(64):返回码
	Retcode string `db:"retcode"`
	//varchar(1024):描述
	Retmsg string `db:"retmsg"`
	//tinyint(4):状态: 0:未支付 1 已支付
	Status int `db:"status"`
	//bigint(20):描述
	Create_time int64 `db:"create_time;n_update"`
	//bigint(20):退款金额
	Refunded int64 `db:"refunded"`
	//bigint(20):当前退款金额
	Refunding int64 `db:"refunding"`
	//varchar(64):当前退款编码
	Refund_no string `db:"refund_no"`
}
