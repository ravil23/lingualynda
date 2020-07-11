package dao

import (
	"github.com/go-pg/pg/v9/orm"

	"github.com/ravil23/lingualynda/telegrambot/entity"
	"github.com/ravil23/lingualynda/telegrambot/postgres"
)

type UserDAO interface {
	Find(userID entity.UserID) (*entity.User, error)
	FindAll() ([]*entity.User, error)
	Upsert(user *entity.User) error
}

var _ UserDAO = (*userDAO)(nil)

type userDAO struct {
	conn *postgres.Connection
}

func NewUserDAO(conn *postgres.Connection) (*userDAO, error) {
	dao := &userDAO{
		conn: conn,
	}
	if err := dao.ensureSchema(); err != nil {
		return nil, err
	}
	return dao, nil
}

func (dao *userDAO) ensureSchema() error {
	options := &orm.CreateTableOptions{
		IfNotExists:   true,
		FKConstraints: true,
	}
	return dao.conn.CreateTable((*entity.User)(nil), options)
}

func (dao *userDAO) Find(userID entity.UserID) (*entity.User, error) {
	user := &entity.User{ID: userID}
	err := dao.conn.Select(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (dao *userDAO) FindAll() ([]*entity.User, error) {
	var users []*entity.User
	err := dao.conn.Model(&users).
		Select()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (dao *userDAO) Upsert(user *entity.User) error {
	_, err := dao.conn.Model(user).
		OnConflict("(id) DO UPDATE").
		Set("updated_at = now()").
		Set("nick_name = coalesce(EXCLUDED.nick_name, u.nick_name)").
		Set("first_name = coalesce(EXCLUDED.first_name, u.first_name)").
		Set("last_name = coalesce(EXCLUDED.last_name, u.last_name)").
		Set("chat_id = coalesce(EXCLUDED.chat_id, u.chat_id)").
		Insert()
	return err
}
