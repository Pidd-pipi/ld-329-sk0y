package repository

import "cyskillswap/internal/model"

// ListNeeds 返回全部需求，并统计每条需求的响应人数
func ListNeeds() []model.Need {
	store.mu.RLock()
	defer store.mu.RUnlock()
	needs := append([]model.Need(nil), store.needs...)
	for i := range needs {
		needs[i].Responses = countResponsesLocked(needs[i].ID)
	}
	return needs
}

// FindNeed 按 ID 查找需求，附带响应人数
func FindNeed(id int) (model.Need, bool) {
	store.mu.RLock()
	defer store.mu.RUnlock()
	for _, need := range store.needs {
		if need.ID == id {
			need.Responses = countResponsesLocked(id)
			return need, true
		}
	}
	return model.Need{}, false
}

// countResponsesLocked 调用前必须已持有锁
func countResponsesLocked(needID int) int {
	count := 0
	for _, resp := range store.responses {
		if resp.NeedID == needID {
			count++
		}
	}
	return count
}
