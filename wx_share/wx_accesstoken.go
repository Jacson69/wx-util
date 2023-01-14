package wx_share

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"
)

//公众号调用各接口时都需使用access_token。开发者需要进行妥善保存。access_token的存储至少要保留512个字符空间。
//access_token的有效期目前为2个小时，需定时刷新，重复获取将导致上次获取的access_token失效。
//建议公众号开发者使用中控服务器统一获取和刷新access_token
//其他业务逻辑服务器所使用的access_token均来自于该中控服务器，不应该各自去刷新，否则容易造成冲突，导致access_token覆盖而影响业务。

// token内容
type WxApiTokenRet struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
}

type TokenInfo struct {
	Token      string
	ExpireTime int64
}

// 由于access_token的获取次数有限，且部署于单点服务器，因此需要加锁来控制并发获取ticket
// 对access_token进行缓存，且运用锁机制来控制并发请求
type TokenCache struct {
	tokens map[string]TokenInfo
	lock   sync.Mutex
}

var token_cache TokenCache

func GetAccessToken(appid string, secret string) (string, error) {
	return token_cache.GetAccessToken(appid, secret)
}

func (cache *TokenCache) GetAccessToken(appid string, secret string) (string, error) {
	cache.lock.Lock()
	defer cache.lock.Unlock()

	if cache.tokens == nil {
		cache.tokens = make(map[string]TokenInfo)
	}

	tm := time.Now().Unix()
	info, ok := cache.tokens[appid]
	//当从缓存中获取不到 access_token 或 访问时间过了 access_token 的有效期， 则重新获取access_token 。 否则每次将从缓存中获取
	if ok == false || tm >= info.ExpireTime {
		ret, err := getAccessToken(appid, secret)
		if err != nil {
			return "", err
		}
		info.Token = ret.AccessToken
		info.ExpireTime = tm + ret.ExpiresIn
		cache.tokens[appid] = info
		return info.Token, nil
	} else {
		return info.Token, nil
	}
}

// 获取小程序全局唯一后台接口调用凭据（access_token）。
// 调用绝大多数后台接口时都需使用 access_token，开发者需要进行妥善保存。
func getAccessToken(appid string, secret string) (WxApiTokenRet, error) {
	wx_api_url := "https://api.weixin.qq.com/cgi-bin/token"
	wx_api_url = fmt.Sprintf("%s?grant_type=client_credential&appid=%s&secret=%s",
		wx_api_url, appid, secret)

	var ent WxApiTokenRet
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
