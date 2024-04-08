package userinfo

type UserinfoService interface {
    ChangeUsername(userID string, request *UsernameChangeRequest) error
    GetUserByID(userID string) (UserProfileResponse, error)
}

type userinfoService struct{
    repo UserinfoRepository
}

func NewUserinfoService(repo UserinfoRepository) UserinfoService{
    return &userinfoService{
        repo: repo,
    }
}

func (service *userinfoService) ChangeUsername(userID string, request *UsernameChangeRequest) error {
    err := service.repo.UpdateUsername(userID, request.NewUsername)

    return err
}

func (service *userinfoService) GetUserByID(userID string) (UserProfileResponse, error) {
    curUser, err := service.repo.GetUserById(userID)

    return UserProfileResponse{
        Email: curUser.Email,
        Username: curUser.Username,
        CreatedAt: curUser.CreatedAt,
    }, err
}
