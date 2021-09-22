package main

import (
	"fegin-go/common"
	"fegin-go/config"
	"fegin-go/invoke"
	"fmt"
)

// main app start
func main() {
	config.FeginConfigurationInstance = &config.FeginConfiguration{LoadBalancing: "RoundRobin"}
	str := invoke.Invoke("register", common.POST, "", nil)
	fmt.Println(str)
}
