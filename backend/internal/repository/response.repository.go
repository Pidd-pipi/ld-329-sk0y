package repository

import "cyskillswap/internal/model"

// ListResponses 返回某条需求的全部响应（候选人）
func ListResponses(needID int) []model.NeedResponse {
	store.mu.RLock()
	defer store.mu.RUnlock()
	result := make([]model.NeedResponse, 0)
	for _, resp := range store.responses {
		if resp.NeedID == needID {
			result = append(result, resp)
		}
	}
	return result
}

// FindResponse 按 ID 查找响应
func FindResponse(id int) (model.NeedResponse, bool) {
	store.mu.RLock()
	defer store.mu.RUnlock()
	for _, resp := range store.responses {
		if resp.ID == id {
			return resp, true
		}
	}
	return model.NeedResponse{}, false
}

// FindResponseByStudent 查找某同学对指定需求的响应，用于防止重复报名
func FindResponseByStudent(needID int, student string) (model.NeedResponse, bool) {
	store.mu.RLock()
	defer store.mu.RUnlock()
	for _, resp := range store.responses {
		if resp.NeedID == needID && resp.Student == student {
			return resp, true
		}
	}
	return model.NeedResponse{}, false
}

// AddResponse 新增响应并分配 ID
func AddResponse(resp model.NeedResponse) model.NeedResponse {
	store.mu.Lock()
	defer store.mu.Unlock()
	resp.ID = store.nextResponseID
	store.nextResponseID++
	store.responses = append(store.responses, resp)
	return resp
}
