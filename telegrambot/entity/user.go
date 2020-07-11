package entity

import (
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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
