package sessionMiddware

import "fmt"

// 总初始化文件：中间件应该让用户选择使用那个版本

var (
	sessionMgr SessionMgr
)

func Init(provider string, addr string, options ...string) (err error) {
	switch provider {
	case "memory":
		sessionMgr = NewMemorySessionMgr()
	case "redis":
		sessionMgr = NewRedisSessionMgr()
	default:
		fmt.Errorf("不支持该类型的 session 操作，本中间件仅支持memory/redis版本操作")
		return
	}
	err = sessionMgr.Init(addr, options...)
	return
}