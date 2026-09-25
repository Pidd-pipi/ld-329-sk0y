package service

import (
	"testing"

	"cyskillswap/internal/constants"
	"cyskillswap/internal/errors"
	"cyskillswap/internal/repository"
)

// resetForTest 用全新内存存储，保证测试间互不干扰。
func resetForTest() {
	resetStore(repository.NewTestStore())
}

func TestRespondDuplicateRejected(t *testing.T) {
	resetForTest()
	req := RespondRequest{OfferNote: "我可以帮忙", FreeSlots: []string{constants.SlotSatMorning}}
	if _, err := Respond("林澈", 1, req); err == nil {
		t.Fatal("种子数据中林澈已响应 need 1，重复报名应失败")
	}
	// 需求仍应处于招募中
	cards := ListNeedCards("林澈")
	for _, c := range cards {
		if c.ID == 1 && c.Status != constants.NeedStatusOpen {
			t.Fatalf("重复报名后需求状态应为 open，实际 %s", c.Status)
		}
	}
}

func TestRespondOwnNeedRejected(t *testing.T) {
	resetForTest()
	_, err := Respond("孟野", 1, RespondRequest{OfferNote: "自己", FreeSlots: []string{constants.SlotSatMorning}})
	if biz, ok := err.(errors.BusinessError); !ok || biz.Code != constants.CodeOwnNeed {
		t.Fatalf("应返回 CodeOwnNeed，实际 %v", err)
	}
}

func TestSelectConflictKeepsNeedOpen(t *testing.T) {
	resetForTest()
	// 孟野 need1（周六上午）选林澈：林澈该时段已与顾远有约
	_, err := Select("孟野", 1, 1)
	biz, ok := err.(errors.BusinessError)
	if !ok || biz.Code != constants.CodeRespondentBusy {
		t.Fatalf("应返回时段冲突错误，实际 %v", err)
	}
	if biz.Details["conflictWith"] != "顾远" {
		t.Fatalf("冲突对象应为顾远，实际 %v", biz.Details["conflictWith"])
	}
	cards := ListNeedCards("孟野")
	for _, c := range cards {
		if c.ID == 1 {
			if c.Status != constants.NeedStatusOpen {
				t.Fatal("冲突选择失败后需求必须保持 open，可改选别人")
			}
			if c.Appointment != nil {
				t.Fatal("冲突时不应生成预约")
			}
		}
	}

	// 改选唐宁（response 2，周六上午有空且无冲突）应成功
	apt, err := Select("孟野", 1, 2)
	if err != nil {
		t.Fatalf("改选应成功: %v", err)
	}
	if apt.Status != constants.AppointmentPending || apt.Requester != "孟野" || apt.Respondent != "唐宁" {
		t.Fatalf("新预约状态不正确: %+v", apt)
	}
	// 选定后停止接收新响应
	_, err = Respond("周芮", 1, RespondRequest{OfferNote: "新来的", FreeSlots: []string{constants.SlotSatMorning}})
	if biz, ok := err.(errors.BusinessError); !ok || biz.Code != constants.CodeNeedClosed {
		t.Fatalf("选定后应停止接收响应，实际 %v", err)
	}
}

func TestSelectRejectsSlotMismatch(t *testing.T) {
	resetForTest()
	// need2 周二晚；唐宁(response 5)空闲只有周四晚/周日全天
	_, err := Select("许安", 2, 5)
	if biz, ok := err.(errors.BusinessError); !ok || biz.Code != constants.CodeSlotMismatch {
		t.Fatalf("应返回时段不包含错误，实际 %v", err)
	}
}

func TestConfirmFlow(t *testing.T) {
	resetForTest()
	apt, _ := Select("孟野", 1, 2)
	// 非当事人不能确认
	if _, err := ConfirmAppointment("韩梅", apt.ID); err == nil {
		t.Fatal("非预约双方不应能确认")
	}
	apt, err := ConfirmAppointment("孟野", apt.ID)
	if err != nil || apt.Status != constants.AppointmentPending || !apt.RequesterOK {
		t.Fatalf("发起人单方确认后仍应为 pending: %+v, %v", apt, err)
	}
	apt, err = ConfirmAppointment("唐宁", apt.ID)
	if err != nil || apt.Status != constants.AppointmentConfirmed {
		t.Fatalf("双方确认后应为 confirmed: %+v, %v", apt, err)
	}
}

func TestCancelReopensNeedForReselect(t *testing.T) {
	resetForTest()
	apt, _ := Select("林澈", 3, 6) // need3 选孟野
	if _, err := CancelAppointment("孟野", apt.ID); err != nil {
		t.Fatalf("取消待确认预约失败: %v", err)
	}
	cards := ListNeedCards("林澈")
	for _, c := range cards {
		if c.ID == 3 {
			if c.Status != constants.NeedStatusOpen {
				t.Fatal("取消后需求应重新开放")
			}
			for _, r := range c.Responses {
				if r.Status != constants.ResponseWaiting {
					t.Fatalf("取消后候选人应恢复 waiting，%s 为 %s", r.Respondent, r.Status)
				}
			}
		}
	}
	// 取消后响应立即恢复接收
	if _, err := Respond("韩梅", 3, RespondRequest{OfferNote: "会吉他", FreeSlots: []string{constants.SlotWedEvening}}); err != nil {
		t.Fatalf("取消后应能重新报名: %v", err)
	}
	// 然后可改选周芮(response 7)
	if _, err := Select("林澈", 3, 7); err != nil {
		t.Fatalf("改选应成功: %v", err)
	}
}

func TestNonRequesterCannotSelect(t *testing.T) {
	resetForTest()
	if _, err := Select("林澈", 1, 2); err == nil {
		t.Fatal("非发起人不应能选定候选人")
	}
}
