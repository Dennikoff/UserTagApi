package sqlstore

import "github.com/Dennikoff/UserTagApi/internal/app/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(user *model.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	if err := user.BeforeCreate(); err != nil {
		return err
	}

	if err := r.store.db.QueryRow("INSERT INTO users (id, email, password, nickname) VALUES (default, $1, $2, $3) RETURNING id",
		user.Email, user.EncryptedPassword, user.NickName,
	).Scan(&user.ID); err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	user := &model.User{}

	if err := r.store.db.QueryRow(
		"SELECT id, email, password, nickname from users where email=$1",
		email).Scan(
		&user.ID,
		&user.Email,
		&user.EncryptedPassword,
		&user.NickName); err != nil {
		return nil, err
	}

	return user, nil
}
