package repository

import (
	"cyskillswap/internal/constants"
	"cyskillswap/internal/model"
)

func GetProfile() model.Profile {
	skills := ListSkills()
	skillWall := skills
	if len(skillWall) > 1 {
		skillWall = skillWall[:1]
	}
	return model.Profile{
		Name: constants.CurrentUser, Major: "新闻传播 2023", CreditScore: 91, CreditLevel: constants.CreditGold,
		SkillWall: skillWall,
		Radar:     map[string]int{"摄影": 92, "修图": 86, "沟通": 90, "编程": 42, "乐器": 35},
		History:   []string{"完成毕业照拍摄交换", "响应 Python 数据分析需求", "预约吉他入门课"},
		Reviews:   ListReviews(),
	}
}
