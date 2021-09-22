package config

// FeginConfiguration 配置类
type FeginConfiguration struct {
	LoadBalancing string `json:"load_balancing" yaml:"load-balancing"` // 随机算法
}

var FeginConfigurationInstance *FeginConfiguration
