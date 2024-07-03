package services

import (
	"context"
	"fmt"
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
	res, err := u.UserRepo.DeleteUser(ctx, in)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func (u *userManagementService) GetUserProfileById(ctx context.Context, in *pb.IdUserRequest) (*pb.UserProfileResponse, error) {
	fmt.Printf("Received request for user ID: %s\n", in.UserId)

	res, err := u.UserRepo.GetUserProfileById(ctx, in)
	if err != nil {
		fmt.Printf("Error fetching user profile: %v\n", err)
		return nil, err
	}

	fmt.Printf("Fetched user profile: %v\n", res)
	return res, nil
}

func (u *userManagementService) UpdateUserProfileById(ctx context.Context, in *pb.UserProfile) (*pb.UserProfileResponse, error) {
	
	res, err := u.UserRepo.UpdateUserProfileById(ctx, in)
	if err != nil {
		return nil, err
	}

	return res, nil
}
