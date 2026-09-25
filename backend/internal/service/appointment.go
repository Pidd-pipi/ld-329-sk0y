package service

import (
	"cyskillswap/internal/constants"
	"cyskillswap/internal/errors"
	"cyskillswap/internal/model"
	"cyskillswap/internal/repository"
)

// ListAppointments 返回当前可见预约（演示环境返回全部）。
func ListAppointments() []model.Appointment {
	db := repository.DB()
	db.RLock()
	defer db.RUnlock()
	return db.ListAppointments()
}

// ConfirmAppointment 预约双方中的任一方确认；双方都确认后预约生效。
func ConfirmAppointment(actor string, appointmentID int) (model.Appointment, error) {
	db := repository.DB()
	db.Lock()
	defer db.Unlock()

	apt, _ := db.FindAppointment(appointmentID)
	if apt == nil {
		return model.Appointment{}, errors.ErrApptNotFound
	}
	if actor != apt.Requester && actor != apt.Respondent {
		return model.Appointment{}, errors.ErrNotApptParty
	}
	if apt.Status == constants.AppointmentCancelled {
		return model.Appointment{}, errors.ErrAppointmentClosed
	}
	if apt.Status == constants.AppointmentConfirmed {
		return model.Appointment{}, errors.ErrAlreadyConfirmed
	}

	if actor == apt.Requester {
		apt.RequesterOK = true
	} else {
		apt.RespondentOK = true
	}
	if apt.RequesterOK && apt.RespondentOK {
		apt.Status = constants.AppointmentConfirmed
	}
	return *apt, nil
}

// CancelAppointment 发起人或被选同学取消待确认预约：需求重新开放、候选人恢复候选，可改选别人。
func CancelAppointment(actor string, appointmentID int) (model.Appointment, error) {
	db := repository.DB()
	db.Lock()
	defer db.Unlock()

	apt, _ := db.FindAppointment(appointmentID)
	if apt == nil {
		return model.Appointment{}, errors.ErrApptNotFound
	}
	if actor != apt.Requester && actor != apt.Respondent {
		return model.Appointment{}, errors.ErrNotApptParty
	}
	if apt.Status == constants.AppointmentCancelled {
		return model.Appointment{}, errors.ErrAppointmentClosed
	}
	if apt.Status == constants.AppointmentConfirmed {
		return model.Appointment{}, errors.New(409, constants.CodeAppointmentLocked, "预约已双方确认，请先联系对方协商后再取消")
	}

	apt.Status = constants.AppointmentCancelled
	if apt.NeedID != 0 {
		db.ReopenNeed(apt.NeedID)
	}
	return *apt, nil
}
