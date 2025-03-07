package core

import (
	"QuakeAPI/log"
	"QuakeAPI/model"
	"QuakeAPI/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type QuakeInterface interface {
	GetUserInfo(key string)
	GetServiceInfo(key string, query string, total int, pid string) (string, string)
}

type Core struct {
}

var httpClient utils.HttpClient

func init() {
	httpClient = utils.HttpClient{}
}

func (c Core) GetUserInfo(key string) {
	url := "https://quake.360.cn/api/v3/user/info"
	data := make(map[string]string)
	headers := make(map[string]string)
	headers["X-QuakeToken"] = key
	headers["Content-Type"] = "application/json"
	res := httpClient.DoGet(url, data, headers)
	var userInfo model.UserInfo
	err := json.Unmarshal(res, &userInfo)
	if err != nil {
		log.Log("unmarshal error:"+err.Error(), log.ERROR)
		return
	}
	if userInfo.Code != 0 {
		log.Log("Error API Key", log.ERROR)
		return
	}
	var roles bytes.Buffer
	for _, role := range userInfo.Data.Role {
		roles.WriteString(role.Fullname + "-")
	}
	println("[*] Connect to Quake success. Please wait...\n")
	println("[+] Query user info success:")
	println(" Role: " + roles.String())
	println(" Name: " + userInfo.Data.User.Username)
	println(" Mail: " + userInfo.Data.User.Email)
	println(" Phone: " + userInfo.Data.MobilePhone)
	println(" Credit: " + strconv.Itoa(userInfo.Data.Credit) + "\n")
}

func (c Core) GetServiceInfo(key string, query string, total int, pid string) (string, string) {
	url := "https://quake.360.cn/api/v3/scroll/quake_service"
	data := make(map[string]string)
	data["query"] = query
	data["size"] = strconv.Itoa(total)
	data["latest"] = "False"
	data["ignore_cache"] = "true"
	if pid != "" {
		data["pagination_id"] = pid
	}
	headers := make(map[string]string)
	headers["X-QuakeToken"] = key
	headers["Content-Type"] = "application/json"
	println("[*] Query: " + query + "\n[*] Wait...\n")
	res := httpClient.DoPost(url, data, headers)
	var serviceInfo model.ServiceInfo
	err := json.Unmarshal(res, &serviceInfo)
	if err != nil {
		println("unmarshal error:" + err.Error())
		return "", ""
	}
	if serviceInfo.Code != 0 {
		println("Error API Key")
		return "", ""
	}
	result := bytes.Buffer{}
	println("[+] Query service info success:")
	for _, value := range serviceInfo.Data {
		currentData := fmt.Sprintf("host:%s\tip:%s\tport:%d\tpath:%s\tcode:%d\ttitle:%s\t", value.Service.HTTP.Host, value.IP, value.Port, value.Service.HTTP.Path, value.Service.HTTP.StatusCode, strings.Trim(strings.Trim(strings.Trim(value.Service.HTTP.Title, "\n"), " "), "\n"))
		println(currentData)
		result.WriteString(currentData + "\n")
	}
	paginationID := serviceInfo.Meta.PaginationID
	return paginationID, result.String()
}
