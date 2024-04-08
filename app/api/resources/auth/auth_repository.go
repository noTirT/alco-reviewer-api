package auth

import (
	"errors"
	"github.com/blockloop/scan/v2"
	"noTirT/alcotracker/app/db"
	"noTirT/alcotracker/app/api/shared"
	"strings"
)

type AuthRepository interface {
	CreateUser(user *shared.User) (*shared.User, error)
	GetUserByEmail(userEmail string) *shared.User
	GetUserByUsername(username string) *shared.User
	GetUserByID(userID string) *shared.User
}

type authRepository struct {
	db *db.PostgresDB
}

func NewAuthRepository(db *db.PostgresDB) AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (repo *authRepository) GetUserByID(userID string) *shared.User {
	row, _ := repo.db.Db.Query(`SELECT * FROM users WHERE id=$1;`, userID)
	var user shared.User
	err := scan.Row(&user, row)
	if err != nil {
		panic(err)
	}
	return &user
}

func (repo *authRepository) GetUserByEmail(userEmail string) *shared.User {
	row, _ := repo.db.Db.Query(`SELECT * FROM users WHERE email=$1;`, userEmail)
	var user shared.User
	err := scan.Row(&user, row)
	if err != nil {
		panic(err)
	}
	return &user
}

func (repo *authRepository) CreateUser(user *shared.User) (*shared.User, error) {
	_, err := repo.db.Db.Exec(`INSERT INTO users(email, username, password, tokenhash) VALUES($1, $2, $3, $4)`, user.Email, user.Username, user.Password, user.TokenHash)

	if err != nil {
		if strings.Contains(err.Error(), "violates unique constraint") {
			err = errors.New("shared.User with that email or username already exists")
		}
	}

	return user, err
}

func (repo *authRepository) GetUserByUsername(username string) *shared.User {
	row, _ := repo.db.Db.Query(`SELECT * FROM users WHERE username=$1;`, username)
	var user shared.User
	err := scan.Row(&user, row)
	if err != nil {
		panic(err)
	}
	return &user
}
