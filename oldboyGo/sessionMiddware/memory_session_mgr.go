package sessionMiddware

import (
	"errors"
	"github.com/satori/go.uuid"
	"sync"
)

type MemorySessionMgr struct {
	sessionMap map[string]Session
	rwlock sync.RWMutex
}

// NewMemorySessionMgr 构造函数
func NewMemorySessionMgr() SessionMgr {
	smr := &MemorySessionMgr{
		sessionMap: make(map[string]Session, 1024),
	}
	return smr
}

// Init 初始化函数
func (m *MemorySessionMgr)Init(addr string, option ...string) (err error) {
	return
}

// CreateSession 创建 session 对象
func (m *MemorySessionMgr)CreateSession() (session Session, err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	// 用 uuid 作为 session id
	uid := uuid.NewV4()
	sessionId := uid.String()
	// 创建一个 seesion 对象
	session = NewMemorySession(sessionId)
	// 加入到 session 管理者的大 map 中
	m.sessionMap[sessionId]=session
	return
}

func (m *MemorySessionMgr)GetSession(sessionID string) (session Session, err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	session, ok := m.sessionMap[sessionID]
	if !ok {
		err = errors.New("session not exists")
		return
	}
	return
}