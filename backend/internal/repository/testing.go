package repository

// NewTestStore 暴露给 service 层测试使用的全新内存存储。
func NewTestStore() *Store { return NewStore() }

// SetStoreForTest 替换全局存储实例。
func SetStoreForTest(s *Store) { defaultStore = s }
