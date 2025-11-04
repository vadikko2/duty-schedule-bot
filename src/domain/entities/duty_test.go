package entities

import "testing"

// Проверяет корректность создания сущности расписания через конструктор
func TestCreateNewDutyPositive(t *testing.T) {
	officer := Officer{
		telegramLogin: "vadkoz",
	}
	newDuty := NewDuty(&officer, DefaultDutyDuration)
	officer.telegramLogin = "vadikko2"

	if newDuty.Officer().TelegramLogin() == "vadikko2" {
		t.Errorf("Expected vadkoz, got %s", newDuty.Officer().TelegramLogin())
	}
}
