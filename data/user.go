package data

import (
	protos "github.com/imrushi/go-grpc-assignment/protos/api/proto/v1/users"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var userList = []*protos.User{
	&protos.User{
		Id:      1,
		Fname:   "Tony",
		City:    "Maliby",
		Phone:   1234567890,
		Height:  5.67,
		Married: true,
	},
	&protos.User{
		Id:      2,
		Fname:   "Thor",
		City:    "Asgard",
		Phone:   1234567890,
		Height:  6.67,
		Married: false,
	},
	&protos.User{
		Id:      3,
		Fname:   "Peter",
		City:    "New York",
		Phone:   1234567890,
		Height:  5.30,
		Married: true,
	},
}

type UserDetail struct {
	log  logrus.Logger
	User []*protos.User
}

func NewUserDetail(l logrus.Logger) *UserDetail {
	ud := &UserDetail{log: l, User: userList}

	return ud
}

func FindIndexByUserID(id int64) (int, error) {
	result, err := OrderAgnosticBS(userList, id)
	if err != nil {
		return -1, err
	}

	return result, nil
}

func OrderAgnosticBS(arr []*protos.User, target int64) (int, error) {
	var start int = 0
	var end int = len(arr) - 1

	// find whether the array is sorted in ascending or descending
	var isAsc bool = arr[start].Id < arr[end].Id

	for start <= end {
		// find the middle element
		// mid  := (start + end) / 2 // If we use this formula if start and end value or sum of both will execed the limit of int range
		mid := start + (end-start)/2

		if arr[mid].Id == target {
			return mid, nil
		}

		if isAsc {

			if target < arr[mid].Id {
				//check on left
				end = mid - 1
			} else if target > arr[mid].Id {
				// check on right
				start = mid + 1
			}
		} else {
			if target > arr[mid].Id {
				//check on right
				end = mid - 1
			} else if target < arr[mid].Id {
				// check on left
				start = mid + 1
			}
		}
	}
	err := status.Errorf(
		codes.NotFound,
		"User with id %d doesn't exist in Database",
		target,
	)
	return -1, err
}
