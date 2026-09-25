package repository

// defaultStore 是单进程演示环境共享的存储实例。
var defaultStore = NewStore()

// DB 返回全局存储，service 层通过它加锁实现跨实体的事务性操作。
func DB() *Store { return defaultStore }
