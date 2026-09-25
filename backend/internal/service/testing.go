package service

import "cyskillswap/internal/repository"

// resetStore 替换 service 层间接使用的全局存储。
func resetStore(s *repository.Store) {
	repository.SetStoreForTest(s)
}
