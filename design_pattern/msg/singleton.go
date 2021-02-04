package msg

import "sync"

// MessagePool 消息池
type MessagePool struct {
	pool *sync.Pool
}
// Message 消息
type Message struct {
	Count int32
}
// 消息池单例
var msgPool = &MessagePool{
	// 如果消息池里没有消息，则新建一个Count值为0的Message实例
	pool: &sync.Pool{New: func() interface{} { return &Message{Count: 0} }},
}
// Instance 访问消息池单例的唯一方法
func Instance() *MessagePool {
	return msgPool
}
// 往消息池里添加消息
func (m *MessagePool) AddMsg(msg *Message) {
	m.pool.Put(msg)
}
// 从消息池里获取消息
func (m *MessagePool) GetMsg() *Message {
	return m.pool.Get().(*Message)
}
