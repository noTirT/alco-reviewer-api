package userinfo

import "time"

type UsernameChangeRequest struct {
	NewUsername string `json:"new_username"`
}

type UserProfileResponse struct {
    Email string `json:"email"`
    Username string `json:"username"`
    CreatedAt time.Time `json:"created_at"`
}
