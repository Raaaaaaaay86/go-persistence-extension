package gorme_test

import (
	"context"
	"fmt"
	"slices"
	"testing"
	"time"

	"github.com/raaaaaaaay86/go-persistence-extension/gorme/entity"
	"github.com/raaaaaaaay86/go-persistence-extension/gorme/repository"
	"github.com/raaaaaaaay86/go-persistence-extension/gorme/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type BasicOperationTestSuite struct {
	suite.Suite
	UserRepository *repository.UserRepository
}

func (s *BasicOperationTestSuite) SetupTest() {
	if err := s.Setup(); err != nil {
		s.T().Fatalf("failed to setup BasicOperationTestSuite: %s", err.Error())
	}
}

func (s *BasicOperationTestSuite) Setup() error {
	db, err := util.CreateGormPostgreSqlConnection(&gorm.Config{TranslateError: true})
	if err != nil {
		return err
	}

	s.UserRepository = repository.NewUserRepository(db)

	return nil
}

func (s *BasicOperationTestSuite) Test_GetById() {
	ctx := context.Background()

	user, err := s.UserRepository.GetById(ctx, 1)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), uint(1), user.ID)
	assert.Equal(s.T(), "user1", user.Username)

	_, err = s.UserRepository.GetById(ctx, 999) 
	assert.ErrorIs(s.T(), gorm.ErrRecordNotFound, err)
}

func (s *BasicOperationTestSuite) Test_GetBy() {
	ctx := context.Background()

	type GetByTestCase struct {
		TestDescription string
		Condition entity.User
		ValidateFn func(*testing.T, *entity.User, error)
	}

	table := []GetByTestCase{
		{
			TestDescription: "Get user by birthday",
			Condition: entity.User{Birthday: time.Date(2000, 3, 3, 0, 0, 0, 0, time.UTC)},
			ValidateFn: func(t *testing.T, user *entity.User, err error) {
				assert.NoError(t, err)
				assert.Equal(t, uint(1), user.ID)
				assert.Equal(t, "user1", user.Username)
			},
		},
		{
			TestDescription: "Get user by username",
			Condition: entity.User{Username: "user1"},
			ValidateFn: func(t *testing.T, user *entity.User, err error) {
				assert.NoError(t, err)
				assert.Equal(t, uint(1), user.ID)
				assert.Equal(t, "user1", user.Username)
			},
		},
		{
			TestDescription: "Get user by age",
			Condition: entity.User{Age: 20},
			ValidateFn: func(t *testing.T, user *entity.User, err error) {
				assert.NoError(t, err)
				assert.Equal(t, uint(1), user.ID)
				assert.Equal(t, "user1", user.Username)
			},
		},
		{
			TestDescription: "Get user by age and username",
			Condition: entity.User{Age: 20, Username: "user1"},
			ValidateFn: func(t *testing.T, user *entity.User, err error) {
				assert.NoError(t, err)
				assert.Equal(t, uint(1), user.ID)
				assert.Equal(t, "user1", user.Username)
			},
		},
		{
			TestDescription: "Get non-exist user by username",
			Condition: entity.User{Username: "non-exist"},
			ValidateFn: func(t *testing.T, user *entity.User, err error) {
				assert.ErrorIs(t, gorm.ErrRecordNotFound, err)
			},
		},
	}

	for _, tc := range table {
		s.T().Logf("Test_GetBy: %s", tc.TestDescription)
		user, err := s.UserRepository.GetBy(ctx, tc.Condition)
		tc.ValidateFn(s.T(), user, err)
	}
}

func (s *BasicOperationTestSuite) Test_FindBy() {
	ctx := context.Background()

	s.T().Log("Test_FindBy: Find users by age 20 with no limit")
	users, err := s.UserRepository.FindBy(ctx, entity.User{Age: 20}, -1)
	assert.NoError(s.T(), err)
	expectedUserNames := []string{"user1", "user3", "user6", "user8", "user10"}
	for _, user := range users {
		assert.Equal(s.T(), 20, user.Age)
		assert.True(s.T(), slices.Contains(expectedUserNames, user.Username))
	}

	s.T().Log("Test_FindBy: Find users by age 20 with limit 1")
	limit := 1
	users, err = s.UserRepository.FindBy(ctx, entity.User{Age: 20}, limit)
	assert.NoError(s.T(), err)
	assert.Len(s.T(), users, limit)

	s.T().Log("Test_FindBy: Find users by age 100 with no limit")
	users, err = s.UserRepository.FindBy(ctx, entity.User{Age: 100}, -1)
	assert.NoError(s.T(), err)
	assert.Len(s.T(), users, 0)
}

func (s *BasicOperationTestSuite) Test_FindByAll() {
	ctx := context.Background()

	s.T().Log("Test_FindAll: Find all users with no limit")
	_, err := s.UserRepository.FindAll(ctx, -1)
	assert.NoError(s.T(), err)

	s.T().Log("Test_FindAll: Find all users with limit 1")
	users, err := s.UserRepository.FindAll(ctx, 1)
	assert.NoError(s.T(), err)
	assert.Len(s.T(), users, 1)
}

func (s *BasicOperationTestSuite) Test_CreateAndDeleteById() {
	ctx := context.Background()

	s.T().Log("Test_CreateAndDelete: Create user")
	user := entity.User{
		Username: fmt.Sprintf( "delete_by_id_%d", time.Now().Unix()),
		Email: "delete_by_id@mail.com",
		Birthday: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		Age: 10,
	}
	err := s.UserRepository.Create(ctx, &user)
	assert.NoError(s.T(), err)
	assert.NotEmpty(s.T(), user.ID)

	s.T().Log("Test_CreateAndDelete: Delete user by ID")
	affectedCount, err := s.UserRepository.DeleteById(ctx, user.ID)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), int64(1), affectedCount)
}

func (s *BasicOperationTestSuite) Test_CreateAndDeleteByStruct() {
	ctx := context.Background()

	s.T().Log("Test_CreateAndDeleteByStruct: Create user")
	user := entity.User{
		Username: fmt.Sprintf( "delete_by_struct_%d", time.Now().Unix()),
		Email: "delete_by_struct@mail.com",
		Birthday: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		Age: 10,
	}
	err := s.UserRepository.Create(ctx, &user)
	assert.NoError(s.T(), err)
	assert.NotEmpty(s.T(), user.ID)

	s.T().Log("Test_CreateAndDeleteByStruct: Delete user by struct")
	affectedCount, err := s.UserRepository.Delete(ctx, &user)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), int64(1), affectedCount)
}

func (s *BasicOperationTestSuite) Test_CreateAndUpdateByStruct() {
	ctx := context.Background()

	s.T().Log("Test_CreateAndUpdateByStruct: Create user")
	user := entity.User{
		Username: fmt.Sprintf( "update_by_struct_%d", time.Now().Unix()),
		Email: "update_by_struct@mail.com",
		Birthday: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		Age: 10,
	}
	err := s.UserRepository.Create(ctx, &user)
	assert.NoError(s.T(), err)
	assert.NotEmpty(s.T(), user.ID)

	s.T().Log("Test_CreateAndUpdateByStruct: Update user by struct")
	user.Age = 11
	affectedCount, err := s.UserRepository.Update(ctx, &user)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), int64(1), affectedCount)

	queryUser, err := s.UserRepository.GetById(ctx, user.ID)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 11, queryUser.Age)
}

func TestRunBasicOperationTestSuite(t *testing.T) {
	suite.Run(t, new(BasicOperationTestSuite))
}
