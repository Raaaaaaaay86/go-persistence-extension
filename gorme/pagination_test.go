package gorme_test

import (
	"context"
	"fmt"
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

		pagination, err := s.UserRepository.PFindTimeBefore(context.Background(), target, beforeAt, currentPage, 1)
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

		pagination, err := s.UserRepository.PFindTimeAfter(context.Background(), target, beforeAt, currentPage, 1)
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

func (s *PaginationOperationTestSuite) TestPFindTimeBetween() {
	s.T().Log("Test_TestPFindTimeBetween: start")
	currentPage := 1
	for {
		target := entity.User{
			Birthday: mark.TargetTime,
		}
		startAt := time.Date(2000, 6, 1, 0, 0, 0, 0, time.UTC)
		endAt := time.Date(2000, 10, 1, 0, 0, 0, 0, time.UTC)

		pagination, err := s.UserRepository.PFindTimeBetween(context.Background(), target, startAt, endAt, currentPage, 1)
		if err != nil {
			s.T().Fatalf("Test_TestPFindTimeBetween: failed (%s)", err.Error())
		}

		for _, user := range pagination.Results {
			s.T().Logf("birthday %v should between %v and %v", user.Birthday, startAt, endAt)
			assert.True(s.T(), user.Birthday.After(startAt))
			assert.True(s.T(), user.Birthday.Before(endAt))
		}

		if !pagination.HasNext() {
			break
		}

		currentPage++
	}
}

func (s *PaginationOperationTestSuite) TestPFindIntGT() {
	s.T().Log("Test_TestPFindIntGT: start")

	currentPage := 1
	for {
		target := entity.User{
			Age: mark.TargetInt,
		}

		greaterThan := 22
		pagination, err := s.UserRepository.PFindIntGT(context.Background(), target, greaterThan, currentPage, 1)
		if err != nil {
			s.T().Fatalf("Test_TestPFindIntGT: failed (%s)", err.Error())
		}

		for _, user := range pagination.Results {
			fmt.Println(user.Age)
			assert.Greater(s.T(), user.Age, greaterThan)
		}

		if !pagination.HasNext() {
			break
		}

		currentPage++
	}
}

func (s *PaginationOperationTestSuite) TestPFindIntGTE() {
	s.T().Log("Test_TestPFindIntGTE: start")

	currentPage := 1
	for {
		target := entity.User{
			Age: mark.TargetInt,
		}

		greaterThan := 23
		pagination, err := s.UserRepository.PFindIntGTE(context.Background(), target, greaterThan, currentPage, 1)
		if err != nil {
			s.T().Fatalf("Test_TestPFindIntGTE: failed (%s)", err.Error())
		}

		for _, user := range pagination.Results {
			fmt.Println(user.Age)
			assert.GreaterOrEqual(s.T(), user.Age, greaterThan)
		}

		if !pagination.HasNext() {
			break
		}

		currentPage++
	}
}

func (s *PaginationOperationTestSuite) TestPFindIntLT() {
	s.T().Log("Test_TestPFindIntLT: start")

	currentPage := 1
	for {
		target := entity.User{
			Age: mark.TargetInt,
		}

		greaterThan := 23
		pagination, err := s.UserRepository.PFindIntLT(context.Background(), target, greaterThan, currentPage, 1)
		if err != nil {
			s.T().Fatalf("Test_TestPFindIntLT: failed (%s)", err.Error())
		}

		for _, user := range pagination.Results {
			fmt.Println(user.Age)
			assert.Less(s.T(), user.Age, greaterThan)
		}

		if !pagination.HasNext() {
			break
		}

		currentPage++
	}
}

func (s *PaginationOperationTestSuite) TestPFindIntLTE() {
	s.T().Log("Test_TestPFindIntLTE: start")

	currentPage := 1
	for {
		target := entity.User{
			Age: mark.TargetInt,
		}

		greaterThan := 23
		pagination, err := s.UserRepository.PFindIntLTE(context.Background(), target, greaterThan, currentPage, 1)
		if err != nil {
			s.T().Fatalf("Test_TestPFindIntLTE: failed (%s)", err.Error())
		}

		for _, user := range pagination.Results {
			fmt.Println(user.Age)
			assert.LessOrEqual(s.T(), user.Age, greaterThan)
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
