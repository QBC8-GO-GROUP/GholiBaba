package service

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/QBC8-GO-GROUP/GholiBaba/api/pb"
	"github.com/QBC8-GO-GROUP/GholiBaba/internal/company"
	"github.com/QBC8-GO-GROUP/GholiBaba/internal/company/domain"
	companyPort "github.com/QBC8-GO-GROUP/GholiBaba/internal/company/port"
	"github.com/QBC8-GO-GROUP/GholiBaba/pkg/adapters/storage/types"
	"github.com/QBC8-GO-GROUP/GholiBaba/pkg/logger"
)

var (
	ErrPasswordNotMatch = errors.New("not match password")
)

type CompanyService struct {
	svc                   companyPort.Service
	authSecret            string
	expMin, refreshExpMin uint
}

func NewCompanyService(svc companyPort.Service, authSecret string, expMin, refreshExpMin uint) *CompanyService {
	return &CompanyService{
		svc:           svc,
		authSecret:    authSecret,
		expMin:        expMin,
		refreshExpMin: refreshExpMin,
	}
}

var (
	ErrCompanyCreationValidation = company.ErrCompanyCreationValidation
	ErrCompanyOnCreate           = company.ErrCompanyOnCreate
	ErrCompanyNotFound           = company.ErrCompanyNotFound
)

type SignUpFirstResponseWrapper struct {
	RequestTimestamp int64                          `json:"requestTimestamp"`
	Data             *pb.CompanySignUpFirstResponse `json:"data"`
}
type SignUpSecondResponseWrapper struct {
	RequestTimestamp int64                           `json:"requestTimestamp"`
	Data             *pb.CompanySignUpSecondResponse `json:"data"`
}

func (s *CompanyService) SignUp(ctx context.Context, req *pb.CompanySignUpFirstRequest) (*SignUpFirstResponseWrapper, error) {
	companyID, err := s.svc.CreateCompany(ctx, domain.Company{
		FirstName:    req.GetFirstName(),
		LastName:     req.GetLastName(),
		Phone:        domain.Phone(req.GetPhone()),
		Email:        domain.Email(req.GetEmail()),
		PasswordHash: req.GetPassword(),
		NationalCode: domain.NationalCode(req.GetNationalCode()),
		City:         req.GetCity(),
		Gender:       req.GetGender(),
	})

	if err != nil {
		return nil, err
	}

	response := &SignUpFirstResponseWrapper{
		RequestTimestamp: time.Now().Unix(),
		Data: &pb.CompanySignUpFirstResponse{
			CompanyId: uint64(companyID),
		},
	}

	return response, nil
}

func (s *CompanyService) GetByID(ctx context.Context, id uint) (*pb.Company, error) {
	company, err := s.svc.GetCompanyByID(ctx, domain.CompanyID(id))
	if err != nil {
		return nil, err
	}

	return &pb.Company{
		Id:                uint64(company.ID),
		FirstName:         company.FirstName,
		LastName:          company.LastName,
		Phone:             string(company.Phone),
		Email:             string(company.Email),
		PasswordHash:      company.PasswordHash,
		NationalCode:      string(company.NationalCode),
		City:              company.City,
		Gender:            company.Gender,
		SurveyLimitNumber: int32(company.SurveyLimitNumber), // Protobuf may require int32 instead of int
		Balance:           int32(company.Balance),
	}, nil
}

func (s *CompanyService) Update(ctx context.Context, company *types.Company) error {
	err := s.svc.UpdateCompany(ctx, domain.Company{
		ID:        domain.CompanyID(company.ID),
		FirstName: company.FirstName,
		LastName:  company.LastName,
		Phone:     domain.Phone(company.Phone),
		CreatedAt: company.CreatedAt,
		UpdatedAt: company.UpdatedAt,
	})
	if err != nil {
		logger.Error("update company error", nil)
		return err
	}
	return nil
}

func (s *CompanyService) DeleteByID(ctx context.Context, companyID int) error {
	err := s.svc.DeleteByID(ctx, domain.CompanyID(companyID))
	if err != nil {
		logger.Error("can not delete company", nil)
		return err
	}

	logger.Info("deleted company with id "+strconv.Itoa(int(companyID)), nil)
	return nil
}
