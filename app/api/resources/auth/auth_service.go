package auth

import (
	"errors"
	"fmt"
	"log"
	"time"

	"noTirT/alcotracker/app/api/shared"
	"noTirT/alcotracker/configs"
	"noTirT/alcotracker/util"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	CreateUser(request *shared.User) (*shared.User, error)
	GetUserByUsername(request *shared.User) (*shared.User, error)
	GetUserByID(userID string) (*shared.User, error)
	Authenticate(reqUser *shared.User, compUser *shared.User) bool
	GenerateAccessToken(user *shared.User) (string, error)
	GenerateRefreshToken(user *shared.User) (string, error)
	ValidateAccessToken(tokenString string) (string, error)
	ValidateRefreshToken(tokenString string) (string, string, error)
}

type authService struct {
	repo   AuthRepository
	config *configs.Configuration
}

type AccessTokenCustomClaims struct {
	UserID  string
	KeyType string
	jwt.StandardClaims
}

type RefreshTokenCustomClaims struct {
	UserID    string
	CustomKey string
	KeyType   string
	jwt.StandardClaims
}

func NewAuthService(repo AuthRepository, config *configs.Configuration) AuthService {
	return &authService{
		repo:   repo,
		config: config,
	}
}

func (service *authService) CreateUser(request *shared.User) (*shared.User, error) {
	user, err := service.repo.CreateUser(&shared.User{
		Id:        "",
		Email:     request.Email,
		Username:  request.Username,
		Password:  request.Password,
		TokenHash: util.GenerateRandomString(15),
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (service *authService) GetUserByUsername(request *shared.User) (*shared.User, error) {
	user := service.repo.GetUserByUsername(request.Username)

	return user, nil
}

func (service *authService) GetUserByID(userID string) (*shared.User, error) {
	user := service.repo.GetUserByID(userID)

	return user, nil
}

func (service *authService) Authenticate(reqUser *shared.User, compUser *shared.User) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(compUser.Password), []byte(reqUser.Password)); err != nil {
		log.Println("Password hashes are not the same")
		return false
	}
	return true
}

func (service *authService) GenerateAccessToken(user *shared.User) (string, error) {
	userId := fmt.Sprint(user.Id)
	tokenType := "access"

	claims := AccessTokenCustomClaims{
		userId,
		tokenType,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(service.config.JwtExpiration)).Unix(),
			Issuer:    "alcohol.user.service",
		},
	}

	privateKey := util.PrivateRSAFromFile("access")
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	return token.SignedString(privateKey)
}

func (service *authService) GenerateRefreshToken(user *shared.User) (string, error) {
	cusKey := util.GenerateCustomKey(fmt.Sprint(user.Id), user.TokenHash)
	tokenType := "refresh"

	claims := RefreshTokenCustomClaims{
		fmt.Sprint(user.Id),
		cusKey,
		tokenType,
		jwt.StandardClaims{
			Issuer: "alcohol.user.service",
		},
	}

	privateKey := util.PrivateRSAFromFile("refresh")
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	return token.SignedString(privateKey)
}

func (service *authService) ValidateAccessToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AccessTokenCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			log.Println("Unexpected signing method in access token")
			return nil, errors.New("Unexpected signing method in access token")
		}
		return util.PublicRSAFromFile("access"), nil
	})

	if err != nil {
		log.Println(err)
		return "", err
	}

	claims, ok := token.Claims.(*AccessTokenCustomClaims)
	if !ok || !token.Valid || claims.UserID == "" || claims.KeyType != "access" {
		return "", errors.New("Invalid token: authentication failed")
	}
	return claims.UserID, nil
}

func (service *authService) ValidateRefreshToken(tokenString string) (string, string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &RefreshTokenCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			log.Println("Unexpected signing method in auth token")
			return nil, errors.New("Unexpected signing method in auth token")
		}
		return util.PublicRSAFromFile("refresh"), nil
	})

	if err != nil {
		log.Println(err)
		return "", "", err
	}

	claims, ok := token.Claims.(*RefreshTokenCustomClaims)
	if !ok || !token.Valid || claims.UserID == "" || claims.KeyType != "refresh" {
		return "", "", errors.New("invalid token: authentication failed")
	}
	return claims.UserID, claims.CustomKey, nil
}
