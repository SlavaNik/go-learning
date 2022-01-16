package services

import (
	"database/sql"
	"go_lect/db/repository"
)

type UserService struct {
	UserRepo        repository.UsersRepository
	UserAddressRepo repository.UsersAddressRepository
}

func NewUserService(conn *sql.DB) UserService {
	return UserService{
		UserRepo:        repository.NewUserRepository(conn),
		UserAddressRepo: repository.NewUserAddressRepository(conn),
	}
}

func (s UserService) CreateUser(user repository.User, address repository.UserAddress) error {
	err := s.UserRepo.BeginTx()
	if err != nil {
		return err
	}
	s.UserAddressRepo.SetTx(s.UserRepo.GetTx())
	defer func() {
		_ = s.UserRepo.RollbackTx()
	}()

	user_id, err := s.UserRepo.CreateUser(user)
	if err != nil {
		return err
	}
	address.UserId = user_id
	_, err = s.UserRepo.AddUserAddress(address)
	if err != nil {
		return err
	}

	err = s.UserRepo.CommitTx()
	if err != nil {
		return err
	}
	s.UserAddressRepo.SetTx(nil)
	return nil
}
