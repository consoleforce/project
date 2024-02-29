// services/contact/delivery/group_delivery.go

package delivery

import (
	"github.com/consoleforce/project/services/contact/internal/usecase"
	"net/http"
)

type GroupDelivery struct {
	groupUsecase usecase.GroupUsecase
}

func NewGroupDelivery(groupUsecase usecase.GroupUsecase) *GroupDelivery {
	return &GroupDelivery{
		groupUsecase: groupUsecase,
	}
}

func (d *GroupDelivery) CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
	// Реализация создания группы
}

func (d *GroupDelivery) GetGroupHandler(w http.ResponseWriter, r *http.Request) {
	// Реализация получения группы
}
