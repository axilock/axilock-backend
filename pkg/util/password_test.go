package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandString(6)
	hashPassword1, err := NewPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashPassword1)
	err = CheckPassword(hashPassword1, password)
	require.NoError(t, err)
	wrongPassword := RandString(6)
	err = CheckPassword(hashPassword1, wrongPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashPassword2, err := NewPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashPassword2)

	require.NotEqual(t, hashPassword1, hashPassword2)
}
