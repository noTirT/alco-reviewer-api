package user

import (
	"noTirT/alcotracker/app/api/shared"
	"noTirT/alcotracker/app/db"
	"time"

	"github.com/blockloop/scan/v2"
)

type UserRepository interface {
	UpdateUsername(userID string, newUserName string) error
	GetUserById(userID string) (*shared.User, error)
	FollowUser(userID string, followedID string) error
	UnfollowUser(userID string, followedID string) error
	IncrementFollowerCount(userID string) error
	IncrementFollowingCount(userID string) error
	DecrementFollowerCount(userID string) error
	DecrementFollowingCount(userID string) error
	GetFollowingUser(userID string, questionUserID string) (bool, error)
}

type userRepository struct {
	db *db.PostgresDB
}

func NewUserRepository(db *db.PostgresDB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (repo *userRepository) DecrementFollowerCount(userID string) error {
	_, err := repo.db.Db.Exec("UPDATE users SET follower_count=follower_count-1 WHERE id=$1", userID)

	return err
}

func (repo *userRepository) DecrementFollowingCount(userID string) error {
	_, err := repo.db.Db.Exec("UPDATE users SET following_count=following_count-1 WHERE id=$1", userID)

	return err
}

func (repo *userRepository) IncrementFollowerCount(userID string) error {
	_, err := repo.db.Db.Exec("UPDATE users SET follower_count=follower_count+1 WHERE id=$1", userID)

	return err
}

func (repo *userRepository) IncrementFollowingCount(userID string) error {
	_, err := repo.db.Db.Exec("UPDATE users SET following_count=following_count+1 WHERE id=$1", userID)

	return err
}

func (repo *userRepository) UpdateUsername(userID string, newUserName string) error {
	_, err := repo.db.Db.Exec("UPDATE users SET username=$1, updated_at=$2 WHERE id=$3", newUserName, time.Now(), userID)

	return err
}

func (repo *userRepository) FollowUser(userID string, followedID string) error {
	_, err := repo.db.Db.Exec("INSERT INTO following(follower_id, followed_id) values($1, $2)", userID, followedID)

	return err
}

func (repo *userRepository) UnfollowUser(userID string, followedID string) error {
	_, err := repo.db.Db.Exec("DELETE FROM following WHERE follower_id=$1 AND followed_id=$2", userID, followedID)

	return err
}

func (repo *userRepository) GetUserById(userID string) (*shared.User, error) {
	row, err := repo.db.Db.Query("SELECT * FROM users WHERE id=$1", userID)
	if err != nil {
		return nil, err
	}

	var user shared.User
	err = scan.Row(&user, row)
	return &user, err
}

func (repo *userRepository) GetFollowingUser(userID string, questionUserID string) (bool, error) {
	rows, err := repo.db.Db.Query("SELECT * FROM following WHERE follower_id=$1 AND followed_id=$2", userID, questionUserID)
	if err != nil {
		return false, err
	}
	counter := 0
	for rows.Next() {
		counter++
	}

	return counter > 0, nil
}
