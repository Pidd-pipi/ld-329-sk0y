package repository

import (
	"time"

	"cyskillswap/internal/constants"
	"cyskillswap/internal/model"
)

// ListAppointments 返回预约快照（含已取消），调用方可按状态过滤。
func (s *Store) ListAppointments() []model.Appointment {
	out := make([]model.Appointment, len(s.appointments))
	copy(out, s.appointments)
	return out
}

// FindAppointment 按 ID 查找预约。
func (s *Store) FindAppointment(id int) (*model.Appointment, int) {
	for i := range s.appointments {
		if s.appointments[i].ID == id {
			return &s.appointments[i], i
		}
	}
	return nil, -1
}

// ActiveAppointmentByNeed 返回需求当前生效中的预约（pending/confirmed），无则 nil。
func (s *Store) ActiveAppointmentByNeed(needID int) *model.Appointment {
	for i := range s.appointments {
		apt := &s.appointments[i]
		if apt.NeedID == needID && apt.Status != constants.AppointmentCancelled {
			return apt
		}
	}
	return nil
}

// FindConflictingAppointment 检查同学在某时段是否已有其他未取消预约。
// excludeNeedID 用于排除同一需求自身的预约（改选时旧预约已取消，不会命中）。
// 同学既可能以响应者身份、也可能以发起人身份占用该时段，两种都算冲突。
func (s *Store) FindConflictingAppointment(person, slotCode string, excludeNeedID int) *model.Appointment {
	for i := range s.appointments {
		apt := &s.appointments[i]
		if apt.SlotCode != slotCode || apt.Status == constants.AppointmentCancelled || apt.NeedID == excludeNeedID {
			continue
		}
		if apt.Respondent == person || apt.Requester == person {
			return apt
		}
	}
	return nil
}

// AddAppointment 按需求时间生成一条待双方确认的预约。
func (s *Store) AddAppointment(needID int, requester, respondent, slotCode, place, agenda string, createdAt time.Time) model.Appointment {
	apt := model.Appointment{
		ID: s.nextAppointmentID, NeedID: needID,
		Requester: requester, Respondent: respondent,
		SlotCode: slotCode, Time: constants.SlotLabel(slotCode),
		Place: place, Agenda: agenda,
		Status: constants.AppointmentPending, CreatedAt: createdAt,
	}
	s.nextAppointmentID++
	s.appointments = append(s.appointments, apt)
	return apt
}
