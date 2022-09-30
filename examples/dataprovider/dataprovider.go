package dataprovider

import "database/sql"

type User struct {
	ID    int
	Name  string
	Email string
}

type Users interface {
	ListUsers() ([]User, error)
	DeleteUser(id int) error
	AddUser(name, email string) (User, error)
}

type UsersDatabase struct {
	db *sql.DB
}

func (ud *UsersDatabase) ListUsers() ([]User, error) {
	/* omitted */

	return nil, nil
}
func (ud *UsersDatabase) DeleteUser(id int) error {
	/* omitted */
	return nil
}

func (ud *UsersDatabase) AddUser(name, email string) (User, error) {
	return User{}, nil
}

func NewUsersDatabase(db *sql.DB) Users {
	return &UsersDatabase{db: db}
}
