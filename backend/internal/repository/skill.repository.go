package repository

import "cyskillswap/internal/model"

func ListSkills() []model.Skill {
	store.mu.RLock()
	defer store.mu.RUnlock()
	return append([]model.Skill(nil), store.skills...)
}
