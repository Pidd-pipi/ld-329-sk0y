package repository

import (
	"time"

	"cyskillswap/internal/constants"
	"cyskillswap/internal/model"
)

// 以下方法假定调用方已通过 DB() 持有相应的锁。

// FindNeed 按 ID 查找需求；返回值指向存储内对象，调用方不得在锁外持有。
func (s *Store) FindNeed(id int) (*model.Need, int) {
	for i := range s.needs {
		if s.needs[i].ID == id {
			return &s.needs[i], i
		}
	}
	return nil, -1
}

// ListNeeds 返回需求快照副本。
func (s *Store) ListNeeds() []model.Need {
	out := make([]model.Need, len(s.needs))
	copy(out, s.needs)
	return out
}

// FindResponse 按 ID 查找响应。
func (s *Store) FindResponse(id int) (*model.Response, int) {
	for i := range s.responses {
		if s.responses[i].ID == id {
			return &s.responses[i], i
		}
	}
	return nil, -1
}

// ResponsesByNeed 返回某条需求下的全部响应（按报名先后）。
func (s *Store) ResponsesByNeed(needID int) []model.Response {
	out := make([]model.Response, 0)
	for i := range s.responses {
		if s.responses[i].NeedID == needID {
			out = append(out, s.responses[i])
		}
	}
	return out
}

// HasResponse 判断同学是否已对该需求报名（含已选中/已落选的历史记录）。
func (s *Store) HasResponse(needID int, respondent string) bool {
	for i := range s.responses {
		if s.responses[i].NeedID == needID && s.responses[i].Respondent == respondent {
			return true
		}
	}
	return false
}

// AddResponse 追加一条报名并返回新记录。
func (s *Store) AddResponse(needID int, respondent, note string, slots []string, createdAt time.Time) model.Response {
	resp := model.Response{
		ID: s.nextResponseID, NeedID: needID, Respondent: respondent,
		OfferNote: note, FreeSlots: slots, Status: constants.ResponseWaiting, CreatedAt: createdAt,
	}
	s.nextResponseID++
	s.responses = append(s.responses, resp)
	return resp
}

// MarkNeedBooked 将需求置为已选定：中选响应 selected，其余在候响应 released。
func (s *Store) MarkNeedBooked(needID, selectedResponseID int) {
	if need, _ := s.FindNeed(needID); need != nil {
		need.Status = constants.NeedStatusBooked
	}
	for i := range s.responses {
		r := &s.responses[i]
		if r.NeedID != needID {
			continue
		}
		switch r.ID {
		case selectedResponseID:
			r.Status = constants.ResponseSelected
		default:
			if r.Status == constants.ResponseWaiting {
				r.Status = constants.ResponseReleased
			}
		}
	}
}

// ReopenNeed 取消预约后恢复需求：需求开放，全部相关响应回到候选中。
func (s *Store) ReopenNeed(needID int) {
	if need, _ := s.FindNeed(needID); need != nil {
		need.Status = constants.NeedStatusOpen
	}
	for i := range s.responses {
		r := &s.responses[i]
		if r.NeedID == needID && r.Status != constants.ResponseWaiting {
			r.Status = constants.ResponseWaiting
		}
	}
}
