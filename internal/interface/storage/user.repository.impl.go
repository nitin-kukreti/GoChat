package storage

import (
	"database/sql"
	"fmt"
	"github.com/nitin-kukreti/GoChat/internal/domain"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) CreateUser(name string) (domain.User, error) {
var user domain.User
	query := `SELECT * FROM create_user($1);`
	err := r.db.QueryRow(query, name).Scan(&user.ID, &user.Name)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to create user: %w", err)
	}
	return user, nil
}

func (r *UserRepositoryImpl) GetUserByID(id int) (domain.User, error) {
	 var user domain.User;
	 query:=`select * from get_user_by_id($1);`;
	 if err :=r.db.QueryRow(query,id).Scan(&user.ID,&user.Name); err != nil {
		if err == sql.ErrNoRows {
			return domain.User{},domain.ErrUserNotFound
		}
		return domain.User{},err
	 }
	 return user,nil
}