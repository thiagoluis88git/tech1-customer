package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/thiagoluis88git/tech1-customer/internal/core/domain/dto"
	"github.com/thiagoluis88git/tech1-customer/internal/core/domain/usecases"
	"github.com/thiagoluis88git/tech1-customer/pkg/httpserver"
)

// @Summary Create new customer
// @Description Create new customer. This process is not required to make an order
// @Tags Customer
// @Accept json
// @Produce json
// @Param product body dto.Customer true "customer"
// @Success 200 {object} dto.CustomerResponse
// @Failure 400 "Customer has required fields"
// @Failure 409 "This Customer is already added"
// @Router /auth/signup [post]
func CreateCustomerHandler(createCustomer usecases.CreateCustomerUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var customer dto.Customer

		err := httpserver.DecodeJSONBody(w, r, &customer)

		if err != nil {
			log.Print("decoding customer body", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		response, err := createCustomer.Execute(r.Context(), customer)

		if err != nil {
			log.Print("create customer", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		httpserver.SendResponseSuccess(w, response)
	}
}

// @Summary Update customer
// @Description Update customer
// @Tags Customer
// @Accept json
// @Produce json
// @Param id path int true "12"
// @Param product body dto.Customer true "customer"
// @Success 204
// @Failure 400 "Customer has required fields"
// @Failure 404 "Customer not found"
// @Router /api/admin/customers/{id} [put]
func UpdateCustomerHandler(updateCustomer usecases.UpdateCustomerUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		customerIdStr, err := httpserver.GetPathParamFromRequest(r, "id")

		if err != nil {
			log.Print("update customer", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendBadRequestError(w, err)
			return
		}

		customerId, err := strconv.Atoi(customerIdStr)

		if err != nil {
			log.Print("update customer", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendBadRequestError(w, err)
			return
		}

		var customer dto.Customer
		err = httpserver.DecodeJSONBody(w, r, &customer)

		if err != nil {
			log.Print("decoding customer body for update", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		customer.ID = uint(customerId)
		err = updateCustomer.Execute(r.Context(), customer)

		if err != nil {
			log.Print("update customer", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		httpserver.SendResponseNoContentSuccess(w)
	}
}

// @Summary Get customer by ID
// @Description Get customer by ID
// @Tags Customer
// @Accept json
// @Produce json
// @Param Id path string true "12"
// @Success 200 {object} dto.Customer
// @Failure 404 "Customer not found"
// @Router /api/customers/{id} [get]
func GetCustomerByIdHandler(getCustomerById usecases.GetCustomerByIdUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		customerIdStr, err := httpserver.GetPathParamFromRequest(r, "id")

		if err != nil {
			log.Print("get customer by id", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendBadRequestError(w, err)
			return
		}

		customerId, err := strconv.Atoi(customerIdStr)

		if err != nil {
			log.Print("get customer by id", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendBadRequestError(w, err)
			return
		}

		customer, err := getCustomerById.Execute(r.Context(), uint(customerId))

		if err != nil {
			log.Print("get customer by id", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		httpserver.SendResponseSuccess(w, customer)
	}
}

// @Summary Get customer by CPF
// @Description Get customer by CPF. This Endpoint can be used as a Login
// @Tags Customer
// @Accept json
// @Produce json
// @Param customer body dto.CustomerForm true "customerForm"
// @Success 200 {object} dto.Customer
// @Failure 404 "Customer not found"
// @Router /api/customers/{cpf} [post]
func GetCustomerByCPFHandler(getCustomerByCPF usecases.GetCustomerByCPFUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cpf, err := httpserver.GetPathParamFromRequest(r, "cpf")

		if err != nil {
			log.Print("update customer", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendBadRequestError(w, err)
			return
		}

		customer, err := getCustomerByCPF.Execute(r.Context(), cpf)

		if err != nil {
			log.Print("get customer by id", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		httpserver.SendResponseSuccess(w, customer)
	}
}
