package repositories

import (
	"context"
	"database/sql"
	"testing"

	"github.com/davidchristie/app/services/app/repositories"
	"github.com/stretchr/testify/suite"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	DB             *sql.DB
	UserRepository repositories.UserRepository
}

func (suite *UserRepositoryTestSuite) SetupSuite() {
	db := connectToDatabase(suite.T())
	suite.DB = db
	suite.UserRepository = repositories.NewUserRepository(db)
}

func (suite *UserRepositoryTestSuite) TeardownSuite() {
	suite.DB.Close()
}

func (suite *UserRepositoryTestSuite) TestFindByID() {
	user := randomUser()
	err := suite.UserRepository.Insert(context.Background(), user)
	suite.Require().NoError(err)
	record, err := suite.UserRepository.FindByID(context.Background(), user.ID)
	suite.Require().NoError(err)
	assertUserEqual(suite.T(), user, record)
}

func (suite *UserRepositoryTestSuite) TestFindByPrimaryEmail() {
	user := randomUser()
	err := suite.UserRepository.Insert(context.Background(), user)
	suite.Require().NoError(err)
	record, err := suite.UserRepository.FindByPrimaryEmail(context.Background(), user.PrimaryEmail)
	suite.Require().NoError(err)
	assertUserEqual(suite.T(), user, record)
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
