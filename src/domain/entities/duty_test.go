package entities

import "testing"

// Проверяет корректность создания сущности расписания через конструктор
func TestCreateNewDutyPositive(t *testing.T) {
	officer := Officer{
		telegramLogin: "vadkoz",
	}
	newDuty, error := NewDuty(&officer, DefaultDutyDuration)
	officer.telegramLogin = "vadikko2"

	if error != nil {
		t.Fatal("Duty created with error", error)
	}
	if newDuty.Officer().TelegramLogin() == "vadikko2" {
		t.Errorf("Expected vadkoz, got %s", newDuty.Officer().TelegramLogin())
	}
}

// Проверяет что нельзя создать дежурство с нулевой длительностью
func TestCreatenewDutyNegative(t *testing.T) {
	officer := Officer{
		telegramLogin: "vadkoz",
	}
	var zeroDuration uint8 = 0
	newDuty, error := NewDuty(&officer, zeroDuration)

	if error == nil {
		t.Errorf("Duty with zero duration can not be created")
	}
	if newDuty != nil {
		t.Errorf("Duty pointer is not nil but excepted")
	}
}
