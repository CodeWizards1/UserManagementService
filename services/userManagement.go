package services

import (
	"context"
	pb "userManagement/genproto/UserManagementService"
	"userManagement/storage/postgres"

	"github.com/jmoiron/sqlx"
)

type userManagementService struct {
	UserRepo *postgres.UserRepository
	pb.UnimplementedUserManagementServiceServer
}

func NewuserManagementService(db *sqlx.DB) *userManagementService {
	return &userManagementService{UserRepo: postgres.NewUserRepository(db)}
}

func (u *userManagementService) GetUserById(ctx context.Context, in *pb.IdUserRequest) (*pb.UserResponse, error) {

	res, err := u.UserRepo.GetUserById(ctx, in)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *userManagementService) UpdateUserById(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UserResponse, error) {

	res, err := u.UserRepo.UpdateUser(ctx, in)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *userManagementService) DeleteUserById(ctx context.Context, in *pb.IdUserRequest) (*pb.DeleteUserResponse, error) {
	return nil, nil
}

func (u *userManagementService) GetUserProfile(ctx context.Context, in *pb.IdUserRequest) (*pb.UserProfileResponse, error) {
	return nil, nil
}

func (u *userManagementService) UpdateUserProfile(ctx context.Context, in *pb.UserProfileRequest) (*pb.UserProfileResponse, error) {
	return nil, nil
}
