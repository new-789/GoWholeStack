package registry

import (
	"context"
	"fmt"
	"sync"
)

/* 插件管理类
可以用一个大 map 管理，key 字符串，value是 Registry 接口对象
用户自定义去调用，自定义插件。
实现注册中心的初始化，供系统使用。
*/

// PluginMgr 声明管理者结构体
type PluginMgr struct {
	// map 维护所有的插件
	plugins map[string]Registry
	lock sync.Mutex
}

var (
	pluginMgr = &PluginMgr{
		plugins: make(map[string]Registry),
	}
)

// RegisterPlugin 插件注册，提供给外部注册服务方法
func RegisterPlugin(registry Registry) (err error) {
	return pluginMgr.registerPlugin(registry)
}

// 注册插件, 服务内部注册服务实现使用
func (p *PluginMgr)registerPlugin(plugin Registry) (err error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	// 先判断大 map 中需要注册的服务在 map 中是否以存在
	if _, ok := p.plugins[plugin.Name()]; ok {
		err = fmt.Errorf("registry plugin exists")
		return
	}
	p.plugins[plugin.Name()] = plugin
	return
}

// InitRegistry 进行初始化注册中心，提供给外部调用的方法
func InitRegistry(ctx context.Context, name string, opts ...Option) (registry Registry, err error) {
	return pluginMgr.initRegistry(ctx, name, opts...)
}

// 注册插件后台内部具体实现
func (p *PluginMgr)initRegistry(ctx context.Context, name string, opts ...Option) (registry Registry, err error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	// 先查看判断 map 中是否存在需要注册的插件列表，若不存在则报错
	plugin, ok := p.plugins[name]
	if !ok {
		err = fmt.Errorf("plugin %s not exists", name)
		return
	}
	// 存在，返回值赋值
	registry = plugin
	// 进行组件初始化
	err = plugin.Init(ctx, opts...)
	return
}