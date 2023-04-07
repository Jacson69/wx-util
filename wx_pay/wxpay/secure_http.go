package wxpay

import (
	"bytes"
	"crypto/tls"
	"fmt"
	log "github.com/haming123/wego/dlog"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

var _tlsConfig *tls.Config

// 采用单例模式初始化ca
// ca证书的位置，采用固定位置：./cert/cert.pem
func getTLSConfig(pem_name string) (*tls.Config, error) {

	if _tlsConfig != nil {
		return _tlsConfig, nil
	}

	//ca证书的位置，需要绝对路径
	conf_dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	wechatCertPath := fmt.Sprintf("%s/cert/cert_1640064324.pem", conf_dir)
	wechatKeyPath := fmt.Sprintf("%s/cert/key_1640064324.pem", conf_dir)
	if len(pem_name) > 0 {
		wechatCertPath = fmt.Sprintf("%s/cert/cert_%s.pem", conf_dir, pem_name)
		wechatKeyPath = fmt.Sprintf("%s/cert/key_%s.pem", conf_dir, pem_name)
	}
	//gowf.LogDf("cert_file=%s\n", wechatCertPath)
	log.Info(wechatCertPath)

	certPEMBlock, err := ioutil.ReadFile(wechatCertPath)
	if err != nil {
		return nil, err
	}
	keyPEMBlock, err := ioutil.ReadFile(wechatKeyPath)
	if err != nil {
		return nil, err
	}

	// load cert
	//cert, err := tls.LoadX509KeyPair(wechatCertPath, wechatKeyPath)
	cert, err := tls.X509KeyPair(certPEMBlock, keyPEMBlock)
	if err != nil {
		return nil, err
	}

	_tlsConfig = &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      nil,
	}

	return _tlsConfig, nil
}

// 携带ca证书的安全请求
func SecurePost(url string, xmlContent []byte, pem_name string) (*http.Response, error) {
	tlsConfig, err := getTLSConfig(pem_name)
	if err != nil {
		return nil, err
	}

	tr := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: tr}
	return client.Post(url, "application/xml", bytes.NewBuffer(xmlContent))
}
