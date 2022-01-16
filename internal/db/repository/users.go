package repository

import "database/sql"

type UsersRepository struct {
	DB *sql.DB
	TX *sql.Tx
}

type User struct {
	Id           int
	Email        string
	PasswordHash string
	Name         string
}

func NewUserRepository(conn *sql.DB) UsersRepository {
	return UsersRepository{
		DB: conn,
	}
}

func (r UsersRepository) CreateUser(u User) (int, error) {
	var id int

	if r.TX != nil {
		err := r.TX.QueryRow("INSERT INTO users(name, email, password_hash) VALUES(?, ?, ?) RETURNING id", u.Name, u.Email, u.PasswordHash).Scan(&id)
		return id, err
	}
	err := r.DB.QueryRow("INSERT INTO users(name, email, password_hash) VALUES(?, ?, ?) RETURNING id", u.Name, u.Email, u.PasswordHash).Scan(&id)

	return id, err
}

func (r UsersRepository) GetByEmail(email string) (User, error) {
	var user User

	err := r.DB.QueryRow("SELECT id, email, name FROM users WHERE email = ?", email).Scan(&user.Id, &user.Email, &user.Name)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UsersRepository) BeginTx() error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	r.TX = tx
	return nil
}

func (r *UsersRepository) CommitTx() error {
	defer func() {
		r.TX = nil
	}()
	if r.TX != nil {
		return r.CommitTx()
	}
	return nil
}

func (r *UsersRepository) RollbackTx() error {
	defer func() {
		r.TX = nil
	}()
	if r.TX != nil {
		return r.RollbackTx()
	}
	return nil
}

func (r *UsersRepository) GetTx() *sql.Tx {
	return r.TX
}

func (r *UsersRepository) SetTx(tx *sql.Tx) {
	r.TX = tx
}
