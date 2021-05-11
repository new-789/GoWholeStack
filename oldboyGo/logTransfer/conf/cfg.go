package conf

// LogTransferCfg 全局配置信息结构体
type LogTransferCfg struct {
	KafkaCfg `ini:"kafka"`
	ESCfg `ini:"es"`
}

// KafkaCfg kafka 配置信息
type KafkaCfg struct {
	Address string `ini:"address"`
	Topic string `ini:"topic"`
}

// ESCfg ElasticSearch 配置信息
type ESCfg struct {
	Address string `ini:"address"`
	ChanSize int `ini:"chan_size"`
	MaxGoroutine int `ini:"max_goroutine"`
}


