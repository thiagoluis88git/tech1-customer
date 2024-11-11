package handler_test

import (
	"context"
	"os"

	"github.com/stretchr/testify/mock"
	"github.com/thiagoluis88git/tech1-customer/internal/core/domain/dto"
	"github.com/thiagoluis88git/tech1-customer/pkg/environment"
)

func setup() {
	os.Setenv(environment.QRCodeGatewayRootURL, "ROOT_URL")
	os.Setenv(environment.DBHost, "HOST")
	os.Setenv(environment.DBPort, "1234")
	os.Setenv(environment.DBUser, "User")
	os.Setenv(environment.DBPassword, "Pass")
	os.Setenv(environment.DBName, "Name")
	os.Setenv(environment.CognitoClientID, "ClienId")
	os.Setenv(environment.CognitoGroupAdmin, "Admin")
	os.Setenv(environment.CognitoGroupUser, "CognitoUser")
	os.Setenv(environment.CognitoUserPoolID, "USerPool")
	os.Setenv(environment.WebhookMercadoLivrePaymentURL, "WEBHOOK")
	os.Setenv(environment.QRCodeGatewayToken, "token")
	os.Setenv(environment.Region, "Region")
}

func mockCreateUserForm() dto.UserAdmin {
	return dto.UserAdmin{
		Name:  "Name",
		CPF:   "12345678910",
		Email: "teste@email.com",
	}
}

func mockUpdateUserForm() dto.UserAdmin {
	return dto.UserAdmin{
		ID:    uint(3),
		Name:  "Name",
		CPF:   "12345678910",
		Email: "teste@email.com",
	}
}

func mockGetUserByCPF() dto.UserAdminForm {
	return dto.UserAdminForm{
		CPF: "12345678910",
	}
}

type MockCreateCustomerUseCase struct {
	mock.Mock
}

type MockUpdateCustomerUseCase struct {
	mock.Mock
}

type MockGetCustomerByIdUseCase struct {
	mock.Mock
}

type MockGetCustomerByCPFUseCase struct {
	mock.Mock
}

type MockLoginCustomerUseCase struct {
	mock.Mock
}

type MockLoginUnknownCustomerUseCase struct {
	mock.Mock
}

type MockPayOrderUseCase struct {
	mock.Mock
}

type MockGetPaymentTypesUseCase struct {
	mock.Mock
}

type MockCreateOrderUseCase struct {
	mock.Mock
}

type MockGetOrderByIdUseCase struct {
	mock.Mock
}

type MockGetOrdersToPrepareUseCase struct {
	mock.Mock
}

type MockGetOrdersToFollowUseCase struct {
	mock.Mock
}

type MockGetOrdersWaitingPaymentUseCase struct {
	mock.Mock
}

type MockUpdateToPreparingUseCase struct {
	mock.Mock
}

type MockUpdateToDoneUseCase struct {
	mock.Mock
}

type MockUpdateToDeliveredUseCase struct {
	mock.Mock
}

type MockUpdateToNotDeliveredUseCase struct {
	mock.Mock
}

type MockCreateProductUseCase struct {
	mock.Mock
}

type MockGetProductsByCategoryUseCase struct {
	mock.Mock
}

type MockGetProductsByIDUseCase struct {
	mock.Mock
}

type MockDeleteProductUseCase struct {
	mock.Mock
}

type MockUpdateProductUseCase struct {
	mock.Mock
}

type MockGenerateQRCodePaymentUseCase struct {
	mock.Mock
}

type MockCreateUserUseCase struct {
	mock.Mock
}

type MockUpdateUserUseCase struct {
	mock.Mock
}

type MockGetUserByIdUseCase struct {
	mock.Mock
}

type MockGetUserByCPFUseCase struct {
	mock.Mock
}

type MockGetCategoryUseCase struct {
	mock.Mock
}

type MockLoginUserUseCase struct {
	mock.Mock
}

func (mock *MockCreateCustomerUseCase) Execute(ctx context.Context, customer dto.Customer) (dto.CustomerResponse, error) {
	args := mock.Called(ctx, customer)
	err := args.Error(1)

	if err != nil {
		return dto.CustomerResponse{}, err
	}

	return args.Get(0).(dto.CustomerResponse), nil
}

func (mock *MockUpdateCustomerUseCase) Execute(ctx context.Context, customer dto.Customer) error {
	args := mock.Called(ctx, customer)
	err := args.Error(0)

	if err != nil {
		return err
	}

	return nil
}

func (mock *MockGetCustomerByIdUseCase) Execute(ctx context.Context, id uint) (dto.Customer, error) {
	args := mock.Called(ctx, id)
	err := args.Error(1)

	if err != nil {
		return dto.Customer{}, err
	}

	return args.Get(0).(dto.Customer), nil
}

func (mock *MockGetCustomerByCPFUseCase) Execute(ctx context.Context, cpf string) (dto.Customer, error) {
	args := mock.Called(ctx, cpf)
	err := args.Error(1)

	if err != nil {
		return dto.Customer{}, err
	}

	return args.Get(0).(dto.Customer), nil
}

func (mock *MockLoginCustomerUseCase) Execute(ctx context.Context, cpf string) (dto.Token, error) {
	args := mock.Called(ctx, cpf)
	err := args.Error(1)

	if err != nil {
		return dto.Token{}, err
	}

	return args.Get(0).(dto.Token), nil
}

func (mock *MockLoginUnknownCustomerUseCase) Execute(ctx context.Context) (dto.Token, error) {
	args := mock.Called(ctx)
	err := args.Error(1)

	if err != nil {
		return dto.Token{}, err
	}

	return args.Get(0).(dto.Token), nil
}

func (mock *MockCreateUserUseCase) Execute(ctx context.Context, user dto.UserAdmin) (dto.UserAdminResponse, error) {
	args := mock.Called(ctx, user)
	err := args.Error(1)

	if err != nil {
		return dto.UserAdminResponse{}, err
	}

	return args.Get(0).(dto.UserAdminResponse), nil
}

func (mock *MockGetUserByIdUseCase) Execute(ctx context.Context, id uint) (dto.UserAdmin, error) {
	args := mock.Called(ctx, id)
	err := args.Error(1)

	if err != nil {
		return dto.UserAdmin{}, err
	}

	return args.Get(0).(dto.UserAdmin), nil
}

func (mock *MockGetUserByCPFUseCase) Execute(ctx context.Context, cpf string) (dto.UserAdmin, error) {
	args := mock.Called(ctx, cpf)
	err := args.Error(1)

	if err != nil {
		return dto.UserAdmin{}, err
	}

	return args.Get(0).(dto.UserAdmin), nil
}

func (mock *MockLoginUserUseCase) Execute(ctx context.Context, cpf string) (dto.Token, error) {
	args := mock.Called(ctx, cpf)
	err := args.Error(1)

	if err != nil {
		return dto.Token{}, err
	}

	return args.Get(0).(dto.Token), nil
}

func (mock *MockUpdateUserUseCase) Execute(ctx context.Context, user dto.UserAdmin) error {
	args := mock.Called(ctx, user)
	err := args.Error(0)

	if err != nil {
		return err
	}

	return nil
}
