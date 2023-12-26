package usecase_test

import (
	domainUser "api/domain/model/user"
	mock_repository "api/domain/repository/mock"
	"api/usecase"
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func Test_NewUserUsecase(t *testing.T) {
	userRepository := &mock_repository.MockUserRepository{}
	userUsecase := usecase.NewUserUsecase(userRepository)
	if userUsecase == nil {
		t.Errorf("UserUsecase.NewUserUsecase: should return \"user\" usecase, but got: nil")
	}
}

func Test_FindAllUsers(t *testing.T) {
	type TestUserRepository struct {
		TestFindAll func() ([]*domainUser.User, error)
	}
	type arguments struct{}
	type wantError struct {
		isErr bool
		err   error
	}
	tests := []struct {
		name         string
		testUserRepo TestUserRepository
		args         arguments
		want         []*domainUser.User
		wantErr      wantError
	}{
		{
			name: "Successfully",
			testUserRepo: TestUserRepository{
				TestFindAll: func() ([]*domainUser.User, error) {
					users := make([]*domainUser.User, 3)
					for i := range users {
						users[i] = &domainUser.User{
							ID:             domainUser.UserID(fmt.Sprintf("%04d", i+1)),
							Name:           "test user " + strconv.Itoa(i+1),
							Administrative: false,
						}
					}
					return users, nil
				},
			},
			args: arguments{},
			want: []*domainUser.User{
				{ID: "0001", Name: "test user 1", Administrative: false},
				{ID: "0002", Name: "test user 2", Administrative: false},
				{ID: "0003", Name: "test user 3", Administrative: false},
			},
			wantErr: wantError{
				isErr: false,
				err:   nil,
			},
		},
		{
			name: "Error",
			testUserRepo: TestUserRepository{
				TestFindAll: func() ([]*domainUser.User, error) {
					return nil, fmt.Errorf("FAILED TO FIND ALL USERS.")
				},
			},
			args: arguments{},
			want: nil,
			wantErr: wantError{
				isErr: true,
				err:   fmt.Errorf("FAILED TO FIND ALL USERS."),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepository := &mock_repository.MockUserRepository{}
			userRepository.MockFindAllUsers = tt.testUserRepo.TestFindAll
			userUsecase := usecase.NewUserUsecase(userRepository)

			got, err := userUsecase.FindAllUsers()
			if tt.wantErr.isErr {
				if err.Error() != tt.wantErr.err.Error() {
					t.Errorf("UserUsecase.FIndAllUsers: error = %v, wantErr.err = %v", err, tt.wantErr.err)
				}
			}
			if (err != nil) != tt.wantErr.isErr {
				t.Errorf("UserUsecase.FindAllUsers: error = %v, wantErr.isErr = %v", err, tt.wantErr.isErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserUsecase.FindAllUsers: got = %v, tt.want = %v", got, tt.want)
			}
		})
	}
}
