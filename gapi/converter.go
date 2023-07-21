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
		Phone:             user.Phone,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
		Role:              int32(user.Role),
	}
}

func convertType(t model.Type) *pb.Type {
	return &pb.Type{
		Id:   t.Id,
		Name: t.Name,
	}
}

func convertTypes(types []model.Type) []*pb.Type {
	result := make([]*pb.Type, 0, len(types))
	for _, t := range types {
		result = append(result, convertType(t))
	}
	return result
}

func convertPackage(pkg model.Package) *pb.Package {
	return &pb.Package{
		Id:           pkg.ID,
		DesId:        pkg.DesID,
		Name:         pkg.Name,
		Details:      pkg.Details,
		Price:        pkg.Price,
		ImgUrl:       pkg.ImgURL,
		Duration:     pkg.Duration,
		NumberPeople: pkg.NumberPeople,
		Types:        convertTypes(pkg.Types),
	}
}

func convertPackages(packages []model.Package) []*pb.Package {
	result := make([]*pb.Package, 0, len(packages))
	for _, p := range packages {
		result = append(result, convertPackage(p))
	}
	return result
}

func convertTrip(trip model.Trip) *pb.Trip {
	return &pb.Trip{
		Id:        trip.Id,
		Username:  trip.Username,
		PId:       trip.PId,
		Total:     trip.Total,
		Notes:     trip.Notes,
		StartDate: timestamppb.New(trip.StartDate),
	}
}
