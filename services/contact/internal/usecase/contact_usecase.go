// services/contact/usecase/contact_usecase.go

package usecase

import (
	"context"
	"github.com/consoleforce/project/services/contact/internal/domain"
	"github.com/consoleforce/project/services/contact/internal/repository"
)

type ContactUsecase interface {
	CreateContact(ctx context.Context, firstName, lastName, middleName, phoneNumber string) (*domain.Contact, error)
}

type contactUsecase struct {
	contactRepository repository.ContactRepository
}

func NewContactUsecase(contactRepository repository.ContactRepository) ContactUsecase {
	return &contactUsecase{
		contactRepository: contactRepository,
	}
}

func (uc *contactUsecase) CreateContact(ctx context.Context, firstName, lastName, middleName, phoneNumber string) (*domain.Contact, error) {
	// Получаем ID из контекста
	id := ctx.Value("ID").(string)

	// Реализация создания контакта
	// Пример передачи ID в репозиторий
	// createdContact, err := uc.contactRepository.Create(ctx, id, firstName, lastName, middleName, phoneNumber)
	return nil, nil
}
