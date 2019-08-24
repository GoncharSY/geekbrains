package phonebook

// Phone - описывает телефонный номер.
type Phone struct {
	Number    int    `json:"number"    xml:"Number"`
	FirstName string `json:"firstName" xml:"Name>First"`
	LastName  string `json:"lastName"  xml:"Name>Last"`
}
