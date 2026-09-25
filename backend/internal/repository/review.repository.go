package repository

import "cyskillswap/internal/model"

func ListReviews() []model.Review {
	store.mu.RLock()
	defer store.mu.RUnlock()
	return append([]model.Review(nil), store.reviews...)
}
