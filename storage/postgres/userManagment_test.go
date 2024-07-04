package postgres

import (
	"context"
	"fmt"
	"log"
	// "reflect"

	// "log"
	"testing"
	"userManagement/config"
	pb "userManagement/genproto/UserManagementService"

	_ "github.com/lib/pq"

	// "userManagement/storage/postgres"

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
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	return db, nil
}

func ConnectDB() *UserRepository {
	db, err := sqlx.Open("postgres", "host=localhost user=postgres password=root port=5432 dbname=authentification sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return NewUserRepository(db)
}

func TestGetUserById(t *testing.T) {
	
	repo := ConnectDB()
	test := pb.IdUserRequest{
		UserId: "2a8c33b9-f3c0-4551-88e4-30e3957b331c",
	}

	ctx := context.Background()
	resp, err := repo.GetUserById(ctx, &test)
	if err != nil {
		t.Fatalf("Error getting user: %v", err)
	}

	Wait := pb.UserResponse{
		Username:  "test1221",
		Email:     "test1221@mail.com",
	}

	fmt.Println("response:  ",resp)
	fmt.Println("wait :   ",Wait)

	if Wait.Username != resp.Username || Wait.Email != resp.Email {
		t.Errorf("User data does not match. Wait: %+v, got: %+v", Wait, resp)
	}
}

func TestUpdateUserById(t *testing.T) {
	
	repo := ConnectDB()
	test := pb.IdUserRequest{
		UserId: "2a8c33b9-f3c0-4551-88e4-30e3957b331c",
	}

	ctx := context.Background()
	resp, err := repo.GetUserById(ctx, &test)
	if err != nil {
		t.Fatalf("Error getting user: %v", err)
	}

	//update qilingan malumotni qo'ying
	Wait := pb.UserResponse{
		Username:  "test1221",
		Email:     "test1221@mail.com",
	}

	fmt.Println("response:  ",resp)
	fmt.Println("wait :   ",Wait)

	if Wait.Username != resp.Username || Wait.Email != resp.Email {
		t.Errorf("User data does not match. Wait: %+v, got: %+v", Wait, resp)
	}
}

func TestDeleteUserById(t *testing.T) {

    repo := ConnectDB()
    test := pb.IdUserRequest{
        UserId: "2a8c33b9-f3c0-4551-88e4-30e3957b331c",
    }

    ctx := context.Background()
    massage, err := repo.DeleteUser(ctx, &test)
    if err!= nil {
        t.Fatalf("Error deleting user: %v", err)
    }

    Wait := pb.DeleteUserResponse{
        Message: "Deteted user",
    }

    fmt.Println("wait :   ",Wait)

    if Wait.Message!= massage.Message {
        t.Errorf("User data does not match. Wait: %+v, got: %+v", Wait, massage)
    }
}

// func TestGetUserProfileById(t *testing.T) {
// 	repo := ConnectDB()
// 	test := pb.IdUserRequest{
// 		UserId: "2a8c33b9-f3c0-4551-88e4-30e3957b331c",
// 	}

// 	ctx := context.Background()
// 	resp, err := repo.GetUserProfileById(ctx, &test)
// 	if err!= nil {
//         t.Fatalf("Error getting user profile: %v", err)
//     }

// 	Wait := pb.UserProfileResponse{
// 		FullName: "Javohir Abdusamatov",
// 		Bio: "bio for Abbos",
// 		Expertise: "beginner",
// 		Location: "nyu york",
//         AvatarUrl: "https://example.com/jack.jpg",
// 	}



// }