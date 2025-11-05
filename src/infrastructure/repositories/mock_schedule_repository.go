package repositories

import (
	"duty-schedule-bot/src/domain/entities"
	"sort"
	"time"

	"github.com/google/uuid"
)

type MockScheduleRepository struct {
	elements map[uuid.UUID]entities.ScheduleElement
}

func (r *MockScheduleRepository) GetScheduleByDates(startDate, endDate time.Time) (*entities.Schedule, error) {
	var foundedElements []entities.ScheduleElement
	if r.elements == nil {
		return &entities.Schedule{}, nil
	}
	for _, element := range r.elements {
		if element.DutyStart().After(startDate) && element.DutyEnd().Before(endDate) {
			foundedElements = append(foundedElements, element)
		}
	}
	sort.Slice(foundedElements, func(i, j int) bool {
		return foundedElements[i].DutyStart().Before(foundedElements[j].DutyStart())
	})
	result, error := entities.ConstructSchedule(&foundedElements)
	if error != nil {
		return nil, error
	}
	return result, nil
}

func (r *MockScheduleRepository) Save(schedule *entities.Schedule) error {
	for _, element := range schedule.Elements() {
		r.elements[element.ID()] = *element
	}
	return nil
}
