package dj

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"dofun/config"
	"dofun/pkg/errno"
	"dofun/pkg/gredis"
	"encoding/json"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type responseBody struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	djPublicKey = "pw:other:publickey"
)

// 加密
func encrypt(publicKey []byte, data []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, data)
}

// 解密
func decrypt(privateKey []byte, ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

//获取公钥
func getPublicKey() ([]byte, error) {
	publicKey, err := gredis.Remember(djPublicKey, 10800, func() interface{} {
		reqUrl := config.AppConfig.DfDjDomainUrl + "/business/api/getKey.html"
		rep, err := djCurlToData(http.MethodGet, reqUrl+"?business_id="+config.AppConfig.DfDjApiPublicBusinessId, "")

		if err != nil {
			return nil
		}
		if value, ok := rep.Data.(map[string]interface{}); ok {
			if pkey, ok := value["publicKey"].(string); ok {
				return pkey
			}
		}
		return nil
	})
	if err != nil {
		return publicKey, err
	}

	return publicKey, nil
}

//发起请求
func djCurlToData(method string, reqUrl string, param string) (*responseBody, error) {
	client := &http.Client{}

	req, err := http.NewRequest(method, reqUrl, strings.NewReader(param))
	if err != nil {
		// handle error
		return nil, errno.Base(errno.InternalServerError, "请求参数错误")
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")

	resp, _ := client.Do(req)
	if resp == nil {
		return nil, errno.Base(errno.InternalServerError, "请求参数错误")
	}
	defer resp.Body.Close()

	jsonBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return nil, errno.Base(errno.InternalServerError, "数据异常")
	}
	var result responseBody
	if err := json.Unmarshal(jsonBody, &result); err != nil {
		return nil, errno.Base(errno.InternalServerError, "数据异常")
	}
	return &result, nil
}
