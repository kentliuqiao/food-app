package application

import (
	"food-app/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	saveUserRepo                func(*entity.User) (*entity.User, error)
	getUserRepo                 func(userId uint64) (*entity.User, error)
	getUsersRepo                func() ([]*entity.User, error)
	getUserEmailAndPasswordRepo func(*entity.User) (*entity.User, error)
)

type fakeUserRepo struct{}

func (u fakeUserRepo) SaveUser(user *entity.User) (*entity.User, error) {
	return saveUserRepo(user)
}

func (u fakeUserRepo) GetUser(userId uint64) (*entity.User, error) {
	return getUserRepo(userId)
}

func (u fakeUserRepo) GetUsers() ([]*entity.User, error) {
	return getUsersRepo()
}

func (u fakeUserRepo) GetUserByEmailAndPassword(user *entity.User) (*entity.User, error) {
	return getUserEmailAndPasswordRepo(user)
}

var userAppFake UserAppInterface = fakeUserRepo{}

func TestSaveUser_Success(t *testing.T) {
	saveUserRepo = func(u *entity.User) (*entity.User, error) {
		return &entity.User{
			ID:        1,
			FirstName: "victor",
			LastName:  "steven",
			Email:     "steven@example.com",
			Password:  "password",
		}, nil
	}
	user := &entity.User{
		ID:        1,
		FirstName: "victor",
		LastName:  "steven",
		Email:     "steven@example.com",
		Password:  "password",
	}

	u, err := userAppFake.SaveUser(user)
	assert.Nil(t, err)
	assert.EqualValues(t, u.FirstName, "victor")
	assert.EqualValues(t, u.Email, "steven@example.com")
}
