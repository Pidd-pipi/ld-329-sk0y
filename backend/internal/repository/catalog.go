package repository

import "cyskillswap/internal/model"

func ListSkills() []model.Skill {
	defaultStore.RLock()
	defer defaultStore.RUnlock()
	out := make([]model.Skill, len(defaultStore.skills))
	copy(out, defaultStore.skills)
	return out
}

func ListMatches() []model.Match {
	defaultStore.RLock()
	defer defaultStore.RUnlock()
	out := make([]model.Match, len(defaultStore.matches))
	copy(out, defaultStore.matches)
	return out
}

func ListReviews() []model.Review {
	defaultStore.RLock()
	defer defaultStore.RUnlock()
	out := make([]model.Review, len(defaultStore.reviews))
	copy(out, defaultStore.reviews)
	return out
}

func ListMessages() []model.Conversation {
	defaultStore.RLock()
	defer defaultStore.RUnlock()
	out := make([]model.Conversation, len(defaultStore.messages))
	copy(out, defaultStore.messages)
	return out
}

func GetProfile() model.Profile {
	defaultStore.RLock()
	defer defaultStore.RUnlock()
	return defaultStore.profile
}
