package delivery

import (
	"github.com/consoleforce/project/services/contact/internal/usecase"
)

// Delivery представляет слой доставки данных
type Delivery struct {
	uc usecase.UseCase
	// Здесь можно добавить поля, если необходимо
}

// NewDelivery создает новый экземпляр слоя доставки данных
func NewDelivery(uc usecase.UseCase) *Delivery {
	return &Delivery{uc: uc}
}
