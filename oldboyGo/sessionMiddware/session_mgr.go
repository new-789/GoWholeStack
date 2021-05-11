package sessionMiddware

// SessionMgr 定义管理者，算理所有 session
type SessionMgr interface {
	// Init 初始化
	Init(addr string, options ...string) (err error)
	// CreateSession 创建 Session ，返回 session 对象和错误信息
	CreateSession() (session Session, err error)
	// GetSession 获取Session ，返回 session 对象和错误信息
	GetSession(sessionId string) (session Session, err error)
}