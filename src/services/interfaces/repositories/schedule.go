package repositories

import (
	"duty-schedule-bot/src/domain/entities"
	"time"
)

type ScheduleRepository interface {
	GetScheduleByDates(fromDate, toDate time.Time) *entities.Schedule
	Save(schedule *entities.Schedule) error
}
