package sessionMiddware

// Session 定义 session 操作规范
type Session interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
	Del(key string) error
	Save() error
}
