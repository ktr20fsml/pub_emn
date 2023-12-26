package user

import (
	domain "api/domain/model/user"
	"time"
)

type User struct {
	ID             string    `db:"mst_user_id"`
	Name           string    `db:"user_name"`
	Password       string    `db:"user_password"`
	Administrative bool      `db:"administrative"`
	StopUsing      time.Time `db:"stop_using"`
	CreatedAt      time.Time `db:"created_at"`
	CreatedBy      string    `db:"created_by"`
}

func ConvertToUserData(reqUser *domain.User) *User {
	return &User{
		ID:             string(reqUser.ID),
		Name:           reqUser.Name,
		Password:       reqUser.Password,
		Administrative: reqUser.Administrative,
		CreatedBy:      string(reqUser.CreatedBy),
	}
}

func ConvertToUsersDatas(reqUsers []*domain.User) []*User {
	users := make([]*User, len(reqUsers))

	for i, reqUser := range reqUsers {
		users[i] = ConvertToUserData(reqUser)
	}

	return users
}

func ConvertToUserDomain(user *User) *domain.User {
	resUser := &domain.User{
		ID:             domain.UserID(user.ID),
		Name:           user.Name,
		Password:       user.Password,
		Administrative: user.Administrative,
		StopUsing:      user.StopUsing,
		CreatedAt:      user.CreatedAt,
		CreatedBy:      domain.UserID(user.ID),
	}

	return resUser
}

func ConvertToUsersDomains(users []*User) []*domain.User {
	resUsers := make([]*domain.User, len(users))

	for i, user := range users {
		resUsers[i] = ConvertToUserDomain(user)
	}

	return resUsers
}
