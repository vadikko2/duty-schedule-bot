package entities

import "errors"

// Дефолтное значение максимального числа дежурантов в очереди
const DefaultMaxOfficerinOrderCount int32 = 10

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
func (o *Officer) FirstName() string { return o.firstName }

// Возвращает фамилию дежурного
func (o *Officer) LastName() string { return o.lastName }

// Возвращает telegram логин дежурного
func (o *Officer) TelegramLogin() string { return o.telegramLogin }

// Возвращает номер телефона дежурного
func (o *Officer) PhoneNumber() string { return o.phoneNumber }

// Сравнивает дежуранто. Считаем, что дежуранты равны, если у них совпадает telegramLogin
func (o *Officer) Equals(officer *Officer) bool { return o.telegramLogin == officer.telegramLogin }

// Элемент очереди дежурантов
type OfficerOrderNode struct {
	value Officer           // Данные о текущем дежуранте
	next  *OfficerOrderNode // Следующий дежурант
}

// Создает новый NewOfficerOrderNode
func NewOfficerOrderNode(officer *Officer, next *OfficerOrderNode) *OfficerOrderNode {
	return &OfficerOrderNode{
		value: *officer,
		next:  next,
	}
}

// Заменяет next у OfficerOrderNode
func (oon *OfficerOrderNode) SetNext(node *OfficerOrderNode) { oon.next = node }

// Очередь дежурантов
type OfficerOrder struct {
	head            *OfficerOrderNode // Ссылка на последнего дежуранта в очереди
	tail            *OfficerOrderNode // Ссылка на первого дежуранта в очереди
	officerCount    int32
	maxOfficerCount int32
}

// Создает пустой OfficerOrder
func NewOfficerOrder(maxOfficerCount *int32) *OfficerOrder {
	var maxCountValue int32
	if maxOfficerCount == nil {
		maxCountValue = DefaultMaxOfficerinOrderCount
	} else {
		maxCountValue = *maxOfficerCount
	}
	return &OfficerOrder{
		maxOfficerCount: maxCountValue,
	}
}

// Добавляет новый элемент в OfficerOrder
// Добавляет новый элемент в OfficerOrder
func (oo *OfficerOrder) AddNewOfficer(officer *Officer) error {
	if oo.officerCount == oo.maxOfficerCount {
		return errors.New("Officer count already has maximum value")
	}

	if oo.head == nil {
		// Если список пустой, новый элемент ссылается сам на себя
		newNode := NewOfficerOrderNode(officer, nil)
		newNode.SetNext(newNode)
		oo.head = newNode
		oo.tail = newNode
	} else {
		// Вставляем новый элемент после head
		newNode := NewOfficerOrderNode(officer, oo.tail)
		oo.head.SetNext(newNode)
		oo.head = newNode
	}

	oo.officerCount++
	return nil
}

// Возвращает текущее значение количиства дежурантов
func (oo *OfficerOrder) OfficerCount() int32 { return oo.officerCount }

// Возвращает текущее значение максимального количества дежурантов
func (oo *OfficerOrder) MaxOfficerCount() int32 { return oo.maxOfficerCount }

// Возвращает количество доступных для заполнения слотов
func (oo *OfficerOrder) AvailableSlots() int32 { return oo.maxOfficerCount - oo.officerCount }

// Проверяет пуста ли очередь
func (oo *OfficerOrder) IsEmpty() bool { return oo.head == nil }
