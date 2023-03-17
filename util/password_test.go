package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPassword(t *testing.T) {
	password := "password"
	hashedPassword, err := HashedPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	require.True(t, checkPasswordHash(password, hashedPassword))
	require.False(t, checkPasswordHash("wrong password", hashedPassword))
}
