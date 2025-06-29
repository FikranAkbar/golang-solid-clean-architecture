package infrastructures

import (
	"database/sql"
	"errors"
	"golang-solid-clean-architecture/database"
	"golang-solid-clean-architecture/entities"
)

type UserRepository struct {
	db database.Database
}

func NewUserRepository(db database.Database) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *entities.User) error {
	err := r.db.Execute(`INSERT INTO users (username, email) VALUES ($1, $2)`, user.Name, user.Email)
	return err
}

func (r *UserRepository) GetByID(id int64) (*entities.User, error) {
	row, err := r.db.QueryRow(`SELECT id, username, email FROM users WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}

	user := &entities.User{}
	row.Scan(&user.ID, &user.Name, &user.Email)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil // User not found
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}
