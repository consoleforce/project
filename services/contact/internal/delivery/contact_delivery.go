// services/contact/delivery/contact_delivery.go

package delivery

import (
	"context"
	"encoding/json"
	"github.com/consoleforce/project/services/contact/internal/domain"
	"github.com/consoleforce/project/services/contact/internal/usecase"
	"net/http"
)

type ContactDelivery struct {
	contactUsecase usecase.ContactUsecase
}

func NewContactDelivery(contactUsecase usecase.ContactUsecase) *ContactDelivery {
	return &ContactDelivery{
		contactUsecase: contactUsecase,
	}
}

func (d *ContactDelivery) CreateContactHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// Добавляем переменную ID типа UUID в контекст (для примера используем константу)
	ctx = context.WithValue(ctx, "ID", "your-uuid-here")

	var newContact domain.Contact
	err := json.NewDecoder(r.Body).Decode(&newContact)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	createdContact, err := d.contactUsecase.CreateContact(ctx, newContact.FirstName, newContact.LastName, newContact.MiddleName, newContact.PhoneNumber)
	if err != nil {
		http.Error(w, "Error creating contact", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdContact)
}

// Другие методы Delivery
