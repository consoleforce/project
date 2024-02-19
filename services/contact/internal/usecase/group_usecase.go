package usecase

import (
	"errors"
	"github.com/consoleforce/project/services/contact/internal/domain"
	"github.com/consoleforce/project/services/contact/internal/repository"
)

type GroupUsecaseImpl struct {
	groupRepo repository.GroupRepository
}

func NewGroupUsecase(groupRepo repository.GroupRepository) *GroupUsecaseImpl {
	return &GroupUsecaseImpl{
		groupRepo: groupRepo,
	}
}

func (uc *GroupUsecaseImpl) CreateGroup(name string) (*domain.Group, error) {
	// Проверка длины названия группы
	if len(name) > 250 {
		return nil, errors.New("group name is too long")
	}
	// Создание новой группы
	group := &domain.Group{
		Name: name,
	}
	createdGroup, err := uc.groupRepo.Create(group)
	if err != nil {
		return nil, err
	}
	return createdGroup, nil
}

func (uc *GroupUsecaseImpl) GetGroupByID(groupID int) (*domain.Group, error) {
	// Получение группы по ID
	group, err := uc.groupRepo.GetByID(groupID)
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (uc *GroupUsecaseImpl) AddContactToGroup(contactID, groupID int) error {
	// Проверка существования контакта
	_, err := uc.contactRepo.GetByID(contactID)
	if err != nil {
		return errors.New("contact not found")
	}
	// Проверка существования группы
	_, err = uc.groupRepo.GetByID(groupID)
	if err != nil {
		return errors.New("group not found")
	}
	// Добавление контакта в группу
	err = uc.groupRepo.AddContactToGroup(contactID, groupID)
	if err != nil {
		return err
	}
	return nil
}
