// services/contact/internal/constructor.go

package internal

import (
	"github.com/consoleforce/project/services/contact/internal/delivery"
	"github.com/consoleforce/project/services/contact/internal/repository"
	"github.com/consoleforce/project/services/contact/internal/usecase"
)

type Constructor struct {
	ContactRepo repository.ContactRepository
	GroupRepo   repository.GroupRepository
	ContactUC   usecase.ContactUsecase
	GroupUC     usecase.GroupUsecase
	ContactDel  delivery.ContactDelivery
}

func NewConstructor(contactRepo repository.ContactRepository, groupRepo repository.GroupRepository) *Constructor {
	return &Constructor{
		ContactRepo: contactRepo,
		GroupRepo:   groupRepo,
		ContactUC:   usecase.NewContactUsecase(contactRepo, groupRepo),
		GroupUC:     usecase.NewGroupUsecase(groupRepo),
		ContactDel:  delivery.NewContactDelivery(),
	}
}
