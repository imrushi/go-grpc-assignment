package server

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/imrushi/go-grpc-assignment/data"
	"github.com/imrushi/go-grpc-assignment/protos/api/proto/v1/users"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGetUser(t *testing.T) {
	log := &logrus.Logger{
		Out: os.Stdout,
		Formatter: &logrus.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
		},
		Level: logrus.DebugLevel,
	}

	user := data.NewUserDetail(*log)
	u := NewUsers(user, *log)

	t.Run("User exist", func(t *testing.T) {
		res, err := u.GetUser(context.Background(), &users.UserDetailRequest{Id: 1})
		if err != nil {
			log.Error(err)
		}
		got := res.String()
		expected := `users:{id:1  fname:"Tony"  city:"Maliby"  phone:1234567890  height:5.67  Married:true}`
		assert.Equal(t, expected, got)
	})

	t.Run("User id does not exist", func(t *testing.T) {
		res, err := u.GetUser(context.Background(), &users.UserDetailRequest{Id: 4})
		if err != nil {
			log.Error(err)
			got := err
			expected := status.Errorf(
				codes.NotFound,
				"User with id %d doesn't exist in Database",
				4,
			)
			assert.Equal(t, expected, got)
		}
		t.Log(res.String())

		got1 := res
		assert.Empty(t, got1)
	})
}

func TestGetUsersList(t *testing.T) {
	log := &logrus.Logger{
		Out: os.Stdout,
		Formatter: &logrus.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
		},
		Level: logrus.DebugLevel,
	}

	user := data.NewUserDetail(*log)
	u := NewUsers(user, *log)

	t.Run("User List exist", func(t *testing.T) {

		res, err := u.GetUsersList(context.Background(), &users.ListOfUserDetailsRequest{Ids: []int64{1, 2, 3}})
		if err != nil {
			log.Error(err)
		}
		js, _ := json.Marshal(res)
		got := string(js)
		expected := `{"users":[{"id":1,"fname":"Tony","city":"Maliby","phone":1234567890,"height":5.67,"Married":true},{"id":2,"fname":"Thor","city":"Asgard","phone":1234567890,"height":6.67},{"id":3,"fname":"Peter","city":"New York","phone":1234567890,"height":5.3,"Married":true}]}`
		assert.Equal(t, expected, got)
	})

	t.Run("List of user id does not exist", func(t *testing.T) {
		_, err := u.GetUsersList(context.Background(), &users.ListOfUserDetailsRequest{Ids: []int64{4, 5, 6}})
		if err != nil {
			log.Error(err)
			got := err
			expected := status.New(codes.NotFound, fmt.Sprintf("Ids doesn't exist %v in database", []int64{4, 5, 6}))
			assert.Equal(t, expected.Proto(), got)
		}
	})

	t.Run("List of user valid and invalid ids", func(t *testing.T) {
		res, _ := u.GetUsersList(context.Background(), &users.ListOfUserDetailsRequest{Ids: []int64{1, 2, 4}})
		js, _ := json.Marshal(res)
		got := string(js)
		expected := `{"users":[{"id":1,"fname":"Tony","city":"Maliby","phone":1234567890,"height":5.67,"Married":true},{"id":2,"fname":"Thor","city":"Asgard","phone":1234567890,"height":6.67}],"error":{"code":5,"message":"Ids doesn't exist [4] in database"}}`
		assert.Equal(t, expected, got)
	})

	t.Run("List of empty user list", func(t *testing.T) {
		_, err := u.GetUsersList(context.Background(), &users.ListOfUserDetailsRequest{Ids: nil})
		t.Log(err)
		if err != nil {
			got := err
			expected := status.Errorf(codes.InvalidArgument, "User list cann't be empty")
			assert.Equal(t, expected, got)
		}
	})
}
