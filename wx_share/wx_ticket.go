package wx_share

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"
)

// ticket用于加强安全性，ticket的有效期目前为2个小时，需定时刷新。建议公众号开发者使用中控服务器统一获取和刷新ticket。
// 根据access_token获取微信公众号的ticket

// Ticket内容
type WxTicketRet struct {
	Ticket    string `json:"ticket"`
	ExpiresIn int64  `json:"expires_in"`
	ErrCode   int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
}

type TicketInfo struct {
	Ticket     string
	ExpireTime int64
}

// 由于ticket的获取次数有限，且部署于单点服务器，因此需要加锁来控制并发获取ticket
// 对ticket进行缓存

type TicketCache struct {
	tickets map[string]TicketInfo
	lock    sync.Mutex
}

var ticket_cache TicketCache

func GetTicket(appid string, token string) (string, error) {
	return ticket_cache.GetTicket(appid, token)
}

func (cache *TicketCache) GetTicket(appid, token string) (string, error) {
	cache.lock.Lock()
	defer cache.lock.Unlock()

	if cache.tickets == nil {
		cache.tickets = make(map[string]TicketInfo)
	}

	tm := time.Now().Unix()
	info, ok := cache.tickets[appid]
	//当从缓存中获取不到ticket 或 访问时间过了ticket的有效期， 则重新获取ticket 。 否则每次将从缓存中获取
	if ok == false || tm >= info.ExpireTime {
		ret, err := getJsapiTicket(token)
		if err != nil {
			return "", err
		}
		info.Ticket = ret.Ticket
		info.ExpireTime = tm + ret.ExpiresIn
		cache.tickets[appid] = info
		return info.Ticket, nil
	} else {
		return info.Ticket, nil
	}
}

// 获取小程序全局唯一后台接口调用凭据（ticket）。
// 开发者需要进行妥善保存。
//
func getJsapiTicket(token string) (WxTicketRet, error) {
	wx_api_url := "https://api.weixin.qq.com/cgi-bin/ticket/getticket"
	wx_api_url = fmt.Sprintf("%s?access_token=%s&type=jsapi",
		wx_api_url, token)

	var ent WxTicketRet
	res, err := WxApiGet(wx_api_url)
	if err != nil {
		return ent, err
	}

	if err := json.Unmarshal(res, &ent); err != nil {
		return ent, err
	}

	if ent.ErrCode != 0 {
		return ent, errors.New(ent.ErrMsg)
	}
	return ent, nil
}
