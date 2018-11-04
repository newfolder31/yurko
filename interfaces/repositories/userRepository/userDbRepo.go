package userRepository

import "github.com/newfolder31/yurko/usecases/userUsecases"

type UserInDbRepo struct {
	DbHandler SqlxHandler
}

func (repo UserInDbRepo) Store(user *userUsecases.User) error {
	err := repo.DbHandler.Execute("INSERT INTO usr (email, password, is_admin, is_active, first_name, last_name, fathers_name) "+
		"VALUES (:email, :password, :is_admin, :is_active, :first_name, :last_name, :fathers_name)", &user)
	return err
}

func (repo UserInDbRepo) FindById(id int) (*userUsecases.User, error) {
	user := userUsecases.User{}
	err := repo.DbHandler.Get(user, "SELECT * FROM usr WHERE id=$1 LIMIT 1", id)
	return &user, err
}

func (repo UserInDbRepo) FindByEmail(email string) (*userUsecases.User, error) {
	user := userUsecases.User{}

	row := repo.DbHandler.QueryRowx(user, "SELECT id, email, password, is_admin, is_active, first_name, last_name, fathers_name  "+
		"FROM usr WHERE email=$1 LIMIT 1", email)
	err := row.StructScan(&user)
	return &user, err
}

func (repo UserInDbRepo) FindByEmail2(email string) (*userUsecases.User, error) { //todo: remain only one variant to get struct instance from one row
	user := userUsecases.User{}
	err := repo.DbHandler.Get(user, "SELECT id, email, password, is_admin, is_active, first_name, last_name, fathers_name  "+
		"FROM usr WHERE email=$1 LIMIT 1", email)
	return &user, err
}

func (repo UserInDbRepo) FindByEmailAndPassword(email, password string) (*userUsecases.User, error) {
	user := userUsecases.User{}
	err := repo.DbHandler.Get(user, "SELECT * FROM usr WHERE email=$1 and password=$2 LIMIT 1", email, password)
	return &user, err
}
