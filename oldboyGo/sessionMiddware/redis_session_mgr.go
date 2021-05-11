package sessionMiddware

import (
	"errors"
	"github.com/garyburd/redigo/redis"
	uuid "github.com/satori/go.uuid"
	"sync"
	"time"
)

// redisSessionMgr redisSession 管理者

// RedisSessionMgr redisSession 管理者
type RedisSessionMgr struct {
	// redis 地址
	addr string
	// redis 密码
	password string
	// redis 连接池
	pool *redis.Pool
	// 锁
	rwlock sync.RWMutex
	// map
	sessionMap map[string]Session
}

// NewRedisSessionMgr RedisSession 管理者构造函数,返回一个 SessionMgr 管理者
func NewRedisSessionMgr() SessionMgr {
	sr := &RedisSessionMgr{
		sessionMap: make(map[string]Session, 32),
	}
	return sr
}

// Init 初始化
func (r *RedisSessionMgr)Init(addr string, option ...string) (err error) {
	// 若有其他参数
	if len(option) >0 {
		r.password = option[0]
	}
	// 创建连接池
	r.pool = myPool(addr, r.password)
	r.addr = addr
	return
}

func myPool(addr, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle: 64,
		MaxActive: 1000,
		IdleTimeout: time.Second * 240,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}
			// 如果有密码，则判断密码
			if _, err := conn.Do("AUTH", password); err != nil {
				conn.Close()
				return nil, err
			}
			return conn, err
		},
		// redis 连接测试，上线时应注释掉
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func (r *RedisSessionMgr)CreateSession() (session Session, err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	id := uuid.NewV4()
	// 用 uuid 作为 sessionId
	sessionId := id.String()
	session = NewRedisSession(sessionId, r.pool)
	return
}

func (r *RedisSessionMgr)GetSession(sessionID string) (session Session, err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	session, ok := r.sessionMap[sessionID]
	if !ok {
		err = errors.New("session not exists")
		return
	}
	return
}
