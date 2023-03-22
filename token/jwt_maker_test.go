package token

import (
	"fmt"
	"simplebank/util"
	"testing"
	"time"

	"github.com/bxcodec/faker/v4"
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	fmt.Printf(util.RandomString(32))
	stringVal := util.RandomString(32)
	fmt.Println("Length :: ", len(stringVal))
	maker, err := NewJWTMaker(util.RandomString(32))

	fmt.Println(maker)

	require.NoError(t, err)

	userName := faker.FirstName() + "_" + faker.LastName()
	duration := time.Minute

	issueAt := time.Now()
	expireAt := issueAt.Add(duration)

	token, payload, err := maker.CreateToken(userName, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err = maker.VerifyToken(token)
	require.NoError(t, err)
	require.Equal(t, userName, payload.Username)

	require.WithinDuration(t, issueAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expireAt, payload.ExpiredAt, time.Second)

}
