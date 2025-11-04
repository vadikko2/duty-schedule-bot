package entities

import (
	"container/list"
	"errors"
	"time"
)

// Элемент расписания
type ScheduleElement struct {
	dutyStart time.Time
	dutyEnd   time.Time
	duty      Duty
}

func NewScheduleElement(dutyStart, dutyEnd time.Time, duty *Duty) *ScheduleElement {
	return &ScheduleElement{
		dutyStart: dutyStart,
		dutyEnd:   dutyEnd,
		duty:      *duty,
	}
}

// Возвращает дату начала дежурства
func (dse *ScheduleElement) DutyStart() time.Time { return dse.dutyStart }

// Возвращает дату окончания дежурства
func (dse *ScheduleElement) DutyEnd() time.Time { return dse.dutyEnd }

// Возвращает сущность дежурства
func (dse *ScheduleElement) Duty() *Duty { return &dse.duty }

// Расписание дежурств
type Schedule struct {
	firstStar *time.Time
	lastEnd   *time.Time
	elements  *list.List
}

// Возвращает жату окончания последнего дежурства в расписании
func (ds *Schedule) LastEnd() (time.Time, error) {
	if ds.lastEnd == nil {
		return time.Time{}, errors.New("Schedule is empty")
	}
	return *ds.lastEnd, nil
}

// Добавляет новое дежурство в расписание
func (ds *Schedule) addToSchedule(duty *Duty) error {
	// Если расписание пустое, ты мы дробавляем в него первый элемент
	if ds.elements == nil {
		ds.elements = list.New()
		now := time.Now()
		newDutyElement := NewScheduleElement(
			now,
			now.AddDate(0, 0, int(duty.DurationDays())),
			duty,
		)
		ds.elements.PushBack(newDutyElement)
		// Проставляем fromDateTime и toDateTime
		ds.firstStar = &newDutyElement.dutyStart
		ds.lastEnd = &newDutyElement.dutyEnd
		return nil
	}

	// Добавляем duty в конец расписания
	lastEnd, error := ds.LastEnd()
	if error != nil {
		return errors.New("Schedule has no lastEnd")
	}
	newDutyElement := NewScheduleElement(
		lastEnd,
		lastEnd.AddDate(0, 0, int(duty.DurationDays())),
		duty,
	)
	ds.elements.PushBack(newDutyElement)
	// Обновляем toDateTime
	ds.lastEnd = &newDutyElement.dutyEnd
	return nil
}

// Возвращает информацию о текущем дежурстве на заданную дату
func (ds *Schedule) getDuty(dateTime time.Time) (*Duty, error) {
	for element := ds.elements.Front(); element != nil; element = element.Next() {
		scheduleElement, ok := element.Value.(*ScheduleElement)
		if !ok {
			return nil, errors.New("Problem with processing element: incorrect type")
		}
		if dateTime.After(scheduleElement.DutyStart()) && dateTime.Before(scheduleElement.DutyEnd()) {
			return scheduleElement.Duty(), nil
		}
	}
	return nil, errors.New("Duty with specified parameters not found")
}
