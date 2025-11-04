package entities

// Дежурант
type Officer struct {
	firstName     string
	lastName      string
	telegramLogin string
	phoneNumber   string
}

// Конструктор. Создание ссылку на новый иммутабельный объект "Дежурант"
func NewOfficer(firstName, lastName, telegramLogin, phoneNumber string) *Officer {
	return &Officer{
		firstName:     firstName,
		lastName:      lastName,
		telegramLogin: telegramLogin,
		phoneNumber:   phoneNumber,
	}
}

// Возвращает имя дежурного
func (d *Officer) FirstName() string { return d.firstName }

// Возвращает фамилию дежурного
func (d *Officer) LastName() string { return d.lastName }

// Возвращает telegram логин дежурного
func (d *Officer) TelegramLogin() string { return d.telegramLogin }

// Возвращает номер телефона дежурного
func (d *Officer) PhoneNumber() string { return d.phoneNumber }
