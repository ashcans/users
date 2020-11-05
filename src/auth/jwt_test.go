package auth

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateToken(t *testing.T) {
	uid, secret := "1", "123"
	token, err := CreateToken(uid, secret)
	if err != nil {
		panic(err.Error())
	}
	println(token)
}

func TestParseToken(t *testing.T) {
	a := assert.New(t)
	uid, secret, token := "1", "123", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDQyODI5ODksInVpZCI6IjEifQ.JRKqkD6sFJ9xSShEd48DSu8l5UhhXLg8IydvqNfeNyY"
	result, err := ParseToken(token, secret)
	if err != nil {
		panic(err.Error())
	}
	a.Equal(uid, result)
}
func TestCreateToken2(t *testing.T) {
	println(uuid.New())

}
