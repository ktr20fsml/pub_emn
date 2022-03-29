package user

import (
	"api/status"
	"time"
	"unicode/utf8"
)

type User struct {
	ID             UserID    `json:"userID"`
	Name           string    `json:"userName"`
	Password       string    `json:"userPassword,omitempty"`
	Administrative bool      `json:"userAdministrative"`
	StopUsing      time.Time `json:"userStopUsing,omitempty"`
	CreatedAt      time.Time `json:"userCreatedAt,omitempty"`
	CreatedBy      UserID    `json:"userCreatedBy,omitempty"`
}

type UserID string

var (
	errID   string = "User ID is 8 characters only."
	errName string = "User name is within 2 to 16 characters."
)

func Validate(reqUser *User) *status.Statuses {
	stts := &status.Statuses{}
	// validate user id
	if utf8.RuneCountInString(string(reqUser.ID)) != 8 {
		stts.Message = append(stts.Message, &errID)
	}

	// validate user name
	if utf8.RuneCountInString(reqUser.Name) < 2 {
		stts.Message = append(stts.Message, &errName)
	}
	if utf8.RuneCountInString(reqUser.Name) > 16 {
		stts.Message = append(stts.Message, &errName)
	}

	return stts
}

func ResponseUser(user *User) *User {
	resUser := &User{
		ID:             user.ID,
		Name:           user.Name,
		Administrative: user.Administrative,
		StopUsing:      user.StopUsing,
	}

	return resUser
}

func ResponseUsers(users []*User) []*User {
	resUsers := make([]*User, len(users))

	for i, user := range users {
		resUsers[i] = ResponseUser(user)
	}

	return resUsers
}
