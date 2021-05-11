package conf

type AppConfig struct {
	KafkaConfig `ini:"kafka"`
	//TailLogConfig `ini:"taillog"`
	EtcdConf `ini:"etcd"`
}

type KafkaConfig struct {
	Address string `ini:"address"`
	//Topic string `ini:"topic"`
	ChanMaxSize int `ini:"chan_max_size"`
}

type EtcdConf struct {
	Address string `ini:"address"`
	Key string `ini:"collect_log_key"`
	Timeout int `ini:"timeout"`
}


/*
type TailLogConfig struct {
	Filename string `ini:"path"`
}
 */