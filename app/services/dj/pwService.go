package dj

import (
	"dofun/config"
	"encoding/json"
	"log"
	"net/url"
)

const (
	PW_SERVICES_INFO = "pw:pw:service:list"
	REDIS_TIME       = 86400
)

type listParam struct {
	ClassId     string `json:"class_id"`
	AttributeId string `json:"attribute_id"`
	Sex         string `json:"sex"`
	OrderBy     string `json:"orderBy"`
	Page        int    `json:"page"`
	PageSize    int    `json:"pageSize"`
}

func PwList(classId string, attributeId string, sex string, orderBy string, page int, pageSize int) interface{} {
	requestParam := make(url.Values)
	param := listParam{
		ClassId:     classId,
		AttributeId: attributeId,
		Sex:         sex,
		OrderBy:     orderBy,
		Page:        page,
		PageSize:    pageSize,
	}
	jsonBytes, err := json.Marshal(param)
	if err != nil {

	}
	publicKey, _ := getPublicKey()
	encryptStr, _ := encrypt(publicKey, jsonBytes)
	requestParam["d"] = []string{string(encryptStr)}
	requestParam["business_id"] = []string{config.AppConfig.DfDjApiPublicBusinessId}

	reqUrl := config.AppConfig.DfDjDomainUrl + "/business/api/pwList.html"
	response, _ := djCurlToData("POST", reqUrl, requestParam.Encode())

	log.Print(publicKey)
	log.Print(response)
	return response.Data
}
