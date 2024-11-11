package seeders

import (
	"fmt"

	"github.com/fullstackdev42/mp-emailer/database/core"
	"github.com/fullstackdev42/mp-emailer/database/factories"
)

type UserSeeder struct {
	BaseSeeder
}

func NewUserSeeder(db core.Interface) *UserSeeder {
	return &UserSeeder{BaseSeeder{DB: db}}
}

func (s *UserSeeder) Seed() error {
	factory := factories.NewUserFactory(s.DB)
	users := factory.MakeMany(5)

	for _, user := range users {
		if err := s.DB.Create(user); err != nil {
			return fmt.Errorf("failed to seed user: %w", err)
		}
	}
	return nil
}
