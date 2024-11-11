package handler_test

import (
	"context"
	"sync"

	"github.com/stretchr/testify/mock"
	"github.com/thiagoluis88git/tech1-customer/internal/core/domain/dto"
)

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

func (mock *MockPayOrderUseCase) Execute(ctx context.Context, payment dto.Payment) (dto.PaymentResponse, error) {
	args := mock.Called(ctx, payment)
	err := args.Error(1)

	if err != nil {
		return dto.PaymentResponse{}, err
	}

	return args.Get(0).(dto.PaymentResponse), nil
}

func (mock *MockGetOrdersToPrepareUseCase) Execute(ctx context.Context) ([]dto.OrderResponse, error) {
	args := mock.Called(ctx)
	err := args.Error(1)

	if err != nil {
		return []dto.OrderResponse{}, err
	}

	return args.Get(0).([]dto.OrderResponse), nil
}

func (mock *MockGetOrdersToFollowUseCase) Execute(ctx context.Context) ([]dto.OrderResponse, error) {
	args := mock.Called(ctx)
	err := args.Error(1)

	if err != nil {
		return []dto.OrderResponse{}, err
	}

	return args.Get(0).([]dto.OrderResponse), nil
}

func (mock *MockGetOrdersWaitingPaymentUseCase) Execute(ctx context.Context) ([]dto.OrderResponse, error) {
	args := mock.Called(ctx)
	err := args.Error(1)

	if err != nil {
		return []dto.OrderResponse{}, err
	}

	return args.Get(0).([]dto.OrderResponse), nil
}

func (m *MockCreateOrderUseCase) Execute(
	ctx context.Context,
	order dto.Order,
	date int64,
	wg *sync.WaitGroup,
	_ chan bool) (dto.OrderResponse, error) {
	args := m.Called(ctx, order, date, wg, mock.Anything)
	err := args.Error(1)

	if err != nil {
		return dto.OrderResponse{}, err
	}

	return args.Get(0).(dto.OrderResponse), nil
}

func (m *MockCreateProductUseCase) Execute(ctx context.Context, product dto.ProductForm) (uint, error) {
	args := m.Called(ctx, product)
	err := args.Error(1)

	if err != nil {
		return uint(0), err
	}

	return args.Get(0).(uint), nil
}

func (m *MockGetProductsByCategoryUseCase) Execute(ctx context.Context, category string) ([]dto.ProductResponse, error) {
	args := m.Called(ctx, category)
	err := args.Error(1)

	if err != nil {
		return []dto.ProductResponse{}, err
	}

	return args.Get(0).([]dto.ProductResponse), nil
}

func (m *MockGetProductsByIDUseCase) Execute(ctx context.Context, id uint) (dto.ProductResponse, error) {
	args := m.Called(ctx, id)
	err := args.Error(1)

	if err != nil {
		return dto.ProductResponse{}, err
	}

	return args.Get(0).(dto.ProductResponse), nil
}

func (m *MockGenerateQRCodePaymentUseCase) Execute(
	ctx context.Context,
	token string,
	qrOrder dto.QRCodeOrder,
	date int64,
	wg *sync.WaitGroup,
	ch chan bool,
) (dto.QRCodeDataResponse, error) {
	args := m.Called(ctx, token, qrOrder, date, wg, mock.Anything)
	err := args.Error(1)

	if err != nil {
		return dto.QRCodeDataResponse{}, err
	}

	return args.Get(0).(dto.QRCodeDataResponse), nil
}

func (mock *MockCreateOrderUseCase) GenerateTicket(ctx context.Context, date int64) int {
	args := mock.Called(ctx, date)
	err := args.Error(1)

	if err != nil {
		return 0
	}

	return args.Get(0).(int)
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

func (mock *MockUpdateToPreparingUseCase) Execute(ctx context.Context, orderId uint) error {
	args := mock.Called(ctx, orderId)
	err := args.Error(1)

	if err != nil {
		return err
	}

	return nil
}

func (mock *MockDeleteProductUseCase) Execute(ctx context.Context, productId uint) error {
	args := mock.Called(ctx, productId)
	err := args.Error(0)

	if err != nil {
		return err
	}

	return nil
}

func (mock *MockUpdateProductUseCase) Execute(ctx context.Context, product dto.ProductForm) error {
	args := mock.Called(ctx, product)
	err := args.Error(0)

	if err != nil {
		return err
	}

	return nil
}

func (mock *MockUpdateToDoneUseCase) Execute(ctx context.Context, orderId uint) error {
	args := mock.Called(ctx, orderId)
	err := args.Error(1)

	if err != nil {
		return err
	}

	return nil
}

func (mock *MockUpdateToDeliveredUseCase) Execute(ctx context.Context, orderId uint) error {
	args := mock.Called(ctx, orderId)
	err := args.Error(1)

	if err != nil {
		return err
	}

	return nil
}

func (mock *MockUpdateToNotDeliveredUseCase) Execute(ctx context.Context, orderId uint) error {
	args := mock.Called(ctx, orderId)
	err := args.Error(1)

	if err != nil {
		return err
	}

	return nil
}

func (mock *MockGetOrderByIdUseCase) Execute(ctx context.Context, orderId uint) (dto.OrderResponse, error) {
	args := mock.Called(ctx, orderId)
	err := args.Error(1)

	if err != nil {
		return dto.OrderResponse{}, err
	}

	return args.Get(0).(dto.OrderResponse), nil
}

func (mock *MockGetPaymentTypesUseCase) Execute() []string {
	args := mock.Called()
	return args.Get(0).([]string)
}

func (mock *MockGetCategoryUseCase) Execute() []string {
	args := mock.Called()
	return args.Get(0).([]string)
}
