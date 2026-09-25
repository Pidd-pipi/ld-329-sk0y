package repository

import "cyskillswap/internal/model"

func ListMessages() []model.Conversation {
	store.mu.RLock()
	defer store.mu.RUnlock()
	return append([]model.Conversation(nil), store.conversations...)
}
