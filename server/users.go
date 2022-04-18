package server

import (
	"context"

	"github.com/imrushi/go-grpc-assignment/data"
	protos "github.com/imrushi/go-grpc-assignment/protos/api/proto/v1/users"
	"github.com/sirupsen/logrus"
)

type Users struct {
	log  logrus.Logger
	user *data.UserDetail
	protos.UnimplementedUserDetailServer
}

func NewUsers(u *data.UserDetail, l logrus.Logger) *Users {
	c := &Users{l, u, protos.UnimplementedUserDetailServer{}}
	return c
}

func (u *Users) GetUser(ctx context.Context, ur *protos.UserDetailRequest) (*protos.UserDetailResponse, error) {
	u.log.Info("Handling GetUser Request for id: ", ur.GetId())
	// check if id exists or not
	foundIdx, err := data.FindIndexByUserID(ur.GetId())
	if err != nil {
		return nil, err
	}
	// Note: While retruning JSON it will not return false boolen field because it is
	// auto-genrated and it is decleared as omitempty
	return &protos.UserDetailResponse{Users: u.user.User[foundIdx]}, nil
}

func (u *Users) GetUsersList(ctx context.Context, ur *protos.ListOfUserDetailsRequest) (*protos.ListOfUserDetailsResponse, error) {
	u.log.Info("users: ", u.user.User[0])
	
	return &protos.ListOfUserDetailsResponse{Users: u.user.User}, nil
}
