package handler_test

import (
	domainUser "api/domain/model/user"
	"api/interface/adapter/handler"
	mock_usecase "api/usecase/mock"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_NewUserHandler(t *testing.T) {
	userUsecase := &mock_usecase.MockUserUsecase{}
	userHandler := handler.NewUserHandler(userUsecase)
	if userHandler == nil {
		t.Errorf("UserHandler.NewUserHandler: should return \"user\" handler, but got nil.")
	}
}

func Test_FindAllUsers(t *testing.T) {
	type arguments struct{}
	type mock struct {
		ctx context.Context
	}
	type wantErr struct {
		isErr bool
		err   error
	}
	type want struct {
		statusCode int
		result     []*domainUser.User
		err        wantErr
	}
	tests := []struct {
		name            string
		testUserUsecase *mock_usecase.MockUserUsecase
		arguments       arguments
		mock            mock
		want            want
	}{
		{
			name: "Successfully",
			testUserUsecase: &mock_usecase.MockUserUsecase{
				MockFindAllUsers: func() ([]*domainUser.User, error) {
					users := make([]*domainUser.User, 3)
					for i := range users {
						users[i] = &domainUser.User{
							ID:   domainUser.UserID(fmt.Sprintf("%04d", i+1)),
							Name: "test user " + strconv.Itoa(i+1),
						}
					}
					return users, nil
				},
			},
			arguments: arguments{},
			mock: mock{
				ctx: context.Background(),
			},
			want: want{
				statusCode: http.StatusOK,
				result: []*domainUser.User{
					{ID: "0001", Name: "test user 1"},
					{ID: "0002", Name: "test user 2"},
					{ID: "0003", Name: "test user 3"},
				},
				err: wantErr{
					isErr: false,
					err:   nil,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userUsecase := &mock_usecase.MockUserUsecase{}
			userUsecase.MockFindAllUsers = tt.testUserUsecase.MockFindAllUsers
			userHandler := handler.NewUserHandler(userUsecase)

			gin.SetMode(gin.TestMode)
			g := gin.New()
			g.GET("/api/users", userHandler.GetUsers)

			w := httptest.NewRecorder()
			request, err := http.NewRequest("GET", "/api/users", nil)
			if err != nil {
				t.Fatalf("FAILED TO GENERATE HTTP REQUEST: %s", err)
			}
			g.ServeHTTP(w, request)

			if tt.want.statusCode == http.StatusOK {
				got, err := tt.testUserUsecase.FindAllUsers()
				if err != nil {
					t.Fatalf("FAILED TO CALL TestUserUsecase.TestFindAllUsers METHOD: %s", err.Error())
				}
				err = json.Unmarshal(w.Body.Bytes(), &got)
				if err != nil {
					t.Fatalf("FAILED TO UNMARSHAL JSON: %s", err.Error())
				}
				if !reflect.DeepEqual(got, tt.want.result) {
					t.Errorf("GET \"api/users\" RESPONSE %v, but want = %v", got, tt.want.result)
				}
			}
			if !reflect.DeepEqual(w.Code, tt.want.statusCode) {
				t.Errorf("STATUS CODE = %v, want = %v", w.Code, tt.want.statusCode)
			}
		})
	}
}
