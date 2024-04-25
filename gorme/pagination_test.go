package gorme_test

import (
	"context"
	"testing"
	"time"

	"github.com/raaaaaaaay86/go-persistence-extension/gorme/entity"
	"github.com/raaaaaaaay86/go-persistence-extension/gorme/repository"
	"github.com/raaaaaaaay86/go-persistence-extension/gorme/util"
	"github.com/raaaaaaaay86/go-persistence-extension/mark"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type PaginationOperationTestSuite struct {
	suite.Suite
	UserRepository *repository.UserRepository
}

func (s *PaginationOperationTestSuite) SetupTest() {
	if err := s.Setup(); err != nil {
		s.T().Fatalf("failed to setup PaginationOperationTestSuite: %s", err.Error())
	}
}

func (s *PaginationOperationTestSuite) Setup() error {
	db, err := util.CreateGormPostgreSqlConnection(&gorm.Config{TranslateError: true})
	if err != nil {
		return err
	}

	s.UserRepository = repository.NewUserRepository(db.Debug())

	return nil
}

func (s *PaginationOperationTestSuite) TestPFindTimeBefore() {
	s.T().Log("Test_PFindTimeBefore: start")
	currentPage := 1
	for {
		target := entity.User{
			Birthday: mark.TargetTime,
		}
		beforeAt := time.Date(2000, 8, 1, 0, 0, 0, 0, time.UTC)

		pagination, err := s.UserRepository.PFindTimeBefore(context.Background(), target, beforeAt, currentPage, 10)
		if err != nil {
			s.T().Fatalf("Test_PFindTimeBefore: failed (%s)", err.Error())
		}

		for _, user := range pagination.Results {
			s.T().Logf("birthday %v should before %v", user.Birthday, beforeAt)
			assert.True(s.T(), user.Birthday.Before(beforeAt))
		}

		if !pagination.HasNext() {
			break
		}

		currentPage++
	}
}

func (s *PaginationOperationTestSuite) TestPFindTimeAfter() {
	s.T().Log("Test_PFindTimeAfter: start")
	currentPage := 1
	for {
		target := entity.User{
			Birthday: mark.TargetTime,
		}
		beforeAt := time.Date(2000, 8, 1, 0, 0, 0, 0, time.UTC)

		pagination, err := s.UserRepository.PFindTimeAfter(context.Background(), target, beforeAt, currentPage, 10)
		if err != nil {
			s.T().Fatalf("Test_PFindTimeAfter: failed (%s)", err.Error())
		}

		for _, user := range pagination.Results {
			s.T().Logf("birthday %v should after %v", user.Birthday, beforeAt)
			assert.True(s.T(), user.Birthday.After(beforeAt))
		}

		if !pagination.HasNext() {
			break
		}

		currentPage++
	}
}

func TestPaginationOperationTestSuite(t *testing.T) {
	suite.Run(t, new(PaginationOperationTestSuite))
}