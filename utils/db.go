package utils

import (
	"context"
	"ent-demo/ent"
	"ent-demo/ent/migrate"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	MyClient *ent.Client
	MyCtx = context.Background()
)


func ConnectToDB() {
	client, err := ent.Open("mysql", "root:@tcp(localhost:3306)/ent_db?parseTime=True")

	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	
	if err == nil{
		log.Println("db connected")
		MyClient = client
	}
	
	// defer client.Close()
	ctx := context.Background()
	// Run migration.
	err = MyClient.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

}
