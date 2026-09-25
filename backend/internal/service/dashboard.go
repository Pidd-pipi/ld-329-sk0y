package service

import (
	"cyskillswap/internal/constants"
	"cyskillswap/internal/model"
	"cyskillswap/internal/repository"
)

// Overview 聚合看板数据；需求卡按当前同学视角附带响应与预约状态。
func Overview(actor string) model.Overview {
	skills := repository.ListSkills()
	needCards := ListNeedCards(actor)
	matches := repository.ListMatches()
	appointments := ListAppointments()
	reviews := repository.ListReviews()
	messages := repository.ListMessages()

	activeAppointments := 0
	for _, apt := range appointments {
		if apt.Status != constants.AppointmentCancelled {
			activeAppointments++
		}
	}

	return model.Overview{
		Service:         constants.ServiceName,
		CurrentUser:     actor,
		SwitchableUsers: constants.SwitchableUsers,
		Categories:      constants.SkillCategories,
		Slots:           buildSlotOptions(),
		Metrics: map[string]int{
			"skills": len(skills), "needs": len(needCards), "matches": len(matches),
			"appointments": activeAppointments, "reviews": len(reviews), "unread": 3,
		},
		Skills: skills, Needs: needCards, Matches: matches,
		Appointments: appointments, Reviews: reviews, Messages: messages,
		Profile: repository.GetProfile(),
	}
}

func buildSlotOptions() []model.SlotOption {
	out := make([]model.SlotOption, 0, len(constants.FreeSlots))
	for _, slot := range constants.FreeSlots {
		out = append(out, model.SlotOption{Code: slot.Code, Label: slot.Label})
	}
	return out
}

func Skills() []model.Skill          { return repository.ListSkills() }
func Matches() []model.Match         { return repository.ListMatches() }
func Reviews() []model.Review        { return repository.ListReviews() }
func Messages() []model.Conversation { return repository.ListMessages() }
func Profile() model.Profile         { return repository.GetProfile() }
