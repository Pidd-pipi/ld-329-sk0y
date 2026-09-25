package service

import (
	"cyskillswap/internal/constants"
	"cyskillswap/internal/errors"
	"cyskillswap/internal/model"
	"cyskillswap/internal/repository"
)

// Needs 返回当前用户视角的需求列表（含我的响应、我的预约、候选人）
func Needs() []model.NeedView {
	return needViews(constants.CurrentUser)
}

func needViews(user string) []model.NeedView {
	needs := repository.ListNeeds()
	views := make([]model.NeedView, 0, len(needs))
	for _, need := range needs {
		views = append(views, buildNeedView(user, need))
	}
	return views
}

// buildNeedView 组装单张需求卡：自己的响应、自己参与的预约、发起人可见的候选人
func buildNeedView(user string, need model.Need) model.NeedView {
	view := model.NeedView{Need: need}
	if resp, ok := repository.FindResponseByStudent(need.ID, user); ok {
		view.MyResponse = &resp
	}
	if appt, ok := repository.FindAppointmentByNeed(need.ID); ok {
		if appt.Requester == user || appt.Provider == user {
			view.MyAppointment = &appt
		}
	}
	if need.Requester == user {
		view.Candidates = repository.ListResponses(need.ID)
	}
	return view
}

// Candidates 发起人查看某条需求的候选人列表
func Candidates(user string, needID int) ([]model.NeedResponse, error) {
	need, ok := repository.FindNeed(needID)
	if !ok {
		return nil, errors.NeedNotFound()
	}
	if need.Requester != user {
		return nil, errors.Forbidden("只有需求发起人可以查看候选人")
	}
	return repository.ListResponses(needID), nil
}
