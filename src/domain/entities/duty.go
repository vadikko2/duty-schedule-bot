package entities

const DefaultDutyDuration uint8 = 7

// Дежурство
type Duty struct {
	officer      Officer
	durationDays uint8
}

// Конструктор. Вернет ссылку на иммутабельный объект Duty
func NewDuty(officer *Officer, durationDays uint8) *Duty {
	return &Duty{
		officer:      *officer,
		durationDays: durationDays,
	}
}

// Возвращает дежуранта
func (d *Duty) Officer() *Officer { return &d.officer }

// Возвращает длительность дежурства в днях
func (d *Duty) DurationDays() uint8 { return d.durationDays }
