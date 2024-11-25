package bdd_test

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/thiagoluis88git/tech1-customer/internal/core/domain/dto"
)

type MockLoginCustomerUseCase struct {
	mock.Mock
}

func (mock *MockLoginCustomerUseCase) Execute(ctx context.Context, cpf string) (dto.Token, error) {
	args := mock.Called(ctx, cpf)
	err := args.Error(1)

	if err != nil {
		return dto.Token{}, err
	}

	return args.Get(0).(dto.Token), nil
}
