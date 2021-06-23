package main

import (
	"context"
	"fmt"
	"gremlin-orm-sample/ent"
	"gremlin-orm-sample/ent/car"
	"gremlin-orm-sample/ent/user"
	"log"
	"time"

	"entgo.io/ent/dialect"
	"github.com/google/uuid"
)

func main() {
	var err error
	client, err := ent.Open(dialect.Gremlin, "http://localhost:8182")
	if err != nil {
		log.Fatalf("creating client: %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	user, err := CreateUser(ctx, client, uuid.New().String(), 32, "seiichi")
	if err != nil {
		log.Fatalf(err.Error())
	} else {
		log.Println("created user: ", user)
	}

	users, err := SearchUsersById(ctx, client, user.ID)
	if err != nil {
		log.Fatalf(err.Error())
	} else {
		log.Println("searched users by id: ", users)
	}

	car, err := CreateCar(ctx, client, uuid.New().String(), "Tesla")
	if err != nil {
		log.Fatalf(err.Error())
	} else {
		log.Println("created car: ", car)
	}

	newUser, err := AddCarToUser(ctx, client, user, car)
	if err != nil {
		log.Fatalf(err.Error())
	} else {
		log.Println("add car to user: ", newUser)
	}

	newUsers, err := SearchUsersByCarModel(ctx, client, "Tesla")
	if err != nil {
		log.Fatalf(err.Error())
	} else {
		log.Println("searched user by car model: ", newUsers)
	}
}

func CreateUser(ctx context.Context, client *ent.Client, id string, age int, name string) (*ent.User, error) {
	user, err := client.User.
		Create().
		SetID(id).
		SetAge(age).
		SetName(name).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	return user, nil
}

func SearchUsersById(ctx context.Context, client *ent.Client, id string) ([]*ent.User, error) {
	users, err := client.User.
		Query().
		Where(user.IDEQ(id)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	return users, nil
}

func CreateCar(ctx context.Context, client *ent.Client, id string, name string) (*ent.Car, error) {
	car, err := client.Car.
		Create().
		SetID(id).
		SetModel(name).
		SetRegisteredAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %w", err)
	}
	return car, nil
}

func AddCarToUser(ctx context.Context, client *ent.Client, user *ent.User, car *ent.Car) (*ent.User, error) {
	user, err := user.Update().
		AddCars(car).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed add car to user: %w", err)
	}
	return user, nil
}

func SearchUsersByCarModel(ctx context.Context, client *ent.Client, carModel string) ([]*ent.User, error) {
	users, err := client.User.
		Query().
		Where(user.HasCarsWith(car.Model(carModel))).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user cars: %w", err)
	}
	return users, nil
}
