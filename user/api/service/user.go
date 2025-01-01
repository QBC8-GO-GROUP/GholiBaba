package service

import (
	"context"
	"errors"

	"github.com/QBC8-GO-GROUP/GholiBaba/api/pb"
	"github.com/QBC8-GO-GROUP/GholiBaba/internal/user"
	"github.com/QBC8-GO-GROUP/GholiBaba/internal/user/domain"
	userPort "github.com/QBC8-GO-GROUP/GholiBaba/internal/user/port"
	"github.com/QBC8-GO-GROUP/GholiBaba/pkg/jwt"
	"github.com/QBC8-GO-GROUP/GholiBaba/pkg/time"

	jwt2 "github.com/golang-jwt/jwt/v5"
)

type UserService struct {
	svc                   userPort.Service
	authSecret            string
	expMin, refreshExpMin uint
}

func NewUserService(svc userPort.Service, authSecret string, expMin, refreshExpMin uint) *UserService {
	return &UserService{
		svc:           svc,
		authSecret:    authSecret,
		expMin:        expMin,
		refreshExpMin: refreshExpMin,
	}
}

var (
	ErrUserCreationValidation = user.ErrUserCreationValidation
	ErrUserOnCreate           = user.ErrUserOnCreate
	ErrUserNotFound           = user.ErrUserNotFound
	ErrInvalidUserPassword    = errors.New("invalid password")
	ErrWrongOTP               = errors.New("wrong otp")
)

func (s *UserService) SignUp(ctx context.Context, req *pb.UserSignUpRequest) (*pb.UserSignUpResponse, error) {

	if req.GetRole() == pb.Role_ROLE_REGULAR_USER {
		req.Role = pb.Role_ROLE_REGULAR_USER
	}
	userID, err := s.svc.CreateUser(ctx, domain.User{
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
		Phone:     domain.Phone(req.GetPhone()),
		Password:  req.GetPassword(),
		Role:      domain.Role(req.GetRole().String()),
	})

	if err != nil {
		return nil, err
	}

	accessToken, err := jwt.CreateToken([]byte(s.authSecret), &jwt.UserClaims{
		RegisteredClaims: jwt2.RegisteredClaims{
			ExpiresAt: jwt2.NewNumericDate(time.AddMinutes(s.expMin, true)),
		},
		UserID: uint(userID),
	})
	if err != nil {
		return nil, err
	}

	refreshToken, err := jwt.CreateToken([]byte(s.authSecret), &jwt.UserClaims{
		RegisteredClaims: jwt2.RegisteredClaims{
			ExpiresAt: jwt2.NewNumericDate(time.AddMinutes(s.expMin, true)),
		},
		UserID: uint(userID),
	})
	if err != nil {
		return nil, err
	}

	return &pb.UserSignUpResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *UserService) SignIn(ctx context.Context, req *pb.UserSignInRequest) (*pb.UserSignInResponse, error) {

	user, err := s.svc.GetUserByFilter(ctx, &domain.UserFilter{
		Phone: req.GetPhone(),
	})
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrUserNotFound
	}

	if !user.PasswordIsCorrect(req.GetPassword()) {
		return nil, ErrInvalidUserPassword
	}

	access, refresh, err := s.createTokens(uint(user.ID), string(user.Role))
	if err != nil {
		return nil, err
	}

	return &pb.UserSignInResponse{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

// func UpdateUserRoleHandler(ctx context.Context, req *pb.UserSignUpRequest) (*pb.UserSignUpResponse, error){
// 	return nil, errors
// }

func (s *UserService) createTokens(userID uint, role string) (access, refresh string, err error) {
	access, err = jwt.CreateToken([]byte(s.authSecret), &jwt.UserClaims{
		RegisteredClaims: jwt2.RegisteredClaims{
			ExpiresAt: jwt2.NewNumericDate(time.AddMinutes(s.expMin, true)),
		},
		UserID: uint(userID),
		Role:   role,
	})
	if err != nil {
		return
	}

	refresh, err = jwt.CreateToken([]byte(s.authSecret), &jwt.UserClaims{
		RegisteredClaims: jwt2.RegisteredClaims{
			ExpiresAt: jwt2.NewNumericDate(time.AddMinutes(s.refreshExpMin, true)),
		},
		UserID: uint(userID),
		Role:   role,
	})

	if err != nil {
		return
	}

	return
}
