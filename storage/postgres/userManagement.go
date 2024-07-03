package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
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
	row := repo.db.DB.QueryRow(query, user.UserId)

	fmt.Println("user id: ", user.UserId)
	if err := row.Scan(
		&userResponse.Username,
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

	var (
		params = make(map[string]interface{})
		arr    []interface{}
	)

	query := `UPDATE users SET updated_at = NOW()`

	if len(user.Username) > 0 {
		params["username"] = user.Username
		query += ", username = :username"
	}

	if len(user.Email) > 0 {
		params["email"] = user.Email
		query += ", email = :email"
	}

	if len(user.Password) > 0 {
		params["password"] = user.Password
		query += ", password = :password"
	}

	query, arr = ReplaceQueryParamsUser(query, params)
	arr = append(arr, user.UserId) // userId ni qo'shish
	fil := fmt.Sprintf("WHERE user_id = $%d and deleted_at is NULL", len(arr))
	query += fil

	_, err := repo.db.ExecContext(ctx, query, arr...)
	if err != nil {
		return nil, err
	}

	return repo.GetUserById(ctx, &pb.IdUserRequest{UserId: user.UserId})
}

func ReplaceQueryParamsUser(namedQuery string, params map[string]interface{}) (string, []interface{}) {
	var (
		i    int = 1
		args []interface{}
	)

	for k, v := range params {
		if k != "" && strings.Contains(namedQuery, ":"+k) {
			namedQuery = strings.ReplaceAll(namedQuery, ":"+k, "$"+strconv.Itoa(i))
			args = append(args, v)
			i++
		}
	}

	return namedQuery, args
}

func (repo *UserRepository) DeleteUser(ctx context.Context, user *pb.IdUserRequest) (*pb.DeleteUserResponse, error) {
	query := `
        UPDATE users SET deleted_at = NOW() 
        WHERE user_id = $1 AND deleted_at IS NULL
    `
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("prepare error: %v", err)
	}

	if _, err := stmt.ExecContext(ctx, user.UserId); err != nil {
		return nil, fmt.Errorf("exec error: %v", err)
	}

	return &pb.DeleteUserResponse{Message: "Deteted user"}, nil
}

func (repo *UserRepository) GetUserProfileById(ctx context.Context, user *pb.IdUserRequest) (*pb.UserProfileResponse, error) {
	userProfileResponse := &pb.UserProfileResponse{}
	userProfileResponse.UserId = user.UserId

	query := `
    SELECT user_id, 
    full_name, 
    bio, 
    expertise, 
    location, 
    avatar_url 
    FROM user_profiles
    WHERE user_id = $1
    `

	fmt.Println("user id: ", userProfileResponse.UserId)

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("prepare error: %v", err)
	}
	defer stmt.Close() // Ensure the statement is closed after use

	row := stmt.QueryRowContext(ctx, user.UserId)
	if err := row.Scan(
		&userProfileResponse.UserId,
		&userProfileResponse.FullName,
		&userProfileResponse.Bio,
		&userProfileResponse.Expertise,
		&userProfileResponse.Location,
		&userProfileResponse.AvatarUrl); err != nil {

		fmt.Println("user id: ", userProfileResponse.UserId)
		fmt.Println("user profile expertise: ", userProfileResponse.Expertise)
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user profile not found")
		}
		return nil, fmt.Errorf("scan error: %v", err)
	}

	fmt.Println("user profile: ", userProfileResponse)

	return userProfileResponse, nil
}

func (repo *UserRepository) UpdateUserProfileById(ctx context.Context, user *pb.UserProfile) (*pb.UserProfileResponse, error) {

	var (
		params = make(map[string]interface{})
		arr    []interface{}
	)

	query := `UPDATE user_profiles SET updated_at = NOW()`

	if len(user.FullName) > 0 {
		params["full_Name"] = user.FullName
		query += ", full_Name = :full_Name"
	}

	if len(user.Bio) > 0 {
		params["bio"] = user.Bio
		query += ", bio = :bio"
	}

	if len(user.Expertise) > 0 {
		params["expertise"] = user.Expertise
		query += ", expertise = :expertise"
	}

	query, arr = ReplaceQueryParamsUser(query, params)
	arr = append(arr, user.UserId) // userId ni qo'shish
	fil := fmt.Sprintf("WHERE user_id = $%d and deleted_at is NULL", len(arr))
	fmt.Println(query)
	fmt.Println(fil)
	fmt.Println(user.UserId)
	fmt.Println(arr)
	query += fil

	_, err := repo.db.ExecContext(ctx, query, arr...)
	if err != nil {
		return nil, err
	}

	return repo.GetUserProfileById(ctx, &pb.IdUserRequest{UserId: user.UserId})
}
