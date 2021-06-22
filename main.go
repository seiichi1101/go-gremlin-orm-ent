package main

import (
	"context"
	"fmt"
	"gremlin-orm-sample/ent"
	"gremlin-orm-sample/ent/user"
	"log"

	"entgo.io/ent/dialect"
)

func main() {
	var err error
	client, err := ent.Open(dialect.Gremlin, "http://localhost:8182")
	if err != nil {
		log.Fatalf("creating client: %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	nickname := "seiichi"
	u, err := CreateUser(ctx, client, nickname)
	if err != nil {
		log.Fatalf(err.Error())
	} else {
		log.Println("user returned: ", u)
	}

	us, err := QueryUser(ctx, client, nickname)
	if err != nil {
		log.Fatalf(err.Error())
	} else {
		log.Println("user returned: ", us)
	}
}

func CreateUser(ctx context.Context, client *ent.Client, nickname string) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetID("dummyId").
		SetName(nickname).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	return u, nil
}

func QueryUser(ctx context.Context, client *ent.Client, nickname string) ([]*ent.User, error) {
	us, err := client.User.
		Query().
		Where(user.NameEQ(nickname)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	return us, nil
}
