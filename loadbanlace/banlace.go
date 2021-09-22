package loadbanlace

import (
	"bytes"
	"math/rand"
	"strconv"
	"sync"
)

// Server save local host name and ip
// use service name as key , register center Obtain service info as value
var Server map[string][]map[string]string = make(map[string][]map[string]string, 0)

func init() {
	reginster := make([]map[string]string, 0)
	reginster = append(reginster, map[string]string{
		"ip":       "www.baidu.com",
		"port":     "80",
		"weighted": "5",
	})
	Server["register"] = reginster
}

// pos
var pos int = 0

// RoundRobin get service use Round robin algorithm
// @Param serviceName string service name
// @Return string service ip
func RoundRobin(serviceName string) string {
	serverList := Server[serviceName]
	length := len(serverList)
	mutex := sync.Mutex{}
	var serverIp bytes.Buffer
	mutex.Lock()
	if pos > length {
		pos = 0
	} else {
		serverIp.WriteString(serverList[pos]["ip"])
		serverIp.WriteString(":")
		serverIp.WriteString(serverList[pos]["port"])
		pos++
	}
	defer mutex.Unlock()
	return serverIp.String()
}

// Random get service use Random algorithm
// @Param serviceName string service name
// @Return string service ip
func Random(serviceName string) string {
	serverList := Server[serviceName]
	length := len(serverList)
	mutex := sync.Mutex{}
	var serverIp bytes.Buffer
	mutex.Lock()
	serverIp.WriteString(serverList[rand.Intn(length)]["ip"])
	serverIp.WriteString(":")
	serverIp.WriteString(serverList[pos]["port"])
	defer mutex.Unlock()
	return serverIp.String()
}

// weightRoundRobinPos Weighted polling number
var weightRoundRobinPos int = 0

// WeightRoundRobin Weighted polling get service ip
// @Param serviceName string service name
// @Return string service ip
func WeightRoundRobin(serviceName string) string {
	serverList := Server[serviceName]
	mutex := sync.Mutex{}
	var serverIp bytes.Buffer
	var weightedList = make([]string, 0)
	for _, v := range serverList {
		var weight, _ = strconv.Atoi(v["weighted"])
		for i := 0; i < weight; i++ {
			weightedList = append(weightedList, v["ip"])
		}
	}
	mutex.Lock()
	defer mutex.Unlock()
	length := len(weightedList)
	if weightRoundRobinPos > length {
		pos = 0
	} else {
		serverIp.WriteString(serverList[rand.Intn(length)]["ip"])
		serverIp.WriteString(":")
		serverIp.WriteString(serverList[pos]["port"])
		weightRoundRobinPos++
	}
	return serverIp.String()
}

// WeightRandom Random selection of weights get service ip
// @Param serviceName string service name
// @Return string service ip
func WeightRandom(serviceName string) string {
	serverList := Server[serviceName]
	mutex := sync.Mutex{}
	var serverIp bytes.Buffer
	var weightedList = make([]string, 0)
	for _, v := range serverList {
		var weight, _ = strconv.Atoi(v["weighted"])
		for i := 0; i < weight; i++ {
			weightedList = append(weightedList, v["ip"])
		}
	}
	mutex.Lock()

	defer mutex.Unlock()
	serverIp.WriteString(serverList[rand.Intn(len(weightedList))]["ip"])
	serverIp.WriteString(":")
	serverIp.WriteString(serverList[pos]["port"])
	return serverIp.String()
}
