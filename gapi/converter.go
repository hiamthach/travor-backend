package gapi

import (
	"github.com/travor-backend/model"
	"github.com/travor-backend/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user model.User) *pb.User {
	return &pb.User{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}
