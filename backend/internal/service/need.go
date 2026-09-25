package service

import (
	"strings"
	"time"

	"cyskillswap/internal/constants"
	"cyskillswap/internal/errors"
	"cyskillswap/internal/model"
	"cyskillswap/internal/repository"
)

// RespondRequest 是同学提交响应的入参。
type RespondRequest struct {
	OfferNote string   `json:"offerNote"`
	FreeSlots []string `json:"freeSlots"`
}

// ListNeedCards 组装需求卡：候选人、当前用户的响应与预约状态一并返回。
func ListNeedCards(actor string) []model.NeedCard {
	db := repository.DB()
	db.RLock()
	defer db.RUnlock()

	needs := db.ListNeeds()
	cards := make([]model.NeedCard, 0, len(needs))
	for _, need := range needs {
		card := model.NeedCard{Need: need}
		resps := db.ResponsesByNeed(need.ID)
		card.Responses = resps
		card.ResponseCount = len(resps)
		for i := range resps {
			if resps[i].Respondent == actor {
				r := resps[i]
				card.MyResponse = &r
				break
			}
		}
		if apt := db.ActiveAppointmentByNeed(need.ID); apt != nil {
			a := *apt
			card.Appointment = &a
		}
		cards = append(cards, card)
	}
	return cards
}

// Respond 同学对需求提交交换说明与空闲时段；同一条需求不允许重复报名。
func Respond(actor string, needID int, req RespondRequest) (model.NeedCard, error) {
	note := strings.TrimSpace(req.OfferNote)
	if note == "" {
		return model.NeedCard{}, errors.New(400, constants.CodeValidationFailed, "请填写交换说明")
	}
	if len(note) > 200 {
		return model.NeedCard{}, errors.New(400, constants.CodeValidationFailed, "交换说明不能超过 200 字")
	}
	if len(req.FreeSlots) == 0 {
		return model.NeedCard{}, errors.New(400, constants.CodeValidationFailed, "请至少选择一个空闲时段")
	}
	seen := map[string]bool{}
	for _, code := range req.FreeSlots {
		if seen[code] {
			continue
		}
		if !constants.IsValidSlot(code) {
			return model.NeedCard{}, errors.New(400, constants.CodeValidationFailed, "存在无效的空闲时段："+constants.SlotLabel(code))
		}
		seen[code] = true
	}

	db := repository.DB()
	db.Lock()
	defer db.Unlock()

	need, _ := db.FindNeed(needID)
	if need == nil {
		return model.NeedCard{}, errors.ErrNeedNotFound
	}
	if need.Requester == actor {
		return model.NeedCard{}, errors.ErrOwnNeed
	}
	if need.Status == constants.NeedStatusBooked {
		return model.NeedCard{}, errors.ErrNeedClosed
	}
	if db.HasResponse(needID, actor) {
		return model.NeedCard{}, errors.ErrDuplicateResponse
	}

	db.AddResponse(needID, actor, note, dedupeSlots(req.FreeSlots), time.Now())
	return buildNeedCard(db, need, actor), nil
}

// Select 发起人从候选人中选定一人，按需求时间生成待双方确认的预约，其余响应停止接收。
// 若被选同学在该时段已有其他未取消预约，选择失败并返回冲突详情，需求仍可改选别人。
func Select(actor string, needID, responseID int) (model.Appointment, error) {
	db := repository.DB()
	db.Lock()
	defer db.Unlock()

	need, _ := db.FindNeed(needID)
	if need == nil {
		return model.Appointment{}, errors.ErrNeedNotFound
	}
	if need.Requester != actor {
		return model.Appointment{}, errors.ErrNotRequester
	}
	if need.Status == constants.NeedStatusBooked {
		return model.Appointment{}, errors.New(409, constants.CodeNeedAlreadyBooked, "该需求已选定同学，如需更换请先取消当前预约")
	}

	resp, _ := db.FindResponse(responseID)
	if resp == nil || resp.NeedID != needID {
		return model.Appointment{}, errors.ErrResponseNotFound
	}

	// 被选同学的空闲时段需覆盖需求时间。
	covers := false
	for _, code := range resp.FreeSlots {
		if code == need.SlotCode {
			covers = true
			break
		}
	}
	if !covers {
		return model.Appointment{}, errors.New(
			409, constants.CodeSlotMismatch,
			resp.Respondent+" 的空闲时段不包含需求时间（"+constants.SlotLabel(need.SlotCode)+"），请改选其他人",
		)
	}

	// 冲突检测：同一时段与其他人的 pending/confirmed 预约均占用该同学的时间。
	if conflict := db.FindConflictingAppointment(resp.Respondent, need.SlotCode, needID); conflict != nil {
		other := conflict.Requester
		if conflict.Requester == resp.Respondent {
			other = conflict.Respondent
		}
		details := map[string]any{
			"respondent":     resp.Respondent,
			"slotCode":       need.SlotCode,
			"slotLabel":      constants.SlotLabel(need.SlotCode),
			"conflictWith":   other,
			"conflictPlace":  conflict.Place,
			"conflictStatus": conflict.Status,
			"conflictApptId": conflict.ID,
		}
		return model.Appointment{}, errors.NewWithDetails(
			409, constants.CodeRespondentBusy,
			resp.Respondent+" 在 "+constants.SlotLabel(need.SlotCode)+" 已有预约（与 "+other+"，"+conflict.Place+"），请改选其他同学",
			details,
		)
	}

	apt := db.AddAppointment(
		needID, need.Requester, resp.Respondent, need.SlotCode,
		defaultPlace(need), defaultAgenda(need, resp), time.Now(),
	)
	db.MarkNeedBooked(needID, responseID)
	return apt, nil
}

// buildNeedCard 在已持锁的前提下组装单张需求卡。
func buildNeedCard(db *repository.Store, need *model.Need, actor string) model.NeedCard {
	card := model.NeedCard{Need: *need}
	resps := db.ResponsesByNeed(need.ID)
	card.Responses = resps
	card.ResponseCount = len(resps)
	for i := range resps {
		if resps[i].Respondent == actor {
			r := resps[i]
			card.MyResponse = &r
		}
	}
	if apt := db.ActiveAppointmentByNeed(need.ID); apt != nil {
		a := *apt
		card.Appointment = &a
	}
	return card
}

func dedupeSlots(slots []string) []string {
	out := make([]string, 0, len(slots))
	seen := map[string]bool{}
	for _, code := range slots {
		if !seen[code] {
			seen[code] = true
			out = append(out, code)
		}
	}
	return out
}

// defaultPlace / defaultAgenda 在发起人未另行协商前给出预约初始信息。
func defaultPlace(need *model.Need) string {
	if need.Campus != "" {
		return need.Campus + "（待双方确认）"
	}
	return "地点待协商"
}

func defaultAgenda(need *model.Need, resp *model.Response) string {
	return "围绕「" + need.Title + "」开展交换：" + resp.OfferNote
}
