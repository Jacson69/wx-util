package dao

import (
	"fmt"
	"github.com/haming123/wego/worm"
	"time"
)

type JzOrder struct {
	//bigint(20):自增id
	Id int64 `db:"id;autoincr"`
	//bigint(20):机构ID
	Shop_id int64 `db:"shop_id"`
	//varchar(64):订单号
	Order_no string `db:"order_no"`
	//bigint(20):排期id
	Paiqi_id int64 `db:"paiqi_id"`
	//tinyint(4):订单类型: 0 预约挂号 1 门诊挂号 2快速接诊 3患者购药
	Jz_type int `db:"jz_type"`
	//varchar(32):患者来源
	Source string `db:"source"`
	//tinyint(4):开方设备: 0 PC 1 手机
	Device int `db:"device"`
	//varchar(64):就诊日期
	Treat_date string `db:"treat_date"`
	//varchar(64):就诊时段
	Treat_time string `db:"treat_time"`
	//varchar(255):就诊地址
	Treat_addr string `db:"treat_addr"`
	//tinyint(4):0: 首诊 1: 复诊
	Flag_first int `db:"flag_first"`
	//bigint(20):医生ID
	Doctor_id int64 `db:"doctor_id"`
	//varchar(32):医生姓名
	Doctor_name string `db:"doctor_name"`
	//bigint(20):助理id
	Help_id int64 `db:"help_id"`
	//varchar(32):助理姓名
	Help_name string `db:"help_name"`
	//bigint(20):客户id
	User_id int64 `db:"user_id"`
	//varchar(32):客户姓名
	User_name string `db:"user_name"`
	//varchar(24):客户电话号码
	User_mobile string `db:"user_mobile"`
	//varchar(32):用户身份证
	Patient_idno string `db:"patient_idno"`
	//varchar(32):患者姓名
	Patient_name string `db:"patient_name"`
	//smallint(6):患者年龄
	Patient_age int `db:"patient_age"`
	//smallint(6):患者月份
	Patient_month int `db:"patient_month"`
	//tinyint(4):患者性别
	Patient_sex int `db:"patient_sex"`
	//varchar(32):出生日期
	Patient_birth string `db:"patient_birth"`
	//bigint(20):患者体重
	Patient_weight int `db:"patient_weight"`
	//varchar(64):诊断结果
	Diagnose string `db:"diagnose"`
	//varchar(1024):患者病历
	Bingli_info string `db:"bingli_info"`
	//varchar(4096):病历图片
	Bingli_img string `db:"bingli_img"`
	//tinyint(4):处方格式:0电子处方 1 照片处方
	Flag_cfgs int `db:"flag_cfgs"`
	//tinyint(4):挂起标志
	Flag_wait int `db:"flag_wait"`
	//tinyint(4):客服确认
	Flag_kefu int `db:"flag_kefu"`
	//varchar(32):客服名称
	Kefu_name string `db:"kefu_name"`
	//varchar(1024):客服备注
	Kefu_desc string `db:"kefu_desc"`
	//tinyint(4):收货方式  0自提 1邮寄
	Recive_type int `db:"recive_type"`
	//varchar(32):收货联系人
	Recive_name string `db:"recive_name"`
	//varchar(24):收货电话
	Recive_mobile string `db:"recive_mobile"`
	//varchar(1024):收货地址名称
	Recive_addr string `db:"recive_addr"`
	//decimal(11,2):预约挂号金额
	Price_gh float64 `db:"price_gh"`
	//decimal(11,2):挂号已收金额
	Payed_gh float64 `db:"payed_gh"`
	//decimal(11,2):就诊诊金金额
	Price_yszj float64 `db:"price_yszj"`
	//decimal(11,2):就诊处方金额
	Price_cfang float64 `db:"price_cfang"`
	//decimal(11,2):就诊运费金额
	Price_send float64 `db:"price_send"`
	//decimal(11,2):就诊收费金额
	Price_fee float64 `db:"price_fee"`
	//decimal(11,2):就诊已收金额
	Payed_fee float64 `db:"payed_fee"`
	//decimal(11,2):挂号诊金金额
	Price_treat float64 `db:"price_treat"`
	//decimal(11,2):缴费总额
	Price float64 `db:"price"`
	//varchar(255):备注
	Remark string `db:"remark"`
	//tinyint(4):状态 0 取消 1:预约待支付 2待登记 3:挂号待支付 4待接诊 5就诊中 6待计价 7就诊待支付 99完成
	Status int `db:"status"`
	//bigint(20):当前支付金额
	Paying int `db:"paying"`
	//varchar(24):当前支付方式
	Paying_mode string `db:"paying_mode"`
	//varchar(64):当前支付编码
	Paying_no string `db:"paying_no"`
	//bigint(20):创建时间
	Create_time int64 `db:"create_time"`
	//bigint(20):挂号时间
	Time_gh int64 `db:"time_gh"`
	//bigint(20):支付时间
	Time_pay int64 `db:"time_pay"`
	//bigint(20):约待支付时间
	Time_yypay int64 `db:"time_yypay"`
	//bigint(20):挂号支付时间
	Time_ghpay int64 `db:"time_ghpay"`
	//bigint(20):处方支付时间
	Time_cfpay int64 `db:"time_cfpay"`
	//varchar(64):叫号内容
	Voice_text string `db:"voice_text"`
	//bigint(20):叫号时间
	Voice_time int64 `db:"voice_time"`
	//bigint(20):患者id
	Patient_id int64 `db:"patient_id"`
	//tinyint(4):互联网医院标志
	Flag_net int `db:"flag_net"`
	//tinyint(4):数据状态：0未同步 1同步 <0错误
	Flag_sync int `db:"flag_sync"`
	//varchar(64):同步信息
	Info_sync string `db:"info_sync"`
	//bigint(20):开方时间
	Time_kf int64 `db:"time_kf"`
	//varchar(128):中医诊断名称
	Jbname_zy string `db:"jbname_zy"`
	//varchar(128):西医诊断名称
	Jbname_xy string `db:"jbname_xy"`
	//varchar(64):中医诊断编码
	Jbcode_zy string `db:"jbcode_zy"`
	//varchar(64):西医诊断编码
	Jbcode_xy string `db:"jbcode_xy"`
	//tinyint(4):医保上传状态：0未同步 1同步 <0错误
	Yb_sync int `db:"yb_sync"`
	//tinyint(4):费别 0:自费 1:医保
	Pay_type int `db:"pay_type"`
	//varchar(64):社保卡号
	Yb_card string `db:"yb_card"`
	//varchar(32):社保卡电脑号
	Yb_pcno string `db:"yb_pcno"`
	//varchar(32):医保门诊号
	Yb_mzno string `db:"yb_mzno"`
	//varchar(32):医保结算序列号
	Yb_seq_no string `db:"yb_seq_no"`
	//varchar(32):社保人员编号
	Yb_psnno string `db:"yb_psnno"`
	//varchar(1024):医保数据
	Yb_info string `db:"yb_info"`
	//varchar(32):医保收费批次号
	Yb_payno string `db:"yb_payno"`
	//tinyint(4):医保上传状态：0未同步 1同步
	Yb_sync2 int `db:"yb_sync2"`
	//decimal(11,2):附加收费金额
	Price_ext float64 `db:"price_ext"`
	//tinyint(4):处方费别 0:自费 1:医保
	Pay_type2 int `db:"pay_type2"`
	//tinyint(4):赊账标志
	Flag_credit int `db:"flag_credit"`
	//varchar(1024):赊账备注
	Desc_credit string `db:"desc_credit"`
	//bigint(20):就诊提醒时间
	Time_ghtx int64 `db:"time_ghtx"`
	//tinyint(4):咨询订单标志
	Flag_zixun int `db:"flag_zixun"`
	//varchar(1023):症状描述
	Symptom_desc string `db:"symptom_desc"`
	//varchar(4096):症状图片
	Symptom_img string `db:"symptom_img"`
}

func GenPayOrderTradeNo(ddtype string, ddid int64) string {
	tm := time.Now().Unix()
	ddid = ddid + 11365295
	return fmt.Sprintf("%s%d%08d", ddtype, ddid, tm)
}

// 设置就诊订单支付中状态
func SetJzddPaying(dbs *worm.DbSession, jzid int64, paying int, paying_mode string, paying_no string) error {
	_, err := dbs.SQL("update jzorder set paying=?, paying_mode=?, paying_no=? where id=?",
		paying, paying_mode, paying_no, jzid).Exec()
	return err
}

// 清除就诊订单支付中状态
func ClearJzddPaying(dbs *worm.DbSession, ddid int64) error {
	_, err := dbs.SQL("update jzorder set paying=0, paying_mode='', paying_no='' where id=?", ddid).Exec()
	return err
}
