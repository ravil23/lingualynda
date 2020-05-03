package dao

import (
	"log"
	"time"

	"github.com/go-pg/pg/v9/orm"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/ravil23/lingualynda/telegrambot/postgres"
)

type UserID int

type User struct {
	tableName struct{} `pg:"user"`

	ID        UserID    `pg:"id,pk"`
	Nick      string    `pg:"nick,notnull"`
	FirstName string    `pg:"first_name"`
	LastName  string    `pg:"last_name"`
	ChatID    ChatID    `pg:"chat_id"`
	CreatedAt time.Time `pg:"created_at,default:now()"`
	UpdatedAt time.Time `pg:"updated_at,default:now()"`
}

func NewUser(tgUser *tgbotapi.User) *User {
	return &User{
		ID:        UserID(tgUser.ID),
		Nick:      tgUser.UserName,
		FirstName: tgUser.FirstName,
		LastName:  tgUser.LastName,
	}
}

type UserDAO interface {
	Find(userID UserID) (*User, error)
	Upsert(user *User) (*User, error)
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
	log.Printf("[user=%d] find user profile", userID)
	user := new(User)
	err := dao.conn.Model(user).
		Where("id = ?", userID).
		Select()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (dao *userDAO) Upsert(user *User) (*User, error) {
	log.Printf("[user=%d][chat=%d] upsert user profile", user.ID, user.ChatID)
	if _, err := dao.conn.Model(user).
		OnConflict("(id) DO NOTHING").
		Insert(); err != nil {
		return nil, err
	} else {
		return user, nil
	}
}
