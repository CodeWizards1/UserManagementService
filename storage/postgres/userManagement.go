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
	fil := " WHERE user_id = $4"

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
	fil = fmt.Sprintf("WHERE user_id = $%d", len(arr))
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

// userResponse := &pb.UserResponse{}
// 	userResponse.UserId = user.UserId

// 	query := `
// 		SELECT username,
// 				email,
// 				created_at,
// 				updated_at
// 		FROM users
// 		WHERE user_id = $1 AND deleted_at IS NULL
// 	`
// 	row := repo.db.DB.QueryRow(query, user.UserId)

// 	fmt.Println("user id: ", user.UserId)
// 	if err := row.Scan(
// 		&userResponse.Username,
// 		&userResponse.Email,
// 		&userResponse.CreatedAt,
// 		&userResponse.UpdatedAt); err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, fmt.Errorf("user not found")
// 		}
// 		return nil, fmt.Errorf("scan error: %v", err)
// 	}

// 	return userResponse, nil

// func (repo *UserRepository) GetUserProfileById(ctx context.Context, user *pb.IdUserRequest) (*pb.UserProfileResponse, error) {
// 	userProfileResponse := &pb.UserProfileResponse{}
// 	userProfileResponse.UserId = user.UserId

// 	fmt.Println("user id: ", userProfileResponse.UserId)
// 	query := `
// 	    SELECT user_id,
// 			full_name,
// 			bio,
// 			user_expertise,
// 			location,
// 			avatar_url
// 		FROM user_profiles
// 		where user_id = $1
// 		`
// 	// stmt, err := repo.db.Prepare(query)
// 	// if err != nil {
// 	// 	return nil, fmt.Errorf("prepare error: %v", err)
// 	// }

// 	// row := stmt.QueryRowContext(ctx, user.UserId)

// 	fmt.Println("user id : ", user.UserId)
// 	row := repo.db.DB.QueryRow(query, user.UserId)
// 	if err := row.Scan(
// 		&userProfileResponse.UserId,
// 		&userProfileResponse.FullName,
// 		&userProfileResponse.Bio,
// 		&userProfileResponse.Expertise,
// 		&userProfileResponse.Location,
// 		&userProfileResponse.AvatarUrl); err != nil {

// 		fmt.Println("user profile: ", userProfileResponse)
// 		if err == sql.ErrNoRows {
// 			return nil, fmt.Errorf("user profile not found")
// 		}
// 		return nil, fmt.Errorf("scan error: %v", err)
// 	}
// 	// userProfileResponse.UserProfile = userProfile

// 	return userProfileResponse, nil

// }
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

func (repo *UserRepository) UpdateUserProfile(ctx context.Context, user *pb.UserProfileRequest) (*pb.UserProfileResponse, error) {

	userProfileResponse := &pb.UserProfileResponse{}

	query := `
		UPDATE user SET updated_at = NOW() 
	`
	params := []string{}
	args := []interface{}{}

	if userProfileResponse.FullName != "" {
		params = append(params, fmt.Sprintf("username = $%d", len(args)+1))
		args = append(args, userProfileResponse.FullName)
	}

	if userProfileResponse.Bio != "" {
		params = append(params, fmt.Sprintf("email = $%d", len(args)+1))
		args = append(args, userProfileResponse.Bio)
	}

	if userProfileResponse.Expertise != "" {
		params = append(params, fmt.Sprintf("password = $%d", len(args)+1))
		args = append(args, userProfileResponse.Expertise)
	}

	if userProfileResponse.Location != "" {
		params = append(params, fmt.Sprintf("location = $%d", len(args)+1))
		args = append(args, userProfileResponse.Location)
	}

	if userProfileResponse.AvatarUrl != "" {
		params = append(params, fmt.Sprintf("avatar_url = $%d", len(args)+1))
		args = append(args, userProfileResponse.AvatarUrl)
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
