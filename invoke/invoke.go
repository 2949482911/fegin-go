package invoke

import (
	"bytes"
	"fegin-go/config"
	"fegin-go/errors"
	"fegin-go/loadbanlace"
	"fegin-go/tool"
	"log"
)

// Invoke use the method invoke Remote Service,and return json string data
// @Param serviceName string need service name
func Invoke(serviceName string, method string, url string, data interface{}) string {
	ip, err := GetIp(serviceName)
	if ip == "" || err != nil {
		log.Printf("NO exist ip %s", serviceName)
		return ""
	}
	if method == "GET" {
		return tool.Get(GetHttpUrl(ip, url))
	} else if method == "POST" {
		return tool.Post(GetHttpUrl(ip, url), data, "application/json")
	}
	return ""
}

// GetIp get service ip
// @Param serviceName string need service name
// @Return string ip
// @Return error exception
func GetIp(serviceName string) (string, error) {
	var ip string
	switch config.FeginConfigurationInstance.LoadBalancing {
	case "RoundRobin":
		ip = loadbanlace.RoundRobin(serviceName)
		break
	case "Random":
		ip = loadbanlace.Random(serviceName)
		break
	case "WeightRoundRobin":
		ip = loadbanlace.WeightRoundRobin(serviceName)
		break
	case "WeightRandom":
		ip = loadbanlace.WeightRandom(serviceName)
		break
	}
	if ip != "" {
		return ip, nil
	}
	return "", errors.New("不存在的IP")
}

// GetHttpUrl get request url
func GetHttpUrl(host string, url string) string {
	var requestUrl bytes.Buffer
	requestUrl.WriteString("http://")
	requestUrl.WriteString(host)
	requestUrl.WriteString("/")
	requestUrl.WriteString(url)
	return requestUrl.String()
}
