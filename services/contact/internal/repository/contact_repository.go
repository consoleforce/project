package repository

import "github.com/consoleforce/project/services/contact/internal/domain"

type ContactRepository interface {
	Create(contact *domain.Contact) (*domain.Contact, error)
	Update(contact *domain.Contact) error
	Delete(contactID int) error
	GetByID(contactID int) (*domain.Contact, error)
}

type GroupRepository interface {
	Create(group *domain.Group) (*domain.Group, error)
	GetByID(groupID int) (*domain.Group, error)
	AddContactToGroup(contactID, groupID int) error
}
