package services

import (
	"context"
	pb "userManagement/genproto/UserManagementSevice/user"

	"github.com/jmoiron/sqlx"
)

type userManagementRepo struct {
	db *sqlx.DB
	pb.UnimplementedUserManagementServiceServer
}

func NewUserManagementRepo(db *sqlx.DB) *userManagementRepo {
	return &userManagementRepo{db: db}
}

func (u *userManagementRepo) GetUserById(ctx context.Context, in *pb.IdUserRequest) (*pb.UserResponse, error) {
	return nil, nil
}

func (u *userManagementRepo) UpdateUserById(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	return nil, nil
}

func (u *userManagementRepo) DeleteUserById(ctx context.Context, in *pb.IdUserRequest) (*pb.DeleteUserResponse, error) {
	return nil, nil
}

func (u *userManagementRepo) GetUserProfile(ctx context.Context, in *pb.IdUserRequest) (*pb.UserProfileResponse, error) {
	return nil, nil
}

func (u *userManagementRepo) UpdateUserProfile(ctx context.Context, in *pb.UserProfileRequest) (*pb.UserProfileResponse, error) {
	return nil, nil
}
