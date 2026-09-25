package service

import (
	"strings"

	"cyskillswap/internal/constants"
	"cyskillswap/internal/errors"
	"cyskillswap/internal/logger"
	"cyskillswap/internal/model"
	"cyskillswap/internal/repository"
)

// SubmitResponse 同学对需求提交响应：校验入参、需求开放状态，并防止重复报名
func SubmitResponse(student string, needID int, input model.ResponseInput) (model.NeedResponse, error) {
	note := strings.TrimSpace(input.Note)
	if note == "" {
		return model.NeedResponse{}, errors.Validation("请填写交换说明")
	}
	if len(input.Slots) == 0 {
		return model.NeedResponse{}, errors.Validation("请至少选择一个空闲时段")
	}
	need, ok := repository.FindNeed(needID)
	if !ok {
		return model.NeedResponse{}, errors.NeedNotFound()
	}
	if need.Status != constants.NeedStatusOpen {
		return model.NeedResponse{}, errors.NeedClosed()
	}
	if need.Requester == student {
		return model.NeedResponse{}, errors.Validation("不能响应自己发布的需求")
	}
	if _, dup := repository.FindResponseByStudent(needID, student); dup {
		return model.NeedResponse{}, errors.DuplicateResponse()
	}
	resp := repository.AddResponse(model.NeedResponse{
		NeedID:  needID,
		Student: student,
		Note:    note,
		Slots:   append([]string(nil), input.Slots...),
		Status:  constants.ResponseStatusPending,
	})
	logger.Info("need response submitted", "need", needID, "student", student)
	return resp, nil
}
