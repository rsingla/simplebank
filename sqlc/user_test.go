package sqlc

import (
	"context"
	"simplebank/util"
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/stretchr/testify/require"
)

func TestCreateRandomUser(t *testing.T) {

	password, err := util.HashedPassword(util.RandomPassword())

	if err != nil {
		t.Fatal(err)
	}

	firstName := faker.FirstName()
	lastName := faker.LastName()
	userName := firstName + "_" + lastName

	arg := CreateUserParams{
		Username:       userName,
		HashedPassword: password,
		FirstName:      firstName,
		LastName:       lastName,
		Email:          util.RandomEmail(userName),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	if err != nil {
		t.Fatal(err)
	}

	require.NotEmpty(t, user)
	require.NotEmpty(t, user.CreatedAt)
	require.NotEmpty(t, user.Email)
	require.NotEmpty(t, user.FirstName)
}
