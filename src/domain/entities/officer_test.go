package entities

import "testing"

// Проверяет что добавление дежуранта в пустую очередь завершается успехом
func TestAddOfficerIntoEmptyOrderPositive(t *testing.T) {
	var count uint32 = 10
	emptyOfficerOrder := NewOfficerOrder(&count)
	officer := Officer{telegramLogin: "vadkoz"}

	error := emptyOfficerOrder.AddNewOfficer(&officer)

	if error != nil {
		t.Fatal("Officer added int empty order with error ", error)
	}
	if emptyOfficerOrder.IsEmpty() != false {
		t.Errorf("Empty officer order is empty")
	}
	if emptyOfficerOrder.OfficerCount() != 1 {
		t.Errorf("Expected officer count is 1 but got %v", emptyOfficerOrder.OfficerCount())
	}
	if emptyOfficerOrder.head.value.telegramLogin != "vadkoz" {
		t.Errorf("Excepted head of order is vadkoz but got %v", emptyOfficerOrder.head.value.telegramLogin)
	}
	if emptyOfficerOrder.tail.value.telegramLogin != "vadkoz" {
		t.Errorf("Excepted tail of order is vadkoz but got %v", emptyOfficerOrder.head.value.telegramLogin)
	}
	if emptyOfficerOrder.head.next.value != emptyOfficerOrder.tail.value {
		t.Errorf("Excepted head is links to tail but has not")
	}
}

// Проверяет что добавление дежурного в не пусутую очередь завершается успехом
func TestAddOfficerIntoNotEmptyOrderPositive(t *testing.T) {
	var count uint32 = 4
	officerOrder := NewOfficerOrder(&count)
	officer1 := Officer{telegramLogin: "vadkoz"}
	officer2 := Officer{telegramLogin: "vadkoz2"}
	officer3 := Officer{telegramLogin: "vadkoz3"}

	officerOrder.AddNewOfficer(&officer1)
	officerOrder.AddNewOfficer(&officer2)
	error := officerOrder.AddNewOfficer(&officer3)

	if error != nil {
		t.Fatal("Officer added int empty order with error ", error)
	}
	if officerOrder.IsEmpty() != false {
		t.Errorf("Empty officer order is empty")
	}
	if officerOrder.OfficerCount() != 3 {
		t.Errorf("Expected officer count is 3 but got %v", officerOrder.OfficerCount())
	}
	if officerOrder.AvailableSlots() != 1 {
		t.Errorf("Expected AvailableSlots is 1 but got %v", officerOrder.AvailableSlots())
	}
	if officerOrder.head.value.telegramLogin != "vadkoz3" {
		t.Errorf("Excepted head of order is vadkoz2 but got %v", officerOrder.head.value.telegramLogin)
	}
	if officerOrder.tail.value.telegramLogin != "vadkoz" {
		t.Errorf("Excepted tail of order is vadkoz but got %v", officerOrder.head.value.telegramLogin)
	}
	if officerOrder.head.next.value != officerOrder.tail.value {
		t.Errorf("Excepted head is links to tail but has not")
	}
}

func TestAddOfficerIntoFullOrder(t *testing.T) {
	var count uint32 = 1
	officerOrder := NewOfficerOrder(&count)
	officer1 := Officer{telegramLogin: "vadkoz"}
	officer2 := Officer{telegramLogin: "vadkoz2"}
	expectedError := "Officer count already has maximum value"
	officerOrder.AddNewOfficer(&officer1)
	error := officerOrder.AddNewOfficer(&officer2)

	if error == nil {
		t.Fatalf("Add new officer into full order have to return error but hast not: %v", error)
	}
	if error.Error() != expectedError {
		t.Errorf("Error has unexpected error: %v", error)
	}
	if officerOrder.OfficerCount() != 1 {
		t.Errorf("Expected OfficerCount is 1 but got %v", officerOrder.OfficerCount())
	}
	if officerOrder.IsEmpty() != false {
		t.Errorf("Empty officer order is empty")
	}
	if officerOrder.AvailableSlots() != 0 {
		t.Errorf("Expected AvailableSlots is 0 but got %v", officerOrder.AvailableSlots())
	}

}
