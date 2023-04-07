package wxpay

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/xml"
	"math/rand"
	"net"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Params map[string]string

func (p Params) SetString(k, s string) Params {
	p[k] = s
	return p
}
func (p Params) GetString(k string) string {
	s, _ := p[k]
	return s
}
func (p Params) SetInt64(k string, i int64) Params {
	p[k] = strconv.FormatInt(i, 10)
	return p
}
func (p Params) GetInt64(k string) int64 {
	i, _ := strconv.ParseInt(p.GetString(k), 10, 64)
	return i
}
func (p Params) ContainsKey(key string) bool {
	_, ok := p[key]
	return ok
}

func XmlToMap(xmlStr string) Params {

	params := make(Params)
	decoder := xml.NewDecoder(strings.NewReader(xmlStr))

	var key string
	var value string
	for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		case xml.StartElement: // 开始标签
			key = token.Name.Local
		case xml.CharData: // 标签内容
			content := string([]byte(token))
			value = content
		}
		if key != "xml" {
			if value != "\n" {
				params.SetString(key, value)
			}
		}
	}

	return params
}

func MapToXml(params Params) string {
	var buf bytes.Buffer
	buf.WriteString(`<xml>`)
	for k, v := range params {
		buf.WriteString(`<`)
		buf.WriteString(k)
		buf.WriteString(`><![CDATA[`)
		buf.WriteString(v)
		buf.WriteString(`]]></`)
		buf.WriteString(k)
		buf.WriteString(`>`)
	}
	buf.WriteString(`</xml>`)

	return buf.String()
}

func SignMD5(signstr string) string {
	hasher := md5.New()
	hasher.Write([]byte(signstr))
	sign := hex.EncodeToString(hasher.Sum(nil))
	return sign
}

func SignSHA1(signstr string) string {
	hasher := sha1.New()
	hasher.Write([]byte(signstr))
	sign := hex.EncodeToString(hasher.Sum(nil))
	return sign
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func GetNonceStr(n int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}

func GetRemoteIp(req *http.Request) string {
	remoteAddr := req.Header.Get("X-Forwarded-For")
	if len(remoteAddr) > 0 {
		return remoteAddr
	}

	remoteAddr = req.RemoteAddr
	if ip := req.Header.Get("Remote_addr"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}
	return remoteAddr
}

// 微信支付计算签名的函数
func CalWxPaySign2(params Params, key string) string {
	var keys = make([]string, 0, len(params))
	for k := range params {
		if k != "sign" { // 排除sign字段
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)

	var buf bytes.Buffer
	for _, k := range keys {
		if len(params.GetString(k)) > 0 {
			buf.WriteString(k)
			buf.WriteString(`=`)
			buf.WriteString(params.GetString(k))
			buf.WriteString(`&`)
		}
	}
	buf.WriteString(`key=`)
	buf.WriteString(key)

	var dataMd5 [16]byte
	dataMd5 = md5.Sum(buf.Bytes())
	str := hex.EncodeToString(dataMd5[:])

	return strings.ToUpper(str)
}

// 微信数据解密:去除填充
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// 微信数据解密
func AesCBCDncrypt(encryptData, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	blockSize := block.BlockSize()
	if len(encryptData) < blockSize {
		panic("ciphertext too short")
	}
	if len(encryptData)%blockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(encryptData, encryptData)

	// 解填充
	encryptData = PKCS7UnPadding(encryptData)

	return encryptData, nil
}

func WxDncrypt(rawData, key, iv string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(rawData)
	key_b, err_1 := base64.StdEncoding.DecodeString(key)
	iv_b, _ := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return "", err
	}
	if err_1 != nil {
		return "", err_1
	}
	dnData, err := AesCBCDncrypt(data, key_b, iv_b)
	if err != nil {
		return "", err
	}
	return string(dnData), nil
}

// 微信接口调用状态
type WXApiStatus struct {
	Return_code string `xml:"return_code"`
	Return_msg  string `xml:"return_msg"`
}

func (ent *WXApiStatus) GetSimpleXml() string {
	var buffer bytes.Buffer

	buffer.WriteString("<xml>\n")
	buffer.WriteString("<return_code>")
	buffer.WriteString(ent.Return_code)
	buffer.WriteString("</return_code>\n")

	buffer.WriteString("<return_msg>")
	buffer.WriteString(ent.Return_msg)
	buffer.WriteString("</return_msg>\n")
	buffer.WriteString("</xml>\n")

	return buffer.String()
}
