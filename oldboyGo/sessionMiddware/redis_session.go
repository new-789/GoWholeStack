package sessionMiddware

import (
	"encoding/json"
	"errors"
	"github.com/garyburd/redigo/redis"
	"sync"
)

// redisSession Demo 方法实现者

// RedisSession 定义 redisSession 对象
type RedisSession struct {
	sessionId string
	pool *redis.Pool
	// 用来将设置的 session 暂放在内存的 map 中，批量导入 redis 提升性能
	sessionMap map[string]interface{}
	// 读写锁
	rwlock sync.RWMutex
	// 记录内存中 map 是否被操作
	flag int
}

// 用常量定义状态
const (
	// SessionFlagNone 表示内存中的 map 没发生变化
	SessionFlagNone = iota
	// SessionFlagModify 表示内存中的 map 以发生变化
	SessionFlagModify
)

// NewRedisSession 构造函数
func NewRedisSession(id string, pool *redis.Pool) *RedisSession {
	s := &RedisSession{
		sessionId: id,
		sessionMap: make(map[string]interface{}, 16),
		pool: pool,
		flag: SessionFlagNone,
	}
	return s
}

// Set session 存储到内存中的 map
func (r *RedisSession)Set(key string, value interface{}) (err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	r.sessionMap[key] = value
	// 更改标记
	r.flag = SessionFlagModify
	return
}

// Save 将内存中的 map 存入
func (r *RedisSession)Save() (err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	// 若数据没变则无需存储
	if r.flag != SessionFlagModify {
		return
	}
	// 内存中的 sessionMap 进行序列化
	data, err := json.Marshal(r.sessionMap)
	if err != nil {
		return
	}
	// 获取 redis 连接
	conn := r.pool.Get()
	_, err = conn.Do("SET", r.sessionId, string(data)) //存入数据到 redis
	// 保存成功后更改 map 的状态
	r.flag = SessionFlagNone
	if err != nil {
		return
	}
	return
}

func (r *RedisSession)Get(key string) (result interface{}, err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	// 先判断session 在内存中的 map 中内容是否存在数据，没有数据则从 redis 中获取
	result, ok := r.sessionMap[key]
	if !ok {
		err = errors.New("key not exists")
	}
	//result, ok = result.(string)
	//if !ok {
	//	r.loadFromRedis()
	//}
	return
}

// 从 redis 再次加载数据到内存中的 map
func (r *RedisSession)loadFromRedis() (err error) {
	conn := r.pool.Get()
	reply, err := conn.Do("GET", r.sessionId)
	if err != nil {
		return
	}
	data, err := redis.String(reply, err)
	if err != nil {
		return
	}
	// 将取到的数据反序列化到内存的 map 中
	err = json.Unmarshal([]byte(data), &r.sessionMap)
	if err != nil {
		return
	}
	return
}

// Del 删除session 在 map 中的数据
func (r *RedisSession)Del(key string) (err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	r.flag = SessionFlagModify
	delete(r.sessionMap, key)
	return
}