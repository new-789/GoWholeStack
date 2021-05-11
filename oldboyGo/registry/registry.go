package registry

import "context"

/* 总接口
Name()：插件名，例如传 etcd
Init(opts ...Option)：初始化，里面是用选项设计模式做初始化。
Register()：做服务注册。
Unregister()：服务反注册，例如服务端停了，注册列表销毁。
GetServices：服务发现（ip port[] string）
*/

type Registry interface {
	// Name 插件名字
	Name() string
	// Init 初始化
	Init(ctx context.Context, opts ...Option) (err error)
	// Register 服务注册
	Register(ctx context.Context, service *Service) (err error)
	// Unregister 服务反注册
	Unregister(ctx context.Context, service *Service) (err error)
	// GetServices 服务发现
	GetServices(ctx context.Context, name string) (service *Service, err error)
}
