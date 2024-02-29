package repository

import (
	"context"
	"database/sql"
	"github.com/consoleforce/project/services/contact/internal/domain"
)

type ContactRepositoryImpl struct {
	db *sql.DB
}

func NewContactRepository(db *sql.DB) *ContactRepositoryImpl {
	return &ContactRepositoryImpl{
		db: db,
	}
}

func (r *ContactRepositoryImpl) Create(contact *domain.Contact) (*domain.Contact, error) {
	// Реальная логика создания контакта в базе данных
	_, err := r.db.Exec("INSERT INTO contacts (first_name, last_name, middle_name, phone_number) VALUES ($1, $2, $3, $4)", contact.FirstName, contact.LastName, contact.MiddleName, contact.PhoneNumber)
	if err != nil {
		return nil, err
	}
	return contact, nil
}

func (r *ContactRepositoryImpl) Update(contact *domain.Contact) error {
	// Реальная логика обновления контакта в базе данных
	_, err := r.db.Exec("UPDATE contacts SET first_name = $1, last_name = $2, middle_name = $3, phone_number = $4 WHERE id = $5", contact.FirstName, contact.LastName, contact.MiddleName, contact.PhoneNumber, contact.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *ContactRepositoryImpl) Delete(contactID int) error {
	// Реальная логика удаления контакта из базы данных
	_, err := r.db.Exec("DELETE FROM contacts WHERE id = $1", contactID)
	if err != nil {
		return err
	}
	return nil
}

func (r *ContactRepositoryImpl) GetByID(contactID int) (*domain.Contact, error) {
	// Реальная логика получения контакта из базы данных по его ID
	var contact domain.Contact
	err := r.db.QueryRow("SELECT id, first_name, last_name, middle_name, phone_number FROM contacts WHERE id = $1", contactID).Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.MiddleName, &contact.PhoneNumber)
	if err != nil {
		return nil, err
	}
	return &contact, nil
}

type ContactRepository interface {
	Create(ctx context.Context, id, firstName, lastName, middleName, phoneNumber string) (*domain.Contact, error)
}

type contactRepository struct {
	db *sql.DB
}

func NewContactRepository(db *sql.DB) ContactRepository {
	return &contactRepository{
		db: db,
	}
}

func (r *contactRepository) Create(ctx context.Context, id, firstName, lastName, middleName, phoneNumber string) (*domain.Contact, error) {
	// Здесь вы можете использовать ID и передать его в запрос к базе данных
	return nil, nil
}
