package entities

import "errors"

const DefaultDutyDuration uint8 = 7

// Дежурство
type Duty struct {
	officer      Officer
	durationDays uint8
}

// Конструктор. Вернет ссылку на иммутабельный объект Duty
func NewDuty(officer *Officer, durationDays uint8) (*Duty, error) {
	if officer == nil {
		return nil, errors.New("Offser is necessary to create duty")
	}
	if durationDays == 0 {
		return nil, errors.New("Duty duration have to be greater than zero")
	}
	return &Duty{
		officer:      *officer,
		durationDays: durationDays,
	}, nil
}

// Возвращает дежуранта
func (d *Duty) Officer() *Officer { return &d.officer }

// Возвращает длительность дежурства в днях
func (d *Duty) DurationDays() uint8 { return d.durationDays }
