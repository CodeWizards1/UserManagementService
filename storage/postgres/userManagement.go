package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	pb "userManagement/genproto/UserManagementService"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) GetUserById(ctx context.Context, user *pb.IdUserRequest) (*pb.UserResponse, error) {
	userResponse := &pb.UserResponse{}
	userResponse.UserId = user.UserId

	query := `
		SELECT username,
				email,
				created_at,
                updated_at
		FROM users 
		WHERE user_id = $1 AND deleted_at IS NULL
	`

	stmt, err := repo.db.Prepare(query)

	if err != nil {
		return nil, fmt.Errorf("prepare error: %v", err)
	}

	row := stmt.QueryRowContext(ctx, user.UserId)
	if err := row.Scan(&userResponse.Username,
		&userResponse.Email,
		&userResponse.CreatedAt,
		&userResponse.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("scan error: %v", err)
	}
	return userResponse, nil
}

func (repo *UserRepository) UpdateUser(ctx context.Context, user *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	query := `
		UPDATE user SET updated_at = NOW() 
	`
	params := []string{}
	args := []interface{}{}

	if user.Username != "" {
		params = append(params, fmt.Sprintf("username = $%d", len(args)+1))
		args = append(args, user.Username)
	}

	if user.Email != "" {
		params = append(params, fmt.Sprintf("email = $%d", len(args)+1))
		args = append(args, user.Email)
	}

	if user.Password != "" {
		params = append(params, fmt.Sprintf("password = $%d", len(args)+1))
		args = append(args, user.Password)
	}

	if len(params) > 0 {
		query += strings.Join(params, ", ")
	}

	query += " WHERE user_id = $1 AND deleted_at IS NULL"

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("prepare error: %v", err)
	}

	if _, err := stmt.ExecContext(ctx, user.UserId, args); err != nil {
		return nil, fmt.Errorf("exec error: %v", err)
	}

	return repo.GetUserById(ctx, &pb.IdUserRequest{UserId: user.UserId})
}

func (repo *UserRepository) DeleteUser(ctx context.Context, user *pb.IdUserRequest) (*pb.DeleteUserResponse, error) {
	query := `
        UPDATE users SET deleted_at = NOW() 
        WHERE user_id = $1 AND deleted_at IS NULL
    `
    stmt, err := repo.db.Prepare(query)
    if err!= nil {
        return nil, fmt.Errorf("prepare error: %v", err)
    }

    if _, err := stmt.ExecContext(ctx, user.UserId); err!= nil {
        return nil, fmt.Errorf("exec error: %v", err)
    }

    return &pb.DeleteUserResponse{Message: "Deteted user"}, nil
}

func (repo *UserRepository) GetUserProfileById(ctx context.Context, user *pb.IdUserRequest) (*pb.UserProfileResponse, error) {
	userProfileResponse := &pb.UserProfileResponse{}
	userProfileResponse.UserId = user.UserId
	userProfile := &pb.UserProfile{}

	query := `
	    SELECT user_id, 
			full_name, 
			bio, 
			user_expertise, 
			location, 
			avatar_url 
		FROM user_profiles
		where user_id = $1
		`
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("prepare error: %v", err)
	}

	row := stmt.QueryRowContext(ctx, user.UserId)
	if err := row.Scan(
		&userProfileResponse.UserId,
		&userProfile.FullName,
		&userProfile.Bio,
        &userProfile.Expertise,
        &userProfile.Location,
        &userProfile.AvatarUrl,
		); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user profile not found")
		}
		return nil, fmt.Errorf("scan error: %v", err)
	}
	userProfileResponse.UserProfile = userProfile

	return userProfileResponse, nil

}

func (repo *UserRepository) UpdateUserProfile(ctx context.Context, user *pb.UserProfileRequest) (*pb.UserProfileResponse, error) {

	userProfile := &pb.UserProfile{}

	query := `
		UPDATE user SET updated_at = NOW() 
	`
	params := []string{}
	args := []interface{}{}

	if userProfile.FullName != "" {
		params = append(params, fmt.Sprintf("username = $%d", len(args)+1))
		args = append(args, userProfile.FullName)
	}

	if userProfile.Bio != "" {
		params = append(params, fmt.Sprintf("email = $%d", len(args)+1))
		args = append(args, userProfile.Bio)
	}

	if userProfile.Expertise != "" {
		params = append(params, fmt.Sprintf("password = $%d", len(args)+1))
		args = append(args, userProfile.Expertise)
	}

	if userProfile.Location!= "" {
        params = append(params, fmt.Sprintf("location = $%d", len(args)+1))
        args = append(args, userProfile.Location)
    }

	if userProfile.AvatarUrl!= "" {
        params = append(params, fmt.Sprintf("avatar_url = $%d", len(args)+1))
        args = append(args, userProfile.AvatarUrl)
    }

	if len(params) > 0 {
		query += strings.Join(params, ", ")
	}

	query += " WHERE user_id = $1 AND deleted_at IS NULL"

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("prepare error: %v", err)
	}

	if _, err := stmt.ExecContext(ctx, user.UserId, args); err != nil {
		return nil, fmt.Errorf("exec error: %v", err)
	}

	return repo.GetUserProfileById(ctx, &pb.IdUserRequest{UserId: user.UserId})
}
