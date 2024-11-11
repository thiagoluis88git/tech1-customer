package usecases

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/thiagoluis88git/tech1-customer/internal/core/domain/dto"
)

type MockOrderRepository struct {
	mock.Mock
}

type MockCustomerRepository struct {
	mock.Mock
}

type MockPaymentRepository struct {
	mock.Mock
}

type MockPaymentGatewayRepository struct {
	mock.Mock
}

type MockProductRepository struct {
	mock.Mock
}

type MockUserAdminRepository struct {
	mock.Mock
}

type MockQRCodePaymentRepository struct {
	mock.Mock
}

func (mock *MockCustomerRepository) CreateCustomer(ctx context.Context, customer dto.Customer) (uint, error) {
	args := mock.Called(ctx, customer)
	err := args.Error(1)

	if err != nil {
		return 0, err
	}

	return args.Get(0).(uint), nil
}

func (mock *MockCustomerRepository) Login(ctx context.Context, cpf string) (string, error) {
	args := mock.Called(ctx, cpf)
	err := args.Error(1)

	if err != nil {
		return "", err
	}

	return args.Get(0).(string), nil
}

func (mock *MockCustomerRepository) LoginUnknown() (string, error) {
	args := mock.Called()
	err := args.Error(1)

	if err != nil {
		return "", err
	}

	return args.Get(0).(string), nil
}

func (mock *MockCustomerRepository) UpdateCustomer(ctx context.Context, customer dto.Customer) error {
	args := mock.Called(ctx, customer)
	err := args.Error(0)

	if err != nil {
		return err
	}

	return nil
}

func (mock *MockCustomerRepository) GetCustomerByCPF(ctx context.Context, cpf string) (dto.Customer, error) {
	args := mock.Called(ctx, cpf)
	err := args.Error(1)

	if err != nil {
		return dto.Customer{}, err
	}

	return args.Get(0).(dto.Customer), nil
}

func (mock *MockCustomerRepository) GetCustomerById(ctx context.Context, id uint) (dto.Customer, error) {
	args := mock.Called(ctx, id)
	err := args.Error(1)

	if err != nil {
		return dto.Customer{}, err
	}

	return args.Get(0).(dto.Customer), nil
}

func (mock *MockUserAdminRepository) CreateUser(ctx context.Context, customer dto.UserAdmin) (uint, error) {
	args := mock.Called(ctx, customer)
	err := args.Error(1)

	if err != nil {
		return uint(0), err
	}

	return args.Get(0).(uint), nil
}

func (mock *MockUserAdminRepository) GetUserById(ctx context.Context, id uint) (dto.UserAdmin, error) {
	args := mock.Called(ctx, id)
	err := args.Error(1)

	if err != nil {
		return dto.UserAdmin{}, err
	}

	return args.Get(0).(dto.UserAdmin), nil
}

func (mock *MockUserAdminRepository) GetUserByCPF(ctx context.Context, cpf string) (dto.UserAdmin, error) {
	args := mock.Called(ctx, cpf)
	err := args.Error(1)

	if err != nil {
		return dto.UserAdmin{}, err
	}

	return args.Get(0).(dto.UserAdmin), nil
}

func (mock *MockUserAdminRepository) Login(ctx context.Context, cpf string) (string, error) {
	args := mock.Called(ctx, cpf)
	err := args.Error(1)

	if err != nil {
		return "", err
	}

	return args.Get(0).(string), nil
}

func (mock *MockUserAdminRepository) UpdateUser(ctx context.Context, customer dto.UserAdmin) error {
	args := mock.Called(ctx, customer)
	err := args.Error(0)

	if err != nil {
		return err
	}

	return nil
}
