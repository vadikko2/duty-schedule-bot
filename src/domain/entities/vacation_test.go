package entities

import (
	"testing"
	"time"
)

// Проверяет что дата окочания отпуска считается корректно
func TestVacationEndCalculatingPositive(t *testing.T) {
	officer := Officer{telegramLogin: "vadkoz"}
	vacationStart := time.Date(2025, 11, 4, 9, 0, 0, 0, time.UTC)
	expectedVacationEnd := time.Date(2025, 11, 14, 9, 0, 0, 0, time.UTC)
	var vacationDuration uint8 = 10
	vacation, error := NewVacation(&officer, vacationStart, vacationDuration)

	vacationEnd := vacation.VacationEnd()

	if error != nil {
		t.Errorf("Vacation created with error %v", error)
	}
	if expectedVacationEnd != vacationEnd {
		t.Errorf("Expected vacation end %v but got %v", expectedVacationEnd, vacationEnd)
	}
}

// Проверяет что нельщя создать отпуск с нулевой длительностью
func TestVacationEndCalculatingNegative(t *testing.T) {
	officer := Officer{telegramLogin: "vadkoz"}
	vacationStart := time.Date(2025, 11, 4, 9, 0, 0, 0, time.UTC)
	var vacationDuration uint8 = 0
	vacation, error := NewVacation(&officer, vacationStart, vacationDuration)

	if error == nil {
		t.Fatal("Vacation have to be created with error")
	}
	if vacation != nil {
		t.Error("Vacation have to be nil")
	}
}
