// services/contact/internal/domain/contact.go

package domain

type Contact struct {
	ID          int
	FirstName   string
	LastName    string
	MiddleName  string
	PhoneNumber string
}

func (c *Contact) FullName() string {
	return c.LastName + " " + c.FirstName + " " + c.MiddleName
}
