package entities

import (
	"testing"
	"time"
)

// Проверяем что дежурство добавляется в пустое расписание
func TestAddFirstDutyPositive(t *testing.T) {
	emptyDutySchedule := Schedule{}
	newDuty := Duty{
		officer:      Officer{telegramLogin: "vadkoz"},
		durationDays: DefaultDutyDuration,
	}

	error := emptyDutySchedule.addToSchedule(&newDuty)

	if error != nil {
		t.Errorf("Adding to Schedule with error %v", error)
	}
	if emptyDutySchedule.elements == nil {
		t.Errorf("Schedule is empty after add new duty %v", emptyDutySchedule)
	}
	if emptyDutySchedule.firstStar == nil {
		t.Errorf("Schedule has no fromDateTime after add new duty %v", emptyDutySchedule)
	}
	if emptyDutySchedule.lastEnd == nil {
		t.Errorf("Schedule has no fromDateTime after add new duty %v", emptyDutySchedule)
	}
	if emptyDutySchedule.elements.Len() != 1 {
		t.Errorf("Schedule has to has 1 elements, but has %d", emptyDutySchedule.elements.Len())
	}
	if _, error := emptyDutySchedule.LastEnd(); error != nil {
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
	dutySchedule.addToSchedule(&duty1)

	error := dutySchedule.addToSchedule(&duty2)

	if error != nil {
		t.Errorf("Adding to Schedule with error %v", error)
	}
	if dutySchedule.elements == nil {
		t.Errorf("Schedule is empty after add new duty %v", dutySchedule)
	}
	if dutySchedule.firstStar == nil {
		t.Errorf("Schedule has no fromDateTime after add new duty %v", dutySchedule)
	}
	if dutySchedule.lastEnd == nil {
		t.Errorf("Schedule has no fromDateTime after add new duty %v", dutySchedule)
	}
	if dutySchedule.elements.Len() != 2 {
		t.Errorf("Schedule has to has 2 elements, but has %d", dutySchedule.elements.Len())
	}
	if _, error := dutySchedule.LastEnd(); error != nil {
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
	dutySchedule.addToSchedule(&duty1)
	dutySchedule.addToSchedule(&duty2)

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
