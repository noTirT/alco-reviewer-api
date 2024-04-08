package userinfo

import (
	"noTirT/alcotracker/app/api/shared"
	"noTirT/alcotracker/app/db"
	"time"

	"github.com/blockloop/scan/v2"
)

type UserinfoRepository interface{
    UpdateUsername(userID string, newUserName string) error
    GetUserById(userID string) (*shared.User, error)
}

type userinfoRepository struct{
    db *db.PostgresDB
}

func NewUserinfoRepository(db *db.PostgresDB) UserinfoRepository{
    return &userinfoRepository{
        db: db,
    }
}

func (repo *userinfoRepository) UpdateUsername(userID string, newUserName string) error {
    _, err := repo.db.Db.Exec("UPDATE users SET username=$1, updated_at=$2 WHERE id=$3", newUserName, time.Now(), userID)

    return err
}

func (repo *userinfoRepository) GetUserById(userID string) (*shared.User, error) {
    row, err := repo.db.Db.Query("SELECT * FROM users WHERE id=$1", userID)
    if err != nil{
        return nil, err
    }

    var user shared.User
    err = scan.Row(&user, row)
    return &user, err
}
