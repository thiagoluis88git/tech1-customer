package remote_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thiagoluis88git/tech1-customer/internal/core/data/model"
	"github.com/thiagoluis88git/tech1-customer/internal/integrations/remote"
)

func TestCognitoRemote(t *testing.T) {
	t.Parallel()

	t.Run("got error when login cognito remote", func(t *testing.T) {
		sut := remote.NewCognitoRemoteDataSource("region", "userPool", "appClient", "groupUser", "adminUser")

		result, err := sut.Login("cpf")
		assert.Error(t, err)
		assert.Empty(t, result)
	})

	t.Run("got error when login unknown cognito remote", func(t *testing.T) {
		sut := remote.NewCognitoRemoteDataSource("region", "userPool", "appClient", "groupUser", "adminUser")

		result, err := sut.LoginUnknown()
		assert.Error(t, err)
		assert.Empty(t, result)
	})

	t.Run("got error when sign up cognito remote", func(t *testing.T) {
		sut := remote.NewCognitoRemoteDataSource("region", "userPool", "appClient", "groupUser", "adminUser")

		err := sut.SignUp(&model.Customer{})
		assert.Error(t, err)
	})

	t.Run("got error when sign up admin cognito remote", func(t *testing.T) {
		sut := remote.NewCognitoRemoteDataSource("region", "userPool", "appClient", "groupUser", "adminUser")

		err := sut.SignUpAdmin(&model.UserAdmin{})
		assert.Error(t, err)
	})
}
