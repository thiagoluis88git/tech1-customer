package main

import (
	"fmt"
	"net/http"

	"github.com/thiagoluis88git/tech1-customer/internal/core/data/repositories"
	"github.com/thiagoluis88git/tech1-customer/internal/core/domain/usecases"
	"github.com/thiagoluis88git/tech1-customer/internal/core/handler"
	"github.com/thiagoluis88git/tech1-customer/internal/integrations/remote"
	"github.com/thiagoluis88git/tech1-customer/pkg/database"
	"github.com/thiagoluis88git/tech1-customer/pkg/environment"
	"github.com/thiagoluis88git/tech1-customer/pkg/httpserver"
	"github.com/thiagoluis88git/tech1-customer/pkg/responses"
	"gorm.io/driver/postgres"

	"github.com/mvrilo/go-redoc"

	"github.com/go-chi/chi/v5"

	_ "github.com/thiagoluis88git/tech1-customer/docs"

	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title Tech1 Customer Docs
// @version 1.0
// @description This is the API for the Tech1 Customer Project.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localshot:3210
// @BasePath /
func main() {
	environment.LoadEnvironmentVariables()

	doc := redoc.Redoc{
		Title:       "Example API",
		Description: "Example API Description",
		SpecFile:    *environment.RedocFolderPath,
		SpecPath:    "/docs/swagger.json",
		DocsPath:    "/docs",
	}

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v",
		environment.GetDBHost(),
		environment.GetDBUser(),
		environment.GetDBPassword(),
		environment.GetDBName(),
		environment.GetDBPort(),
	)

	db, err := database.ConfigDatabase(postgres.Open(dsn))

	if err != nil {
		panic(fmt.Sprintf("could not open database: %v", err.Error()))
	}

	router := chi.NewRouter()
	router.Use(chiMiddleware.RequestID)
	router.Use(chiMiddleware.RealIP)
	router.Use(chiMiddleware.Recoverer)

	// httpClient := httpserver.NewHTTPClient()

	cognitoRemote := remote.NewCognitoRemoteDataSource(
		environment.GetRegion(),
		environment.GetCognitoUserPoolID(),
		environment.GetCognitoClientID(),
		environment.GetCognitoGroupUser(),
		environment.GetCognitoGroupAdmin(),
	)
	customerRepo := repositories.NewCustomerRepository(db, cognitoRemote)
	userRepo := repositories.NewUserAdminRepository(db, cognitoRemote)
	validateCPFUseCase := usecases.NewValidateCPFUseCase()
	loginCustomerUseCase := usecases.NewLoginCustomerUseCase(customerRepo)
	loginUnknownCustomerUseCase := usecases.NewLoginUnknownCustomerUseCase(customerRepo)
	createCustomerUseCase := usecases.NewCreateCustomerUseCase(validateCPFUseCase, customerRepo)
	updateCustomerUseCase := usecases.NewUpdateCustomerUseCase(validateCPFUseCase, customerRepo)
	getCustomerByCPFUseCase := usecases.NewGetCustomerByCPFUseCase(validateCPFUseCase, customerRepo)

	loginUserUseCase := usecases.NewLoginUserUseCase(userRepo)
	createUserUseCase := usecases.NewCreateUserUseCase(validateCPFUseCase, userRepo)
	updateUserUseCase := usecases.NewUpdateUserUseCase(validateCPFUseCase, userRepo)
	getUserByIdUseCase := usecases.NewGetUserByIdUseCase(userRepo)
	getUserByCPFUseCase := usecases.NewGetUserByCPFUseCase(validateCPFUseCase, userRepo)

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		httpserver.SendResponseSuccess(w, &responses.BusinessResponse{
			StatusCode: 200,
			Message:    "ok",
		})
	})

	router.Post("/auth/login", handler.LoginCustomerHandler(loginCustomerUseCase))
	router.Post("/auth/login/unknown", handler.LoginUnknownCustomerHandler(loginUnknownCustomerUseCase))
	router.Post("/auth/admin/login", handler.LoginUserHandler(loginUserUseCase))
	router.Post("/auth/signup", handler.CreateCustomerHandler(createCustomerUseCase))
	router.Post("/auth/admin/signup", handler.CreateUserHandler(createUserUseCase))

	router.Put("/api/admin/customers/{id}", handler.UpdateCustomerHandler(updateCustomerUseCase))
	router.Get("/api/customers/{cpf}", handler.GetCustomerByCPFHandler(getCustomerByCPFUseCase))

	router.Put("/api/users/{id}", handler.UpdateUserHandler(updateUserUseCase))
	router.Get("/api/users/{id}", handler.GetUserByIdHandler(getUserByIdUseCase))
	router.Post("/api/users/login", handler.GetUserByCPFHandler(getUserByCPFUseCase))

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3210/swagger/doc.json"),
	))

	go http.ListenAndServe(":3211", doc.Handler())

	server := httpserver.New(router)
	server.Start()
}
