package main

import (
	"context"
	"ent-demo/ent"
	"ent-demo/ent/migrate"
	"ent-demo/ent/user"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := ent.Open("mysql", "root:@tcp(localhost:3306)/ent_db?parseTime=True")

	if err != nil {
        log.Fatalf("failed connecting to mysql: %v", err)
    }
    defer client.Close()
    ctx := context.Background()
    // Run migration.
    err = client.Schema.Create(
        ctx, 
        migrate.WithDropIndex(true),
        migrate.WithDropColumn(true), 
    )
    if err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }

    CreateUser(ctx, client)

    // GetUser(ctx, client)
    
    // UpdateUser(ctx, client)
    
    // DeleteUser(ctx, client)
}

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error){
    u, err := client.User.
    Create().
    SetName("ahmed").
    SetEmail("ahmed@gmail.com").
    SetPassword("password").
    SetAge(27).
    Save(ctx)

    if err != nil {
        return nil, fmt.Errorf("canot create user : %w", err)
    }

    log.Panicln("user Created : ", u)

    return u, nil
}

func UpdateUser(ctx context.Context, client *ent.Client) (*ent.User, error){
    u, err := client.User.
    UpdateOneID(1).
    SetName("ateeg").
    SetAge(29).
    Save(ctx)

    if err != nil {
        return nil, fmt.Errorf("canot update user : %w", err)
    }

    log.Panicln("user uodated : ", u)

    return u, nil
}

func GetUser(ctx context.Context, client *ent.Client) (*ent.User, error){
    u, err := client.User.
    Query().
    Where(user.NameEQ("Ateeg")).
    Only(ctx)

    if err != nil {
        return nil, fmt.Errorf("canot query user : %w", err)
    }

    log.Panicln("user query : ", u)

    return u, nil
}

func DeleteUser(ctx context.Context, client *ent.Client) (string, error){
    err := client.User.
    DeleteOneID(2).
    Exec(ctx)

    if err != nil {
        return "error", fmt.Errorf("canot query user : %w", err)
    }

    return "done", nil
}