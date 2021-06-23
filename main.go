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

	id := "dummyId"
	age := 32
	name := "seiichi"
	u, err := CreateUser(ctx, client, id, age, name)
	if err != nil {
		log.Fatalf(err.Error())
	} else {
		log.Println("created user: ", u)
	}

	us, err := SearchUsers(ctx, client, id)
	if err != nil {
		log.Fatalf(err.Error())
	} else {
		log.Println("searched users: ", us)
	}
}

func CreateUser(ctx context.Context, client *ent.Client, id string, age int, name string) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetID(id).
		SetAge(age).
		SetName(name).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	return u, nil
}

func SearchUsers(ctx context.Context, client *ent.Client, id string) ([]*ent.User, error) {
	us, err := client.User.
		Query().
		Where(user.IDEQ(id)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	return us, nil
}
