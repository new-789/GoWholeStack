package registry

import "time"

// Options 定义初始化参数类
type Options struct {
	// 地址
	Addrs []string
	// 超时
	Timeout time.Duration
	// 心跳时间
	HeartBeat int64
	// 注册地址
	// /a/b/c/xxx/10.xxx
	RegistryPath string
}

// Option 定义函数类型变量
type Option func(opts *Options)

// With 初始化函数系列

func WithAddrs(addrs []string)Option {
	return func(opts *Options) {
		opts.Addrs = addrs
	}
}

func WithTimeout(timeout time.Duration)Option {
	return func(opts *Options) {
		opts.Timeout = timeout
	}
}

func WithHeartBeat(heartBeat int64)Option {
	return func(opts *Options) {
		opts.HeartBeat = heartBeat
	}
}

func WithRegistryPath(path string)Option {
	return func(opts *Options) {
		opts.RegistryPath = path
	}
}