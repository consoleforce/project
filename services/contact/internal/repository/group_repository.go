// services/contact/repository/group_repository.go

package repository

import (
	"database/sql"

	"github.com/consoleforce/project/services/contact/internal/domain"
)

type GroupRepositoryImpl struct {
	db *sql.DB
}

func NewGroupRepository(db *sql.DB) *GroupRepositoryImpl {
	return &GroupRepositoryImpl{
		db: db,
	}
}

func (r *GroupRepositoryImpl) Create(group *domain.Group) (*domain.Group, error) {
	// Реальная логика создания группы в базе данных
	_, err := r.db.Exec("INSERT INTO groups (name) VALUES ($1)", group.Name)
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (r *GroupRepositoryImpl) GetByID(groupID int) (*domain.Group, error) {
	// Реальная логика получения группы из базы данных по ее ID
	var group domain.Group
	err := r.db.QueryRow("SELECT id, name FROM groups WHERE id = $1", groupID).Scan(&group.ID, &group.Name)
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *GroupRepositoryImpl) AddContactToGroup(contactID, groupID int) error {
	// Реальная логика добавления контакта в группу в базе данных
	_, err := r.db.Exec("INSERT INTO contact_group (contact_id, group_id) VALUES ($1, $2)", contactID, groupID)
	if err != nil {
		return err
	}
	return nil
}
