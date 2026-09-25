package repository

import "cyskillswap/internal/model"

func ListMatches() []model.Match {
	store.mu.RLock()
	defer store.mu.RUnlock()
	return append([]model.Match(nil), store.matches...)
}
