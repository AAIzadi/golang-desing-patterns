package structural

import (
	"errors"
	"sync"
	"testing"
)

func TestUserMemoryFinder(t *testing.T) {
	t.Run("Add and Find user", func(t *testing.T) {
		finder := NewUserMemoryFinder()
		user := User{ID: "1", Name: "Test User"}

		finder.Add(user)
		found, err := finder.Find("1")

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if found == nil {
			t.Fatal("Expected to find user, got nil")
		}
		if found.ID != user.ID || found.Name != user.Name {
			t.Errorf("Expected %v, got %v", user, *found)
		}
	})

	t.Run("Find non-existent user", func(t *testing.T) {
		finder := NewUserMemoryFinder()
		found, err := finder.Find("999")

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if found != nil {
			t.Errorf("Expected nil, got %v", *found)
		}
	})

	t.Run("Concurrent access", func(t *testing.T) {
		finder := NewUserMemoryFinder()
		var wg sync.WaitGroup

		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				user := User{ID: string(rune(id)), Name: "User"}
				finder.Add(user)
				_, _ = finder.Find(string(rune(id)))
			}(i)
		}

		wg.Wait()
	})
}

func TestUserDBFinder(t *testing.T) {
	finder := NewUserDBFinder()

	tests := []struct {
		name     string
		userID   string
		wantUser *User
	}{
		{"Existing user", "1", &User{ID: "1", Name: "John Doe"}},
		{"Another existing user", "2", &User{ID: "2", Name: "Jane Smith"}},
		{"Non-existent user", "999", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := finder.Find(tt.userID)

			if err != nil {
				t.Fatalf("Find() error = %v", err)
			}
			if got != nil && got.ID != tt.wantUser.ID && got.Name != tt.wantUser.Name {
				t.Errorf("Find() = %v, want %v", got, tt.wantUser)
			}
		})
	}
}

func TestUserFinderProxy(t *testing.T) {
	memoryFinder := NewUserMemoryFinder()
	dbFinder := NewUserDBFinder()
	proxy := NewUserFinderProxy(dbFinder, memoryFinder)

	t.Run("First lookup (should hit DB)", func(t *testing.T) {
		user, err := proxy.Find("1")

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if user == nil {
			t.Fatal("Expected to find user, got nil")
		}
		if user.ID != "1" || user.Name != "John Doe" {
			t.Errorf("Expected John Doe, got %v", *user)
		}
	})

	t.Run("Second lookup (should hit cache)", func(t *testing.T) {
		user, err := proxy.Find("1")

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if user == nil {
			t.Fatal("Expected to find user, got nil")
		}
		if user.ID != "1" || user.Name != "John Doe" {
			t.Errorf("Expected John Doe, got %v", *user)
		}
	})

	t.Run("Non-existent user", func(t *testing.T) {
		user, err := proxy.Find("999")

		if err == nil {
			t.Error("Expected error for non-existent user")
		}
		if user != nil {
			t.Errorf("Expected nil, got %v", *user)
		}
		if !(err.Error() == "user 999 not found\n") {
			t.Errorf("Expected 'user 999 not found' error, got %v", err)
		}
	})

	t.Run("Cache is populated from DB", func(t *testing.T) {
		// Verify the cache was populated from the first test
		found, err := memoryFinder.Find("1")

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if found == nil {
			t.Fatal("Expected user to be in cache")
		}
		if found.ID != "1" || found.Name != "John Doe" {
			t.Errorf("Expected cached John Doe, got %v", *found)
		}
	})
}

// MockUserFinder is a mock implementation for testing
type MockUserFinder struct {
	FindFunc func(userID string) (*User, error)
}

func (m *MockUserFinder) Find(userID string) (*User, error) {
	return m.FindFunc(userID)
}

// TestUserFinderProxyWithMocks tests with mock dependencies
func TestUserFinderProxyWithMocks(t *testing.T) {
	t.Run("DB error propagates", func(t *testing.T) {
		mockDB := &MockUserFinder{
			FindFunc: func(userID string) (*User, error) {
				return nil, errors.New("database connection failed")
			},
		}
		mockCache := NewUserMemoryFinder()
		proxy := NewUserFinderProxy(mockDB, mockCache)

		_, err := proxy.Find("1")
		if err == nil {
			t.Fatal("Expected error, got nil")
		}
		if err.Error() != "db lookup error: database connection failed" {
			t.Errorf("Expected DB error, got %v", err)
		}
	})
}

// BenchmarkUserFinderProxy benchmarks the proxy performance
func BenchmarkUserFinderProxy(b *testing.B) {
	memoryFinder := NewUserMemoryFinder()
	dbFinder := NewUserDBFinder()
	proxy := NewUserFinderProxy(dbFinder, memoryFinder)

	// Warm up the cache
	proxy.Find("1")
	proxy.Find("2")
	proxy.Find("3")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Alternate between cached and non-cached lookups
		id := string(rune(i%3 + 1)) // IDs 1, 2, 3
		_, _ = proxy.Find(id)
	}
}
