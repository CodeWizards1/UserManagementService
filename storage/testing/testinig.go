package postgres

import (
	"context"
	"fmt"
	"log"
	"testing"
	"userManagement/config"
	pb "userManagement/genproto/UserManagementService"

	"userManagement/storage/postgres"

	"github.com/jmoiron/sqlx"
)

func GetDB(path string) (*sqlx.DB, error) {
	cfg := config.Load(path)

	psqlUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.DbHost,
		cfg.Postgres.DbPort,
		cfg.Postgres.DbUser,
		cfg.Postgres.DbPassword,
		cfg.Postgres.DbName,
	)

	db, err := sqlx.Connect("postgres", psqlUrl)
	return db, err
}

func TestPostgres(t *testing.T) {

	db, err := GetDB(".")
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	repo := postgres.NewUserRepository(db)

	test := pb.IdUserRequest{
		UserId: "2a8c33b9-f3c0-4551-88e4-30e3957b331c",
	}

	ctx := context.Background()
	resp, err := repo.GetUserById(ctx, &test)
	if err != nil {
		t.Fatalf("Error getting user: %v", err)
	}
	fmt.Println(resp)

	wait := pb.UserResponse{
		Username: "Ikrom",
		Email: "ikrom@gmail.com",
		CreatedAt: "2024-07-02 15:34:56.768968+05",
		UpdatedAt: "2024-07-02 15:34:56.768968+05",
	}
	if resp.Username!= wait.Username || resp.CreatedAt!= wait.CreatedAt || resp.UpdatedAt!= wait.UpdatedAt {
        t.Errorf("User not found or data not match. Expected: %+v, got: %+v", wait, resp)
		return
    }


}
