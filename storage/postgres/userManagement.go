package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	pb "userManagement/genproto/UserManagementSevice/user"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) GetUserById(ctx context.Context, user *pb.IdUserRequest) (*pb.UserResponse, error) {
	userResponse := &pb.UserResponse{}
	userResponse.UserId = user.UserId

	query := `
		SELECT username,
				email,
				created_at,
                updated_at,
		FROM users 
		WHERE user_id = $1 AND deleted_at IS NULL
	`

	stmt, err := repo.db.Prepare(query)

	if err != nil {
		return nil, fmt.Errorf("prepare error: %v", err)
	}

	row := stmt.QueryRowContext(ctx, stmt, user.UserId)
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

func (repo *UserRepository) GetAllUsers(ctx context.Context, user *pb.IdUserRequest) (*pb)
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
