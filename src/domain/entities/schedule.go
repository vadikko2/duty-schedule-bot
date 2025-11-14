package entities

import (
	"container/list"
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
)

// Элемент расписания
type ScheduleElement struct {
	uid       uuid.UUID
	dutyStart time.Time
	dutyEnd   time.Time
	duty      Duty
}

func NewScheduleElement(dutyUuid *uuid.UUID, dutyStart time.Time, duty *Duty) *ScheduleElement {
	var initalUuid uuid.UUID
	if dutyUuid == nil {
		initalUuid = uuid.New()
	} else {
		initalUuid = *dutyUuid
	}
	return &ScheduleElement{
		uid:       initalUuid,
		dutyStart: dutyStart,
		dutyEnd:   dutyStart.AddDate(0, 0, int(duty.DurationDays())),
		duty:      *duty,
	}
}

// Возвращает дату начала дежурства
func (se *ScheduleElement) DutyStart() time.Time { return se.dutyStart }

// Возвращает дату окончания дежурства
func (se *ScheduleElement) DutyEnd() time.Time { return se.dutyEnd }

// Возвращает сущность дежурства
func (se *ScheduleElement) Duty() *Duty { return &se.duty }

// Возвращает идентификатор дежурства
func (se *ScheduleElement) ID() uuid.UUID { return se.uid }

// Расписание дежурств
type Schedule struct {
	firstDutyStart *time.Time
	lastDutyEnd    *time.Time
	elements       *list.List
	mutex          sync.RWMutex
}

// Собирает расписание из списка ScheduleElement
func ConstructSchedule(elements *[]ScheduleElement) (*Schedule, error) {
	if elements == nil || len(*elements) == 0 {
		return &Schedule{}, nil
	}
	elementsList := list.New()
	firstScheduleElementDutyStart := (*elements)[0].DutyStart()
	lastScheduleElementDutyEnd := (*elements)[len(*elements)-1].DutyEnd()
	for _, element := range *elements {
		elementsList.PushBack(element)
	}
	return &Schedule{&firstScheduleElementDutyStart, &lastScheduleElementDutyEnd, elementsList, sync.RWMutex{}}, nil
}

// Возвращает жату окончания последнего дежурства в расписании
func (s *Schedule) LastDutyEnd() (time.Time, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if s.lastDutyEnd == nil {
		return time.Time{}, errors.New("Schedule is empty")
	}
	return *s.lastDutyEnd, nil
}

// Добавляет новое дежурство в расписание
func (s *Schedule) AddToSchedule(duty *Duty) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Если расписание пустое, ты мы дробавляем в него первый элемент
	if s.elements == nil {
		s.elements = list.New()
		now := time.Now()
		newDutyElement := NewScheduleElement(
			nil,
			now,
			duty,
		)
		s.elements.PushBack(newDutyElement)
		// Проставляем fromDateTime и toDateTime
		s.firstDutyStart = &newDutyElement.dutyStart
		s.lastDutyEnd = &newDutyElement.dutyEnd
		return nil
	}

	// Добавляем duty в конец расписания
	if s.lastDutyEnd == nil {
		return errors.New("Schedule has no lastEnd")
	}
	newDutyElement := NewScheduleElement(nil, *s.lastDutyEnd, duty)
	s.elements.PushBack(newDutyElement)
	s.lastDutyEnd = &newDutyElement.dutyEnd
	return nil
}

// Возвращает информацию о текущем дежурстве на заданную дату
func (s *Schedule) getDuty(dateTime time.Time) (*Duty, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	// Проверяем, что расписание не пустое
	if s.elements == nil {
		return nil, errors.New("Schedule is empty")
	}

	for element := s.elements.Front(); element != nil; element = element.Next() {
		scheduleElement, ok := element.Value.(*ScheduleElement)
		if !ok {
			return nil, errors.New("Problem with processing element: incorrect type")
		}
		eq_or_after := (dateTime.Equal(scheduleElement.DutyStart()) || dateTime.After(scheduleElement.DutyStart()))
		eq_or_before := (dateTime.Equal(scheduleElement.DutyEnd()) || dateTime.Before(scheduleElement.DutyEnd()))
		if eq_or_after && eq_or_before {
			return scheduleElement.Duty(), nil
		}
	}
	return nil, errors.New("Duty with specified parameters not found")
}

// Возвращает количество дежурств в расписании за заданный период
func (s *Schedule) DutyCount() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if s.elements == nil {
		return 0
	} else {
		return s.elements.Len()
	}
}

// Возвращает список элементов расписания дежурств
func (s *Schedule) Elements() []*ScheduleElement {
	if s.elements == nil {
		return nil
	}
	var elements []*ScheduleElement
	for element := s.elements.Front(); element != nil; element = element.Next() {
		scheduleElement, ok := element.Value.(*ScheduleElement)
		if !ok {
			continue
		}
		elements = append(elements, scheduleElement)
	}
	return elements
}
