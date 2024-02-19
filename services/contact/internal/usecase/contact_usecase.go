// services/contact/usecase/contact_usecase.go

package usecase

import (
	"errors"
	"github.com/consoleforce/project/services/contact/internal/domain"
	"github.com/consoleforce/project/services/contact/internal/repository"
)

type ContactUsecaseImpl struct {
	contactRepo repository.ContactRepository
	groupRepo   repository.GroupRepository
}

func NewContactUsecase(contactRepo repository.ContactRepository, groupRepo repository.GroupRepository) *ContactUsecaseImpl {
	return &ContactUsecaseImpl{
		contactRepo: contactRepo,
		groupRepo:   groupRepo,
	}
}

func (uc *ContactUsecaseImpl) CreateContact(firstName, lastName, middleName, phoneNumber string) (*domain.Contact, error) {
	// Проверка валидности номера телефона
	if len(phoneNumber) != 11 {
		return nil, errors.New("invalid phone number")
	}
	// Создание нового контакта
	contact := &domain.Contact{
		FirstName:   firstName,
		LastName:    lastName,
		MiddleName:  middleName,
		PhoneNumber: phoneNumber,
	}
	createdContact, err := uc.contactRepo.Create(contact)
	if err != nil {
		return nil, err
	}
	return createdContact, nil
}

func (uc *ContactUsecaseImpl) UpdateContact(contact *domain.Contact) error {
	// Проверка существования контакта
	_, err := uc.contactRepo.GetByID(contact.ID)
	if err != nil {
		return errors.New("contact not found")
	}
	err = uc.contactRepo.Update(contact)
	if err != nil {
		return err
	}
	return nil
}

func (uc *ContactUsecaseImpl) DeleteContact(contactID int) error {
	// Проверка существования контакта
	_, err := uc.contactRepo.GetByID(contactID)
	if err != nil {
		return errors.New("contact not found")
	}
	err = uc.contactRepo.Delete(contactID)
	if err != nil {
		return err
	}
	return nil
}

func (uc *ContactUsecaseImpl) GetContactByID(contactID int) (*domain.Contact, error) {
	// Получение контакта по ID
	contact, err := uc.contactRepo.GetByID(contactID)
	if err != nil {
		return nil, err
	}
	return contact, nil
}
