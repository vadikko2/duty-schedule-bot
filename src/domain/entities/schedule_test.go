package entities

import (
	"testing"
	"time"
)

// Проверяет сбор расписания через метод ConstructSchedule
func TestConstructSchedulePositive(t *testing.T) {
	fistDuty, _ := NewDuty(&Officer{}, 10)
	secondDuty, _ := NewDuty(&Officer{}, 10)
	ExceptedFirstDutyStart := time.Date(2025, 11, 4, 9, 0, 0, 0, time.UTC)
	ExceptedLastDutyEnd := time.Date(2025, 11, 24, 9, 0, 0, 0, time.UTC)
	elements := []ScheduleElement{
		*NewScheduleElement(
			nil,
			ExceptedFirstDutyStart,
			fistDuty,
		),
		*NewScheduleElement(
			nil,
			time.Date(2025, 11, 14, 9, 0, 0, 0, time.UTC),
			secondDuty,
		),
	}

	schedule, error := ConstructSchedule(&elements)

	if schedule == nil {
		t.Fatal("Schedule has been custructed with error:", error)
	}
	if error != nil {
		t.Errorf("Schedule has been custructed with errur %v", error)
	}
	if (*schedule).DutyCount() != 2 {
		t.Errorf("Excepted duty count is 2, but got %v", schedule.DutyCount())
	}
	if *schedule.firstDutyStart != ExceptedFirstDutyStart {
		t.Errorf("Excepted ExceptedFirstDutyStart %v but got %v", ExceptedFirstDutyStart, *schedule.firstDutyStart)
	}
	if *schedule.lastDutyEnd != ExceptedLastDutyEnd {
		t.Errorf("Excepted ExceptedLastDutyEnd %v but got %v", ExceptedLastDutyEnd, *schedule.lastDutyEnd)
	}
}

// Проверяет сбор расписания через метод ConstructSchedule из пустого среза
func TestConstructScheduleFromEmptySlicePositive(t *testing.T) {
	emtySlice := []ScheduleElement{}
	schedule, error := ConstructSchedule(&emtySlice)

	if error != nil {
		t.Errorf("Empty schedule has been cunstructed with error %v", error)
	}
	if schedule.elements != nil {
		t.Errorf("Excepted elements is nil, but got %v", schedule.elements)
	}

}

// Проверяем что дежурство добавляется в пустое расписание
func TestAddFirstDutyPositive(t *testing.T) {
	emptyDutySchedule := Schedule{}
	newDuty := Duty{
		officer:      Officer{telegramLogin: "vadkoz"},
		durationDays: DefaultDutyDuration,
	}

	error := emptyDutySchedule.AddToSchedule(&newDuty)

	if error != nil {
		t.Errorf("Adding to Schedule with error %v", error)
	}
	if emptyDutySchedule.elements == nil {
		t.Errorf("Schedule is empty after add new duty %v", emptyDutySchedule)
	}
	if emptyDutySchedule.firstDutyStart == nil {
		t.Errorf("Schedule has no fromDateTime after add new duty %v", emptyDutySchedule)
	}
	if emptyDutySchedule.lastDutyEnd == nil {
		t.Errorf("Schedule has no fromDateTime after add new duty %v", emptyDutySchedule)
	}
	if emptyDutySchedule.elements.Len() != 1 {
		t.Errorf("Schedule has to has 1 elements, but has %d", emptyDutySchedule.elements.Len())
	}
	if _, error := emptyDutySchedule.LastDutyEnd(); error != nil {
		t.Errorf("Getting LastEnd ends with error %v", error)
	}
}

// Проверяем что очередное дежуство добавляется в не пустое расписание
func TestAddNewDutyPositive(t *testing.T) {
	dutySchedule := Schedule{}
	duty1 := Duty{
		officer:      Officer{telegramLogin: "vadkoz"},
		durationDays: DefaultDutyDuration,
	}
	duty2 := Duty{
		officer:      Officer{telegramLogin: "nkunov"},
		durationDays: DefaultDutyDuration,
	}
	dutySchedule.AddToSchedule(&duty1)

	error := dutySchedule.AddToSchedule(&duty2)

	if error != nil {
		t.Errorf("Adding to Schedule with error %v", error)
	}
	if dutySchedule.elements == nil {
		t.Errorf("Schedule is empty after add new duty %v", dutySchedule)
	}
	if dutySchedule.firstDutyStart == nil {
		t.Errorf("Schedule has no fromDateTime after add new duty %v", dutySchedule)
	}
	if dutySchedule.lastDutyEnd == nil {
		t.Errorf("Schedule has no fromDateTime after add new duty %v", dutySchedule)
	}
	if dutySchedule.elements.Len() != 2 {
		t.Errorf("Schedule has to has 2 elements, but has %d", dutySchedule.elements.Len())
	}
	if _, error := dutySchedule.LastDutyEnd(); error != nil {
		t.Errorf("Getting LastEnd ends with error %v", error)
	}
}

func TestGetDutyPositive(t *testing.T) {
	dutySchedule := Schedule{}
	duty1 := Duty{
		officer:      Officer{telegramLogin: "vadkoz"},
		durationDays: DefaultDutyDuration,
	}
	duty2 := Duty{
		officer:      Officer{telegramLogin: "nkunov"},
		durationDays: DefaultDutyDuration,
	}
	dutySchedule.AddToSchedule(&duty1)
	dutySchedule.AddToSchedule(&duty2)

	nKunovDuty, error := dutySchedule.getDuty(time.Now().AddDate(0, 0, 8))

	if nKunovDuty == nil {
		t.Fatal("getDuty returned nil duty: ", error)
	}
	if error != nil {
		t.Errorf("Getting duty ends with error: %v", error)
	}
	if nKunovDuty.Officer().TelegramLogin() != duty2.Officer().TelegramLogin() {
		t.Errorf("Expected telegram login %v but got %v", duty2.Officer().TelegramLogin(), nKunovDuty.Officer().TelegramLogin())
	}
}
