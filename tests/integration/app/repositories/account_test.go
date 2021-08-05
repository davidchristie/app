package repositories

import (
	"context"
	"database/sql"
	"testing"

	"github.com/davidchristie/app/services/app/repositories"
	"github.com/stretchr/testify/suite"
)

type AccountRepositoryTestSuite struct {
	suite.Suite
	DB                *sql.DB
	AccountRepository repositories.AccountRepository
	UserRepository    repositories.UserRepository
}

func (suite *AccountRepositoryTestSuite) SetupSuite() {
	db := connectToDatabase(suite.T())
	suite.DB = db
	suite.AccountRepository = repositories.NewAccountRepository(db)
	suite.UserRepository = repositories.NewUserRepository(db)
}

func (suite *AccountRepositoryTestSuite) TeardownSuite() {
	suite.DB.Close()
}

func (suite *AccountRepositoryTestSuite) TestFindByID() {
	user := randomUser()
	err := suite.UserRepository.Insert(context.Background(), user)
	suite.Require().NoError(err)
	account := randomAccount(user.ID)
	err = suite.AccountRepository.Insert(context.Background(), account)
	suite.Require().NoError(err)
	record, err := suite.AccountRepository.FindByID(context.Background(), account.ID)
	suite.Require().NoError(err)
	assertAccountEqual(suite.T(), account, record)
}

func (suite *AccountRepositoryTestSuite) TestFindByProvider() {
	user := randomUser()
	err := suite.UserRepository.Insert(context.Background(), user)
	suite.Require().NoError(err)
	account := randomAccount(user.ID)
	err = suite.AccountRepository.Insert(context.Background(), account)
	suite.Require().NoError(err)
	record, err := suite.AccountRepository.FindByProvider(
		context.Background(),
		account.ProviderType,
		account.ProviderID,
		account.ProviderAccountID,
	)
	suite.Require().NoError(err)
	assertAccountEqual(suite.T(), account, record)
}

func TestAccountRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(AccountRepositoryTestSuite))
}
