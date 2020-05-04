package dao_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/stretchr/testify/require"

	"github.com/ravil23/lingualynda/telegrambot/dao"
	"github.com/ravil23/lingualynda/telegrambot/postgres"
)

type QuestionDAOTestSuite struct {
	suite.Suite

	conn        *postgres.Connection
	questionDAO dao.QuestionDAO
}

func TestQuestionDAOTestSuite(t *testing.T) {
	suite.Run(t, new(QuestionDAOTestSuite))
}

func (s *QuestionDAOTestSuite) SetupTest() {
	s.conn = postgres.NewConnection()

	var err error
	s.questionDAO, err = dao.NewQuestionDAO(s.conn)
	s.NoError(err)
}

func (s *QuestionDAOTestSuite) TearDownTest() {
	err := s.conn.Close()
	s.NoError(err)
}

func (s *QuestionDAOTestSuite) deleteQuestion(questionID dao.QuestionID) {
	err := s.questionDAO.Delete(questionID)
	if err != nil {
		s.EqualError(err, "pg: no rows in result set")
	}
}

func (s *QuestionDAOTestSuite) TestFind() {
	s.Run("NotFound", func() {
		questionID := dao.QuestionID(1)
		s.deleteQuestion(questionID)

		foundQuestion, err := s.questionDAO.Find(questionID)

		require.Error(s.T(), err)
		s.Equal("pg: no rows in result set", err.Error())
		require.Nil(s.T(), foundQuestion)
	})

	s.Run("Found", func() {
		questionID := dao.QuestionID(1)
		s.deleteQuestion(questionID)
		question := &dao.Question{
			ID:   questionID,
			Text: "text",
			Options: []dao.Option{
				{Text: "textOption1", IsCorrect: true},
				{Text: "textOption2", IsCorrect: false},
			},
		}
		err := s.questionDAO.Upsert(question)
		s.NoError(err)

		foundQuestion, err := s.questionDAO.Find(question.ID)

		s.NoError(err)
		s.Equal(question.ID, foundQuestion.ID)
		s.Equal(question.Text, foundQuestion.Text)
		s.Len(foundQuestion.Options, len(question.Options))
		for i := range question.Options {
			s.Equal(question.Options[i].Text, foundQuestion.Options[i].Text)
			s.Equal(question.Options[i].IsCorrect, foundQuestion.Options[i].IsCorrect)
		}
		s.NotEmpty(foundQuestion.CreatedAt)
		s.NotEmpty(foundQuestion.UpdatedAt)
		s.Equal(foundQuestion.CreatedAt.UnixNano(), foundQuestion.UpdatedAt.UnixNano())
	})
}

func (s *QuestionDAOTestSuite) TestUpsert() {
	s.Run("InsertWithoutID", func() {
		question := &dao.Question{
			Text: "text",
			Options: []dao.Option{
				{Text: "textOption", IsCorrect: true},
			},
		}

		err := s.questionDAO.Upsert(question)

		s.NoError(err)
		s.NotEmpty(question.ID)
	})

	s.Run("InsertNewFull", func() {
		questionID := dao.QuestionID(1)
		s.deleteQuestion(questionID)
		question := &dao.Question{
			ID:   questionID,
			Text: "text",
			Options: []dao.Option{
				{Text: "textOption", IsCorrect: true},
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		err := s.questionDAO.Upsert(question)

		s.NoError(err)
		s.Equal(questionID, question.ID)
		s.Equal("text", question.Text)
		s.ElementsMatch([]dao.Option{{Text: "textOption", IsCorrect: true}}, question.Options)
		s.NotEmpty(question.CreatedAt)
		s.NotEmpty(question.UpdatedAt)
	})

	s.Run("InsertNewEmpty", func() {
		questionID := dao.QuestionID(1)
		s.deleteQuestion(questionID)
		question := &dao.Question{
			ID: questionID,
		}

		err := s.questionDAO.Upsert(question)

		s.EqualError(err, `ERROR #23502 null value in column "text" violates not-null constraint`)
	})

	s.Run("OverrideOnConflict", func() {
		questionID := dao.QuestionID(1)
		s.deleteQuestion(questionID)
		question := &dao.Question{
			ID:   questionID,
			Text: "text1",
			Options: []dao.Option{
				{Text: "textOption1", IsCorrect: true},
			},
		}
		err := s.questionDAO.Upsert(question)
		s.NoError(err)
		question.Text = "text2"
		question.Options = []dao.Option{{Text: "textOption2", IsCorrect: false}}

		err = s.questionDAO.Upsert(question)

		s.NoError(err)
		s.Equal(questionID, question.ID)
		s.Equal("text2", question.Text)
		s.ElementsMatch([]dao.Option{{Text: "textOption2", IsCorrect: false}}, question.Options)
		s.NotEmpty(question.CreatedAt)
		s.NotEmpty(question.UpdatedAt)
	})

	s.Run("FillOnConflict", func() {
		questionID := dao.QuestionID(1)
		s.deleteQuestion(questionID)
		question := &dao.Question{
			ID:   questionID,
			Text: "text1",
			Options: []dao.Option{
				{Text: "textOption1", IsCorrect: true},
			},
		}
		err := s.questionDAO.Upsert(question)
		s.NoError(err)
		question.Text = ""
		question.Options = []dao.Option{}

		err = s.questionDAO.Upsert(question)

		s.EqualError(err, `ERROR #23502 null value in column "text" violates not-null constraint`)
	})
}
