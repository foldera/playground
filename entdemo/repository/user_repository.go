package repository

import (
	"context"
	"fmt"
	"github.com/foldera/playground/entdemo/ent"
)

type UserRepository interface {
	Create(ctx context.Context, age int, name string) (*ent.User, error)
	AddCars(ctx context.Context, userId int, cars []*ent.Car) error
}

func NewUserRepository(c *ent.Client) UserRepository {
	return &userRepository{
		client: c,
	}
}

type userRepository struct {
	client *ent.Client
}

func (u *userRepository) Create(ctx context.Context, age int, name string) (*ent.User, error) {
	user, err := u.client.User.Create().SetAge(age).SetName(name).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	return user, nil
}

func (u *userRepository) AddCars(ctx context.Context, userId int, cars []*ent.Car) error {
	model, err := u.client.User.Get(ctx, userId)
	if err != nil {
		return fmt.Errorf("user with id=%d not found: %w", userId, err)
	}
	return model.Update().AddCars(cars...).Exec(ctx)
}
