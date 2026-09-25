package repository

import (
	"cyskillswap/internal/constants"
	"cyskillswap/internal/model"
)

func ListAppointments() []model.Appointment {
	store.mu.RLock()
	defer store.mu.RUnlock()
	return append([]model.Appointment(nil), store.appointments...)
}

// FindAppointmentByNeed 查找由指定需求生成的预约
func FindAppointmentByNeed(needID int) (model.Appointment, bool) {
	store.mu.RLock()
	defer store.mu.RUnlock()
	for _, appt := range store.appointments {
		if appt.NeedID == needID {
			return appt, true
		}
	}
	return model.Appointment{}, false
}

// ConflictingAppointment 查找某同学在指定时段已有的预约
func ConflictingAppointment(student, slot string) (model.Appointment, bool) {
	store.mu.RLock()
	defer store.mu.RUnlock()
	return findConflictLocked(student, slot)
}

// findConflictLocked 调用前必须已持有锁
func findConflictLocked(student, slot string) (model.Appointment, bool) {
	for _, appt := range store.appointments {
		if appt.Slot != slot {
			continue
		}
		if appt.Requester == student || appt.Provider == student {
			return appt, true
		}
	}
	return model.Appointment{}, false
}

// TryCreateAppointment 在单把写锁内完成冲突复检、预约落库、响应状态更新与需求关闭，
// 保证同一时段不会被重复预约。返回已创建的预约、冲突的预约和是否创建成功。
func TryCreateAppointment(appt model.Appointment, responseID int) (model.Appointment, model.Appointment, bool) {
	store.mu.Lock()
	defer store.mu.Unlock()
	if conflict, ok := findConflictLocked(appt.Provider, appt.Slot); ok {
		return model.Appointment{}, conflict, false
	}
	appt.ID = store.nextAppointmentID
	store.nextAppointmentID++
	store.appointments = append(store.appointments, appt)
	for i := range store.responses {
		resp := &store.responses[i]
		if resp.ID == responseID {
			resp.Status = constants.ResponseStatusSelected
		} else if resp.NeedID == appt.NeedID {
			// 需求已定人，其余响应标记为未入选，停止接收
			resp.Status = constants.ResponseStatusRejected
		}
	}
	for i := range store.needs {
		if store.needs[i].ID == appt.NeedID {
			store.needs[i].Status = constants.NeedStatusMatched
		}
	}
	return appt, model.Appointment{}, true
}
