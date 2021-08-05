package repositories

import (
	"context"
	"database/sql"
	"testing"

	"github.com/davidchristie/app/services/app/repositories"
	"github.com/stretchr/testify/suite"
)

type SessionRepositoryTestSuite struct {
	suite.Suite
	DB                *sql.DB
	SessionRepository repositories.SessionRepository
	UserRepository    repositories.UserRepository
}

func (suite *SessionRepositoryTestSuite) SetupSuite() {
	db := connectToDatabase(suite.T())
	suite.DB = db
	suite.SessionRepository = repositories.NewSessionRepository(db)
	suite.UserRepository = repositories.NewUserRepository(db)
}

func (suite *SessionRepositoryTestSuite) TeardownSuite() {
	suite.DB.Close()
}

func (suite *SessionRepositoryTestSuite) TestFindByID() {
	user := randomUser()
	err := suite.UserRepository.Insert(context.Background(), user)
	suite.Require().NoError(err)
	session := randomSession(suite.T(), user.ID)
	err = suite.SessionRepository.Insert(context.Background(), session)
	suite.Require().NoError(err)
	record, err := suite.SessionRepository.FindByID(context.Background(), session.ID)
	suite.Require().NoError(err)
	assertSessionEqual(suite.T(), session, record)
}

func (suite *SessionRepositoryTestSuite) TestFindBySessionToken() {
	user := randomUser()
	err := suite.UserRepository.Insert(context.Background(), user)
	suite.Require().NoError(err)
	session := randomSession(suite.T(), user.ID)
	err = suite.SessionRepository.Insert(context.Background(), session)
	suite.Require().NoError(err)
	record, err := suite.SessionRepository.FindBySessionToken(context.Background(), session.SessionToken)
	suite.Require().NoError(err)
	assertSessionEqual(suite.T(), session, record)
}

func TestSessionRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(SessionRepositoryTestSuite))
}
