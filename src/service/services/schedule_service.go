package services

import "duty-schedule-bot/src/service/interfaces/repositories"

// Сервис для работы с расписанием
type ScheduleService struct {
	repo repositories.ScheduleRepository
}

// Конструктор для создания нового экземпляра сервиса расписания
func NewScheduleService(repo repositories.ScheduleRepository) *ScheduleService {
	return &ScheduleService{repo: repo}
}
