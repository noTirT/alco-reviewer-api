package user

type UserService interface {
	ChangeUsername(userID string, request *UsernameChangeRequest) error
	GetUserByID(userID string, ownUserID string) (UserProfileFollowingResponse, error)
	FollowUser(userID string, followId string) error
	UnfollowUser(userID string, unfollowId string) error
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (service *userService) FollowUser(userID string, followId string) error {
	err := service.repo.FollowUser(userID, followId)
	if err != nil {
		return err
	}
	err = service.repo.IncrementFollowingCount(userID)
	if err != nil {
		return err
	}

	err = service.repo.IncrementFollowerCount(followId)

	return err
}

func (service *userService) UnfollowUser(userID string, unfollowId string) error {
	err := service.repo.UnfollowUser(userID, unfollowId)
	if err != nil {
		return err
	}

	err = service.repo.DecrementFollowingCount(userID)
	if err != nil {
		return err
	}

	err = service.repo.DecrementFollowerCount(unfollowId)

	return err
}

func (service *userService) ChangeUsername(userID string, request *UsernameChangeRequest) error {
	err := service.repo.UpdateUsername(userID, request.NewUsername)

	return err
}

func (service *userService) GetUserByID(userID string, ownUserID string) (UserProfileFollowingResponse, error) {
	curUser, err := service.repo.GetUserById(userID)
	exists, err := service.repo.GetFollowingUser(userID, ownUserID)

	return UserProfileFollowingResponse{
		Email:          curUser.Email,
		Username:       curUser.Username,
		CreatedAt:      curUser.CreatedAt,
		FollowerCount:  curUser.FollowerCount,
		FollowingCount: curUser.FollowingCount,
		Following:      exists,
	}, err
}
