package repositories

import (
	"context"

	"github.com/thiagoluis88git/tech1-customer/internal/core/data/model"
	"github.com/thiagoluis88git/tech1-customer/internal/core/domain/dto"
	"github.com/thiagoluis88git/tech1-customer/internal/core/domain/repository"
	"github.com/thiagoluis88git/tech1-customer/internal/integrations/remote"
	"github.com/thiagoluis88git/tech1-customer/pkg/database"
	"github.com/thiagoluis88git/tech1-customer/pkg/responses"

	"gorm.io/gorm"
)

type CustomerRepository struct {
	db            *database.Database
	cognitoRemote remote.CognitoRemoteDataSource
}

func NewCustomerRepository(db *database.Database, cognitoRemote remote.CognitoRemoteDataSource) repository.CustomerRepository {
	return &CustomerRepository{
		db:            db,
		cognitoRemote: cognitoRemote,
	}
}

func (repository *CustomerRepository) CreateCustomer(ctx context.Context, customer dto.Customer) (uint, error) {
	customerEntity := &model.Customer{
		Name:  customer.Name,
		CPF:   customer.CPF,
		Email: customer.Email,
	}

	err := repository.cognitoRemote.SignUp(customerEntity)

	if err != nil {
		return 0, responses.GetCognitoError(err)
	}

	err = repository.db.Connection.WithContext(ctx).Create(customerEntity).Error

	if err != nil {
		return 0, responses.GetDatabaseError(err)
	}

	return customerEntity.ID, nil
}

func (repository *CustomerRepository) UpdateCustomer(ctx context.Context, customer dto.Customer) error {
	customerEntity := &model.Customer{
		Model: gorm.Model{ID: customer.ID},
		Name:  customer.Name,
		CPF:   customer.CPF,
		Email: customer.Email,
	}

	err := repository.db.Connection.WithContext(ctx).Save(&customerEntity).Error

	if err != nil {
		return responses.GetDatabaseError(err)
	}

	return nil
}

func (repository *CustomerRepository) GetCustomerById(ctx context.Context, id uint) (dto.Customer, error) {
	var customerEntity model.Customer

	err := repository.
		db.Connection.WithContext(ctx).
		First(&customerEntity, id).
		Error

	if err != nil {
		return dto.Customer{}, responses.GetDatabaseError(err)
	}

	return repository.populateCustomer(customerEntity), nil
}

func (repository *CustomerRepository) GetCustomerByCPF(ctx context.Context, cpf string) (dto.Customer, error) {
	var customerEntity model.Customer

	err := repository.
		db.Connection.WithContext(ctx).
		Where("cpf = ?", cpf).
		First(&customerEntity).
		Error

	if err != nil {
		return dto.Customer{}, responses.GetDatabaseError(err)
	}

	return repository.populateCustomer(customerEntity), nil
}

func (repository *CustomerRepository) populateCustomer(customerEntity model.Customer) dto.Customer {
	return dto.Customer{
		ID:    customerEntity.ID,
		Name:  customerEntity.Name,
		CPF:   customerEntity.CPF,
		Email: customerEntity.Email,
	}
}

func (repository *CustomerRepository) Login(ctx context.Context, cpf string) (string, error) {
	token, err := repository.cognitoRemote.Login(cpf)

	if err != nil {
		return "", responses.GetDatabaseError(err)
	}

	return token, nil
}

func (repository *CustomerRepository) LoginUnknown() (string, error) {
	token, err := repository.cognitoRemote.LoginUnknown()

	if err != nil {
		return "", responses.GetDatabaseError(err)
	}

	return token, nil
}
