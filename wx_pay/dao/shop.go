package dao

type Shop struct {
	//int(11):id
	Id int64 `db:"id;autoincr"`
	//bigint(20):所属渠道
	Organ_id int64 `db:"organ_id"`
	//tinyint(4):功能类型：1药房 2医疗机构
	Category int `db:"category"`
	//varchar(32):机构级别
	Title string `db:"title"`
	//varchar(32):机构名称
	Name string `db:"name"`
	//varchar(22):机构简称
	Name_short string `db:"name_short"`
	//varchar(32):所在区域
	Area string `db:"area"`
	//varchar(32):联系人
	Contact string `db:"contact"`
	//varchar(32):联系电话
	Mobile string `db:"mobile"`
	//varchar(255):营业地址
	Address string `db:"address"`
	//varchar(6):营业执照
	Code_yyzz string `db:"code_yyzz"`
	//varchar(6):执业许可
	Code_zyxk string `db:"code_zyxk"`
	//varchar(6):邮编号码
	Code_yzbm string `db:"code_yzbm"`
	//decimal(11,5):省内运费
	Price_send_inner float64 `db:"price_send_inner"`
	//decimal(11,5):省外运费
	Price_send_outer float64 `db:"price_send_outer"`
	//varchar(512):京东物流配置
	Cfg_send_jd string `db:"cfg_send_jd"`
	//varchar(512):顺丰物流配置
	Cfg_send_sf string `db:"cfg_send_sf"`
	//varchar(32):物流公司
	Wuliu_name string `db:"wuliu_name"`
	//varchar(255):发货地址
	Send_address string `db:"send_address"`
	//varchar(256):备注
	Remark string `db:"remark"`
	//varchar(64):微信公众号
	Wx_app_name string `db:"wx_app_name"`
	//varchar(64):微信APPID
	Wx_appid string `db:"wx_appid"`
	//varchar(255):微信app密钥
	Wx_secret string `db:"wx_secret"`
	//varchar(64):微信商户号
	Wxpay_sub_mchid string `db:"wxpay_sub_mchid"`
	//tinyint(4):状态: 1: 正常 0: 暂停
	Status int `db:"status"`
	//bigint(20):创建时间
	Create_time int64 `db:"create_time;n_update"`
	//varchar(255):机构图片
	Logo_url string `db:"logo_url"`
	//bigint(20):机构介绍
	Content_id int64 `db:"content_id"`
	//varchar(2048):提成设置：诊金、咨询、治疗、中药处方 成药处方 产品处方 商品处方 治疗检查检验 营养处方
	Setting_usertag string `db:"setting_usertag"`
	//varchar(2048):文章类型：类型1,类型2
	Setting_wztype string `db:"setting_wztype"`
	//varchar(2048):商品类型：类型1,类型2
	Setting_sptype string `db:"setting_sptype"`
	//varchar(2048):商品标签：标签1,标签2
	Setting_sptag string `db:"setting_sptag"`
	//varchar(2048):收费设置
	Setting_shoufei string `db:"setting_shoufei"`
	//varchar(2048):发药设置
	Setting_fayao string `db:"setting_fayao"`
	//varchar(2048):处方配置
	Setting_cfset string `db:"setting_cfset"`
	//varchar(2048):药房配置;;；‘；；
	Setting_yfset string `db:"setting_yfset"`
	//varchar(64):客服电话
	Phone_kefu string `db:"phone_kefu"`
	//varchar(64):微信商户号
	Wx_mchid string `db:"wx_mchid"`
	//varchar(32):所在省份
	Area1 string `db:"area1"`
	//varchar(32):所在城市
	Area2 string `db:"area2"`
	//varchar(32):机构编码
	Code string `db:"code"`
	//varchar(255):app密钥
	Appkey string `db:"appkey"`
	//varchar(255):处方状态通知接口
	Callback_cf string `db:"callback_cf"`
	//varchar(16):积分发放平台编码
	Scoredk_code string `db:"scoredk_code"`
	//varchar(255):积分发放平台参数
	Scoredk_param string `db:"scoredk_param"`
	//varchar(64):微信支付商户号绑定的appid
	Wx_mchapp string `db:"wx_mchapp"`
	//varchar(64):微信支付商户号密钥
	Wx_mchkey string `db:"wx_mchkey"`
	//varchar(64):微信支付商户号证书名称
	Wx_mchpem string `db:"wx_mchpem"`
	//tinyint(4):云药房标志
	Flag_yyf int `db:"flag_yyf"`
	//tinyint(4):互联网医院标志
	Flag_net int `db:"flag_net"`
	//tinyint(4):医保标志
	Flag_yb int `db:"flag_yb"`
	//varchar(32):医保host
	Host_yb string `db:"host_yb"`
	//患者来源
	Setting_source string `db:"setting_source"`
}
