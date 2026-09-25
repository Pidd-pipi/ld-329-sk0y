package repository

import (
	"sync"

	"cyskillswap/internal/model"
)

// Store 内存数据存储，读写加锁保证并发安全
type Store struct {
	mu                sync.RWMutex
	skills            []model.Skill
	needs             []model.Need
	responses         []model.NeedResponse
	appointments      []model.Appointment
	matches           []model.Match
	reviews           []model.Review
	conversations     []model.Conversation
	nextResponseID    int
	nextAppointmentID int
}

var store = newSeededStore()
