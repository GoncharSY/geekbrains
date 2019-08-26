package phonebook

import "fmt"

// Phones - моделирует список телефонов.
type Phones []Phone

// Len - возвращает количество записей в списке.
func (p *Phones) Len() int {
	return len(*p)
}

// Swap - меняет местами записи в списке.
func (p *Phones) Swap(i, j int) {
	var list = *p

	list[i], list[j] = list[j], list[i]
}

// Less - проверяет, что первый элемент меньше второго.
// Сравнивает по фамилии, имени и номеру телефона.
func (p *Phones) Less(i, j int) bool {
	var phones = *p
	var iP, jP = phones[i], phones[j]

	// Сравним по фамилии.
	switch {
	case iP.LastName < jP.LastName:
		return true
	case iP.LastName > jP.LastName:
		return false
	}

	// Сравним по имени.
	switch {
	case iP.FirstName < jP.FirstName:
		return true
	case iP.FirstName > jP.FirstName:
		return false
	}

	// Сравним по номеру.
	return iP.Number < jP.Number
}

// Print - выведет список в консоль.
func (p *Phones) Print() {
	var phones = *p

	for i := range phones {
		phone := phones[i]
		fmt.Printf("%d - %s %s\n", phone.Number, phone.LastName, phone.FirstName)
	}
}
