package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
	pb "userManagement/genproto/UserManagementService"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func checkPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func (repo *UserRepository) CreateUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	query := `
		INSERT INTO users 
			(username, email, password) 
		VALUES 
			($1, $2, $3)
		RETURNING 
			user_id, 
			username, 
			email,
			created_at,
			updated_at
	`
	var (
		UserResponse pb.UserResponse
		created_at   time.Time
		updated_at   time.Time
	)

	hashedPassword, err := hashPassword(in.Password)
	if err != nil {
		log.Println("postgres/user Error while hashing password", err)
		return nil, err
	}

	row := repo.db.QueryRowxContext(ctx, query, in.Username, in.Email, hashedPassword)
	err = row.Scan(&UserResponse.UserId, &UserResponse.Username, &UserResponse.Email, &created_at, &updated_at)
	if err != nil {
		log.Println("postgres/user Error while scanning UserResponse", err)
		return nil, err
	}
	UserResponse.CreatedAt = created_at.Format("2006-01-02 15:04:05")
	UserResponse.UpdatedAt = updated_at.Format("2006-01-02 15:04:05")

	query = `
		INSERT INTO user_profiles 
			(user_id)
		VALUES
			($1)
	`

	_, err = repo.db.ExecContext(ctx, query, UserResponse.UserId)
	if err != nil {
		log.Println("postgres/user Error while creating user_profiles", err)
		return nil, err
	}

	return &UserResponse, nil
}

func (repo *UserRepository) Login(ctx context.Context, in *pb.AutorizationRequest) (*pb.AutorizationResponse, error) {
	query := `
		SELECT
			email,
			password
		FROM users
		WHERE
			email = $1 
	`
	var (
		email    string
		password string
	)

	row := repo.db.QueryRowxContext(ctx, query, in.Email)
	err := row.Scan(&email, &password)
	if err != nil {
		log.Println("postgres/user", err)
		return nil, err
	}

	if !checkPassword(password, in.Password) {
		return &pb.AutorizationResponse{Message: "Password is incorrect!"}, nil
	}

	return &pb.AutorizationResponse{Message: "Login successful"}, nil
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

func (repo *UserRepository) UpdateUserProfileById(ctx context.Context, user *pb.UpdateUserProfileRequest) (*pb.UserProfileResponse, error) {

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
