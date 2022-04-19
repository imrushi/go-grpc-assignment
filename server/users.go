package server

import (
	"context"
	"fmt"

	"github.com/imrushi/go-grpc-assignment/data"
	protos "github.com/imrushi/go-grpc-assignment/protos/api/proto/v1/users"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		u.log.Error(err)
		return nil, err
	}
	// Note: While retruning JSON it will not return false boolen field because it is
	// auto-genrated and it is decleared as omitempty
	return &protos.UserDetailResponse{Users: u.user.User[foundIdx]}, nil
}

func (u *Users) GetUsersList(ctx context.Context, ur *protos.ListOfUserDetailsRequest) (*protos.ListOfUserDetailsResponse, error) {
	u.log.Info("Handling GetUsersList Request for ids: ", ur.GetIds())
	if ur.GetIds() == nil {
		err := status.Errorf(codes.InvalidArgument, "User list cann't be empty")
		return nil, err
	}

	var tempArr []*protos.User
	var invalidIds []int64
	for _, val := range ur.GetIds() {
		result, err := data.FindIndexByUserID(val)
		if err != nil {
			invalidIds = append(invalidIds, val)
			continue
		}
		tempArr = append(tempArr, u.user.User[result])
	}

	if len(invalidIds) != 0 {
		err := status.New(codes.NotFound, fmt.Sprintf("Ids doesn't exist %v in database", invalidIds))
		u.log.Errorf("Ids doesn't exist %v in database", invalidIds)
		return &protos.ListOfUserDetailsResponse{Users: tempArr, Error: err.Proto()}, nil
	}
	return &protos.ListOfUserDetailsResponse{Users: tempArr}, nil
}
