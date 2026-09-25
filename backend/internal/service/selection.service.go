package service

import (
	"cyskillswap/internal/constants"
	"cyskillswap/internal/errors"
	"cyskillswap/internal/logger"
	"cyskillswap/internal/model"
	"cyskillswap/internal/repository"
)

// SelectCandidate 发起人从候选人中选定一人：
// 被选同学同时段已有其他预约则选择失败并返回冲突说明；
// 成功后按需求时间生成待双方确认的预约，需求关闭、其余响应标记为未入选。
func SelectCandidate(requester string, needID, responseID int) (model.Appointment, error) {
	need, ok := repository.FindNeed(needID)
	if !ok {
		return model.Appointment{}, errors.NeedNotFound()
	}
	if need.Requester != requester {
		return model.Appointment{}, errors.Forbidden("只有需求发起人可以选定候选人")
	}
	if need.Status != constants.NeedStatusOpen {
		return model.Appointment{}, errors.NeedClosed()
	}
	resp, ok := repository.FindResponse(responseID)
	if !ok || resp.NeedID != needID {
		return model.Appointment{}, errors.ResponseNotFound()
	}
	appt := model.Appointment{
		NeedID:    need.ID,
		Pair:      need.Requester + " ↔ " + resp.Student,
		Requester: need.Requester,
		Provider:  resp.Student,
		Time:      need.ExpectTime,
		Slot:      need.Slot,
		Place:     constants.DefaultAppointmentPlace,
		Status:    constants.AppointmentStatusPending,
		Agenda:    resp.Note,
	}
	created, conflict, ok := repository.TryCreateAppointment(appt, responseID)
	if !ok {
		return model.Appointment{}, errors.ScheduleConflict(resp.Student, need.Slot, conflict.Time+" · "+conflict.Place)
	}
	logger.Info("candidate selected", "need", needID, "student", resp.Student, "appointment", created.ID)
	return created, nil
}
