package dao

import (
	"time"

	"github.com/go-pg/pg/v9/orm"

	"github.com/ravil23/lingualynda/telegrambot/entity"
	"github.com/ravil23/lingualynda/telegrambot/postgres"
)

type UserMemorizedTermDAO interface {
	FindByUserID(userID entity.UserID, from time.Time) ([]*entity.UserMemorizedTerm, error)
	Upsert(userMemorizedTerm *entity.UserMemorizedTerm) error
}

var _ UserMemorizedTermDAO = (*userMemorizedTermDAO)(nil)

type userMemorizedTermDAO struct {
	conn *postgres.Connection
}

func NewUserMemorizedTermDAO(conn *postgres.Connection) (*userMemorizedTermDAO, error) {
	dao := &userMemorizedTermDAO{
		conn: conn,
	}
	if err := dao.ensureSchema(); err != nil {
		return nil, err
	}
	return dao, nil
}

func (dao *userMemorizedTermDAO) ensureSchema() error {
	options := &orm.CreateTableOptions{
		IfNotExists:   true,
		FKConstraints: true,
	}
	return dao.conn.CreateTable((*entity.UserMemorizedTerm)(nil), options)
}

func (dao *userMemorizedTermDAO) FindByUserID(userID entity.UserID, from time.Time) ([]*entity.UserMemorizedTerm, error) {
	var userMemorizedTerms []*entity.UserMemorizedTerm
	err := dao.conn.Model(&userMemorizedTerms).
		Where("user_id = ?", userID).
		Where("timestamp >= ?", from).
		Select()
	if err != nil {
		return nil, err
	}
	return userMemorizedTerms, nil
}

func (dao *userMemorizedTermDAO) Upsert(userMemorizedTerm *entity.UserMemorizedTerm) error {
	_, err := dao.conn.Model(userMemorizedTerm).
		OnConflict("(timestamp, user_id, term) DO NOTHING").
		Insert()
	return err
}
