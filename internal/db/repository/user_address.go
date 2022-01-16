package repository

import "database/sql"

type UsersAddressRepository struct {
	DB *sql.DB
	TX *sql.Tx
}

type UserAddress struct {
	Id      int
	Country string
	City    string
	Address string
	UserId  int
}

func NewUserAddressRepository(conn *sql.DB) UsersAddressRepository {
	return UsersAddressRepository{
		DB: conn,
	}
}

func (r UsersRepository) AddUserAddress(u UserAddress) (int, error) {
	var id int

	if r.TX != nil {
		err := r.TX.QueryRow("INSERT INTO user_address (country, city, address, user_id) VALUES(?, ?, ?) RETURNING id", u.Country, u.City, u.Address, u.UserId).Scan(&id)
		return id, err
	}
	err := r.DB.QueryRow("INSERT INTO user_address (country, city, address, user_id) VALUES(?, ?, ?) RETURNING id", u.Country, u.City, u.Address, u.UserId).Scan(&id)

	return id, err
}

func (r *UsersAddressRepository) BeginTx() error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	r.TX = tx
	return nil
}

func (r *UsersAddressRepository) CommitTx() error {
	defer func() {
		r.TX = nil
	}()
	if r.TX != nil {
		return r.CommitTx()
	}
	return nil
}

func (r *UsersAddressRepository) RollbackTx() error {
	defer func() {
		r.TX = nil
	}()
	if r.TX != nil {
		return r.RollbackTx()
	}
	return nil
}

func (r *UsersAddressRepository) GetTx() *sql.Tx {
	return r.TX
}

func (r *UsersAddressRepository) SetTx(tx *sql.Tx) {
	r.TX = tx
}
