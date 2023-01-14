package wx_share

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	mathRand "math/rand"
	"net/http"
)

// 生成指定长度的字符串
func RandStringBytes(n int) string {
	const letterBytes = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[mathRand.Intn(len(letterBytes))]
	}
	return string(b)
}

// SHA1加密
func GetSha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

// 对微信服务平台发起get请求
func WxApiGet(wx_addr string) ([]byte, error) {
	//wx_addr 代表url地址
	res, err := http.Get(wx_addr)
	if err != nil {
		return nil, err
	}

	raw, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("http statusCode=%v", res.StatusCode))
	}

	return raw, nil
}
