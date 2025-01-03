package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"

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
	ErrInvalidRole            = errors.New("invalid role provided")
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

// func (s *UserService) UpdateUserRoleHandler(ctx context.Context, req *pb.ChangeRoleRequest) error {

// 	if req.Role.String() == string(domain.Admin) {
// 		return ErrInvalidRole
// 	}

// 	requestedRole := domain.Role(req.Role)

// 	if !domain.IsValidRole(requestedRole) {
// 		return ErrInvalidRole
// 	}

// 	requesterID, ok := ctx.Value("user_id").(string)
// 	if !ok || requesterID == "" {
// 		return errors.New("failed to retrieve requester ID from context")
// 	}

// 	user, err := s.svc.GetUserById(ctx, req.UserId)
// 	if err != nil {
// 		return ErrUserNotFound
// 	}

// 	user.Role = requestedRole

// 	if err := s.svc.UpdateUser(ctx, user); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (s *UserService) UpdateUserRoleHandler(ctx context.Context, req *pb.ChangeRoleRequest) error {

// 	if req.Role.String() == string(domain.Admin) {
// 		return ErrInvalidRole
// 	}

// 	requestedRole := domain.Role(req.Role)
// 	// if !domain.IsValidRole(requestedRole) {

// 	// 	return ErrInvalidRole
// 	// }

// 	fmt.Println("yeeeeeees", requestedRole)

// 	userIDStr, ok := ctx.Value("user_id").(string)
// 	if !ok || userIDStr == "" {
// 		return errors.New("failed to retrieve user ID from context")
// 	}

// 	// Convert userIDStr to domain.UserID (assuming domain.UserID is uint)
// 	userIDUint, err := strconv.ParseUint(userIDStr, 10, 64)
// 	if err != nil {
// 		return fmt.Errorf("invalid user ID format: %w", err)
// 	}

// 	userID := domain.UserID(userIDUint)

// 	// Fetch the user by ID
// 	user, err := s.svc.GetUserById(ctx, userID)
// 	if err != nil {
// 		return ErrUserNotFound
// 	}

// 	// Update the user's role
// 	user.Role = requestedRole
// 	if err := s.svc.UpdateUser(ctx, user); err != nil {
// 		return fmt.Errorf("failed to update user role: %w", err)
// 	}

// 	return nil
// }

// func (s *UserService) UpdateUserRoleHandler(ctx context.Context, req *pb.ChangeRoleRequest) error {

// 	requestedRole, err := domain.MapProtoRoleToDomain(req.Role)
// 	if err != nil {
// 		return ErrInvalidRole
// 	}

// 	fmt.Println("Parsed Role:", requestedRole)

// 	u := ctx.Value("user_id")
// 	fmt.Println("User ID:", u)

// 	userIDStr, ok := ctx.Value("user_id").(string)

// 	if !ok || userIDStr == "" {
// 		return errors.New("failed to retrieve user ID from context")
// 	}

// 	userIDUint, err := strconv.ParseUint(userIDStr, 10, 64)
// 	if err != nil {
// 		return fmt.Errorf("invalid user ID format: %w", err)
// 	}

// 	userID := domain.UserID(userIDUint)

// 	// Fetch the user by ID
// 	user, err := s.svc.GetUserById(ctx, userID)
// 	if err != nil {
// 		return ErrUserNotFound
// 	}

// 	// Update the user's role
// 	user.Role = requestedRole
// 	if err := s.svc.UpdateUser(ctx, user); err != nil {
// 		return fmt.Errorf("failed to update user role: %w", err)
// 	}

// 	return nil
// }

func (s *UserService) UpdateUserRoleHandler(ctx context.Context, req *pb.ChangeRoleRequest) error {
	requestedRole, err := domain.MapProtoRoleToDomain(req.Role)
	if err != nil {
		return ErrInvalidRole
	}

	fmt.Println("Parsed Role:", requestedRole)

	// Retrieve user_id from context
	u := ctx.Value("user_id")
	fmt.Println("Raw User ID from context:", u)

	var userIDUint uint64

	switch v := u.(type) {
	case string:
		// If it's a string, try parsing it into uint64
		userIDUint, err = strconv.ParseUint(v, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid user ID format from string: %w", err)
		}
	case int:
		// If it's an int, cast directly
		userIDUint = uint64(v)
	case uint:
		// If it's a uint, cast directly
		userIDUint = uint64(v)
	case uint64:
		// If it's already uint64, use it directly
		userIDUint = v
	default:
		return errors.New("user_id is of an unsupported type in context")
	}

	// Convert to your domain-specific type
	userID := domain.UserID(userIDUint)

	// Fetch the user by ID
	user, err := s.svc.GetUserById(ctx, userID)
	if err != nil {
		return ErrUserNotFound
	}
	fmt.Println("requestedRole", requestedRole)

	// Update the user's role
	user.Role = requestedRole
	if err := s.svc.UpdateUser(ctx, user); err != nil {
		return fmt.Errorf("failed to update user role: %w", err)
	}

	return nil
}

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
