package storage

import (
	"database/sql"
	"fmt"
	"github.com/nitin-kukreti/GoChat/internal/domain"
)

type GroupRepositoryImpl struct {
	db *sql.DB
}

func NewGroupRepository(db *sql.DB) domain.GroupRepository {
	return &GroupRepositoryImpl{db: db}
}

func (r *GroupRepositoryImpl) CreateGroup(name string) (domain.Group, error) {
	var group domain.Group
	query := "SELECT * FROM create_group($1);"
	if err := r.db.QueryRow(query, name).Scan(&group.ID, &group.Name); err != nil {
		return domain.Group{}, fmt.Errorf("error while creating group %s: %w", name, err)
	}
	return group, nil
}

func (r *GroupRepositoryImpl) AddUserToGroup(userID, groupID int) error {
	query := "SELECT add_user_to_group($1, $2);"
	if _, err := r.db.Exec(query, userID, groupID); err != nil {
		return fmt.Errorf("unable to map user %d to group %d: %w", userID, groupID, err)
	}
	return nil
}
