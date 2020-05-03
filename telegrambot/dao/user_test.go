package dao_test

import (
	"testing"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stretchr/testify/suite"

	"github.com/stretchr/testify/require"

	"github.com/ravil23/lingualynda/telegrambot/dao"
	"github.com/ravil23/lingualynda/telegrambot/postgres"
)

type UserDAOTestSuite struct {
	suite.Suite

	conn    *postgres.Connection
	userDAO dao.UserDAO
}

func TestUserDAOTestSuite(t *testing.T) {
	suite.Run(t, new(UserDAOTestSuite))
}

func (s *UserDAOTestSuite) SetupTest() {
	s.conn = postgres.NewConnection()

	var err error
	s.userDAO, err = dao.NewUserDAO(s.conn)
	s.NoError(err)
}

func (s *UserDAOTestSuite) TearDownTest() {
	err := s.conn.Close()
	s.NoError(err)
}

func (s *UserDAOTestSuite) deleteUser(userID dao.UserID) {
	err := s.userDAO.Delete(userID)
	if err != nil {
		s.EqualError(err, "pg: no rows in result set")
	}
}

func TestNewUser(t *testing.T) {
	tgUser := &tgbotapi.User{
		ID:        1,
		FirstName: "firstName",
		LastName:  "lastName",
		UserName:  "userName",
	}

	user := dao.NewUser(tgUser)

	require.Equal(t, user.ID, dao.UserID(tgUser.ID))
	require.Equal(t, user.FirstName, tgUser.FirstName)
	require.Equal(t, user.LastName, tgUser.LastName)
	require.Equal(t, user.NickName, tgUser.UserName)
	require.Empty(t, user.ChatID)
	require.Empty(t, user.CreatedAt)
	require.Empty(t, user.UpdatedAt)
}

func (s *UserDAOTestSuite) TestFind() {
	s.Run("NotFound", func() {
		userID := dao.UserID(1)
		s.deleteUser(userID)

		foundUser, err := s.userDAO.Find(userID)

		require.Error(s.T(), err)
		s.Equal("pg: no rows in result set", err.Error())
		require.Nil(s.T(), foundUser)
	})

	s.Run("Found", func() {
		userID := dao.UserID(1)
		s.deleteUser(userID)
		user := &dao.User{
			ID:        userID,
			NickName:  "nickName",
			FirstName: "firstName",
			LastName:  "lastName",
			ChatID:    2,
		}
		err := s.userDAO.Upsert(user)
		s.NoError(err)

		foundUser, err := s.userDAO.Find(user.ID)
		s.NoError(err)
		s.Equal(user.ID, foundUser.ID)
		s.Equal(user.FirstName, foundUser.FirstName)
		s.Equal(user.LastName, foundUser.LastName)
		s.Equal(user.NickName, foundUser.NickName)
		s.Equal(user.ChatID, foundUser.ChatID)
		s.NotEmpty(foundUser.CreatedAt)
		s.NotEmpty(foundUser.UpdatedAt)
		s.Equal(foundUser.CreatedAt.UnixNano(), foundUser.UpdatedAt.UnixNano())
	})
}

func (s *UserDAOTestSuite) TestUpsert() {
	s.Run("InsertNewFull", func() {
		userID := dao.UserID(1)
		s.deleteUser(userID)
		user := &dao.User{
			ID:        userID,
			NickName:  "nickName",
			FirstName: "firstName",
			LastName:  "lastName",
			ChatID:    dao.ChatID(1),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		err := s.userDAO.Upsert(user)

		s.NoError(err)
		s.Equal(userID, user.ID)
		s.Equal("firstName", user.FirstName)
		s.Equal("lastName", user.LastName)
		s.Equal("nickName", user.NickName)
		s.Equal(dao.ChatID(1), user.ChatID)
		s.NotEmpty(user.CreatedAt)
		s.NotEmpty(user.UpdatedAt)
	})

	s.Run("InsertNewEmpty", func() {
		userID := dao.UserID(1)
		s.deleteUser(userID)
		user := &dao.User{
			ID: userID,
		}

		err := s.userDAO.Upsert(user)

		s.NoError(err)
		s.Equal(userID, user.ID)
		s.Empty(user.FirstName)
		s.Empty(user.LastName)
		s.Empty(user.NickName)
		s.Empty(user.ChatID)
		s.NotEmpty(user.CreatedAt)
		s.NotEmpty(user.UpdatedAt)
	})

	s.Run("OverrideOnConflict", func() {
		userID := dao.UserID(1)
		s.deleteUser(userID)
		user := &dao.User{
			ID:        userID,
			NickName:  "nickName1",
			FirstName: "firstName1",
			LastName:  "lastName1",
			ChatID:    dao.ChatID(1),
		}
		err := s.userDAO.Upsert(user)
		s.NoError(err)
		user.NickName = "nickName2"
		user.FirstName = "firstName2"
		user.LastName = "lastName2"
		user.ChatID = dao.ChatID(2)

		err = s.userDAO.Upsert(user)
		s.NoError(err)

		s.NoError(err)
		s.Equal(userID, user.ID)
		s.Equal("firstName2", user.FirstName)
		s.Equal("lastName2", user.LastName)
		s.Equal("nickName2", user.NickName)
		s.Equal(dao.ChatID(2), user.ChatID)
		s.NotEmpty(user.CreatedAt)
		s.NotEmpty(user.UpdatedAt)
	})

	s.Run("FillOnConflict", func() {
		userID := dao.UserID(1)
		s.deleteUser(userID)
		user := &dao.User{
			ID:        userID,
			NickName:  "nickName1",
			FirstName: "firstName1",
			LastName:  "lastName1",
			ChatID:    dao.ChatID(1),
		}
		err := s.userDAO.Upsert(user)
		s.NoError(err)
		user.NickName = ""
		user.FirstName = ""
		user.LastName = ""
		user.ChatID = dao.ChatID(0)

		err = s.userDAO.Upsert(user)
		s.NoError(err)

		s.NoError(err)
		s.Equal(userID, user.ID)
		s.Equal("firstName1", user.FirstName)
		s.Equal("lastName1", user.LastName)
		s.Equal("nickName1", user.NickName)
		s.Equal(dao.ChatID(1), user.ChatID)
		s.NotEmpty(user.CreatedAt)
		s.NotEmpty(user.UpdatedAt)
	})
}
