package repositories

import (
	"duty-schedule-bot/src/domain/entities"
	"time"
)

type ScheduleRepository interface {
	// Возвращает дежурства по указанным датам
	GetScheduleByDates(fromDate, toDate time.Time) (*entities.Schedule, error)
	// Сохраняет дежурства в базе данных
	Save(schedule *entities.Schedule) error
}
