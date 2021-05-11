package sessionMiddware

import (
	"errors"
	"sync"
)

// MemorySession 定义结构体对象
type MemorySession struct {
	SessionId string
	// 存 k-v map
	Data map[string]interface{}
	// 读写锁
	RwLock sync.RWMutex
}

// NewMemorySession 构造函数
func NewMemorySession(id string) *MemorySession {
	s := &MemorySession{
		SessionId: id,
		Data: make(map[string]interface{}, 16),
	}
	return s
}

// Set 操作 session Set 方法具体实现
func (m *MemorySession)Set(key string, value interface{}) (err error) {
	// 加锁
	m.RwLock.Lock()
	defer m.RwLock.Unlock()
	m.Data[key] = value
	return
}

// Get 操作 session Get 方法具体实现
func (m *MemorySession)Get(key string) (value interface{}, err error) {
	m.RwLock.Lock()
	defer m.RwLock.Unlock()
	value, ok := m.Data[key]
	if !ok {
		err = errors.New("key not exits in session")
		return
	}
	return
}

// Del 操作 session Del 方法具体实现
func (m *MemorySession)Del(key string) (err error) {
	m.RwLock.Lock()
	defer m.RwLock.Unlock()
	delete(m.Data, key)
	return
}

func (m *MemorySession)Save() (err error) {
	return
}

