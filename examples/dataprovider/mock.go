package dataprovider

type UsersMock struct {
	MockedListUsers  func() ([]User, error)
	MockedDeleteUser func(id int) error
	MockedAddUser    func(name, email string) (User, error)
}

func (um *UsersMock) ListUsers() ([]User, error) {
	return um.MockedListUsers()
}

func (um *UsersMock) DeleteUser(id int) error {
	return um.MockedDeleteUser(id)
}

func (um *UsersMock) AddUser(name, email string) (User, error) {
	return um.MockedAddUser(name, email)
}
