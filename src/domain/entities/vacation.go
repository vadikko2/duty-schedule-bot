package entities

import (
	"errors"
	"time"
)

// Отпуск
type Vacation struct {
	officer       Officer
	vacationStart time.Time
	durationDays  uint8
}

// Конструктор. Возвращает ссылку на новый объект отпуска
func NewVacation(officer *Officer, vacationStart time.Time, durationDays uint8) (*Vacation, error) {
	if durationDays == 0 {
		return nil, errors.New("Vacation duration have to be greater than zero")
	}
	return &Vacation{
		officer:       *officer,
		vacationStart: vacationStart,
		durationDays:  durationDays,
	}, nil
}

// Возвращает ссылку на дежуранта, который находится в отпуске
func (v *Vacation) Officer() *Officer { return &v.officer }

// Возвращает дату начала отпуска
func (v *Vacation) VacationStart() time.Time { return v.vacationStart }

// Возвращает дату окончания отпуска
func (v *Vacation) VacationEnd() time.Time { return v.vacationStart.AddDate(0, 0, int(v.durationDays)) }

// Возвращает длительность отпуска
func (v *Vacation) DurationDays() uint8 { return v.durationDays }
