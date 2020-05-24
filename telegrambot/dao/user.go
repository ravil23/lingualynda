package dao

import (
	"fmt"
	"time"

	"github.com/go-pg/pg/v9/orm"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/ravil23/lingualynda/telegrambot/postgres"
)

type UserID int

type User struct {
	tableName struct{} `pg:"user,alias:u"`

	ID        UserID    `pg:"id,pk"`
	NickName  string    `pg:"nick_name"`
	FirstName string    `pg:"first_name"`
	LastName  string    `pg:"last_name"`
	ChatID    ChatID    `pg:"chat_id"`
	CreatedAt time.Time `pg:"created_at,default:now()"`
	UpdatedAt time.Time `pg:"updated_at,default:now()"`
}

func (u *User) GetFormattedName() string {
	userString := fmt.Sprintf("[id=%d]", u.ID)
	if u.NickName != "" {
		userString = "@" + u.NickName + " " + userString
	}
	if u.LastName != "" {
		userString = u.LastName + " " + userString
	}
	if u.FirstName != "" {
		userString = u.FirstName + " " + userString
	}
	return userString
}

func NewUser(tgUser *tgbotapi.User) *User {
	return &User{
		ID:        UserID(tgUser.ID),
		NickName:  tgUser.UserName,
		FirstName: tgUser.FirstName,
		LastName:  tgUser.LastName,
	}
}

type UserDAO interface {
	Find(userID UserID) (*User, error)
	Upsert(user *User) error
	Delete(userID UserID) error
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
	return dao.conn.CreateTable((*User)(nil), options)
}

func (dao *userDAO) Find(userID UserID) (*User, error) {
	user := &User{ID: userID}
	err := dao.conn.Select(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (dao *userDAO) Delete(userID UserID) error {
	user := &User{ID: userID}
	return dao.conn.Delete(user)
}

func (dao *userDAO) Upsert(user *User) error {
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
