package shared

import (
	"time"
)

type User struct {
	Id             string    `json:"id" sql:"id" db:"id"`
	Email          string    `json:"email" sql:"email" db:"email"`
	Password       string    `json:"password" validate:"required" sql:"password" db:"password"`
	Username       string    `json:"username" validate:"required" sql:"username" db:"username"`
	TokenHash      string    `json:"tokenhash" sql:"tokenhash" db:"tokenhash"`
	CreatedAt      time.Time `json:"created_at" sql:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" sql:"updated_at" db:"updated_at"`
	FollowerCount  int       `json:"follower_count" sql:"follower_count" db:"follower_count"`
	FollowingCount int       `json:"following_count" sql:"following_count" db:"following_count"`
} //@Name User

type GenericResponse struct {
	Status     bool        `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	StatusCode int         `json:"status_code"`
} //@Name GenericResponse

type UserKey struct{}
type UserIDKey struct{}
