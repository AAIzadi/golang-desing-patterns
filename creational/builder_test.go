package creational

import (
	"testing"
)

func TestUserBuilder(t *testing.T) {
	t.Run("Empty User", func(t *testing.T) {
		builder := NewUserBuilder()
		user := builder.Build()

		if user.Name != "" {
			t.Errorf("Expected empty name, got %s", user.Name)
		}
		if user.Email != "" {
			t.Errorf("Expected empty email, got %s", user.Email)
		}
		if user.Age != 0 {
			t.Errorf("Expected age 0, got %d", user.Age)
		}
	})

	t.Run("Fully Specified User", func(t *testing.T) {
		name := "John Doe"
		email := "john@example.com"
		age := 30

		user := NewUserBuilder().
			Name(name).
			Email(email).
			Age(age).
			Build()

		if user.Name != name {
			t.Errorf("Expected name %s, got %s", name, user.Name)
		}
		if user.Email != email {
			t.Errorf("Expected email %s, got %s", email, user.Email)
		}
		if user.Age != age {
			t.Errorf("Expected age %d, got %d", age, user.Age)
		}
	})

	t.Run("Partial Specification", func(t *testing.T) {
		name := "Alice"
		age := 25

		user := NewUserBuilder().
			Name(name).
			Age(age).
			Build()

		if user.Name != name {
			t.Errorf("Expected name %s, got %s", name, user.Name)
		}
		if user.Email != "" {
			t.Errorf("Expected empty email, got %s", user.Email)
		}
		if user.Age != age {
			t.Errorf("Expected age %d, got %d", age, user.Age)
		}
	})

	t.Run("Chaining Methods", func(t *testing.T) {
		user := NewUserBuilder().
			Name("Bob").
			Email("bob@builder.com").
			Age(40).
			Name("Robert"). // Override previous name
			Build()

		if user.Name != "Robert" {
			t.Errorf("Expected name Robert, got %s", user.Name)
		}
		if user.Email != "bob@builder.com" {
			t.Errorf("Expected email bob@builder.com, got %s", user.Email)
		}
		if user.Age != 40 {
			t.Errorf("Expected age 40, got %d", user.Age)
		}
	})

	t.Run("Zero Values", func(t *testing.T) {
		user := NewUserBuilder().
			Name("").
			Email("").
			Age(0).
			Build()

		if user.Name != "" {
			t.Errorf("Expected empty name, got %s", user.Name)
		}
		if user.Email != "" {
			t.Errorf("Expected empty email, got %s", user.Email)
		}
		if user.Age != 0 {
			t.Errorf("Expected age 0, got %d", user.Age)
		}
	})
}
