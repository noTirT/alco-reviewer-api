package user

import "time"

type UsernameChangeRequest struct {
	NewUsername string `json:"new_username"`
}

type UserProfileResponse struct {
	Email          string    `json:"email"`
	Username       string    `json:"username"`
	CreatedAt      time.Time `json:"created_at"`
	FollowerCount  int       `json:"follower_count"`
	FollowingCount int       `json:"following_count"`
}

type UserProfileFollowingResponse struct {
	Email          string    `json:"email"`
	Username       string    `json:"username"`
	CreatedAt      time.Time `json:"created_at"`
	FollowerCount  int       `json:"follower_count"`
	FollowingCount int       `json:"following_count"`
	Following      bool      `json:"following"`
}
