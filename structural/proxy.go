package structural

import (
	"fmt"
	"sync"
)

type User struct {
	ID   string
	Name string
}

type UserFinder interface {
	Find(userId string) (*User, error)
}

type UserMemoryFinder struct {
	users map[string]User
	mu    sync.RWMutex
}

func NewUserMemoryFinder() *UserMemoryFinder {
	return &UserMemoryFinder{
		users: make(map[string]User),
	}
}

func (f *UserMemoryFinder) Add(user User) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.users[user.ID] = user
}

func (f *UserMemoryFinder) Find(userId string) (*User, error) {
	f.mu.RLock()
	defer f.mu.RUnlock()
	if user, ok := f.users[userId]; ok {
		return &user, nil
	}
	return nil, nil
}

type UserDBFinder struct {
}

func NewUserDBFinder() *UserDBFinder {
	return &UserDBFinder{}
}

func (f *UserDBFinder) Find(userID string) (*User, error) {

	users := map[string]User{
		"1": {ID: "1", Name: "John Doe"},
		"2": {ID: "2", Name: "Jane Smith"},
		"3": {ID: "3", Name: "Bob Johnson"},
	}

	if user, ok := users[userID]; ok {
		return &user, nil
	}
	return nil, nil
}

type UserFinderProxy struct {
	dbFinder     UserFinder
	memoryFinder *UserMemoryFinder
}

func NewUserFinderProxy(dbFinder UserFinder, memoryFinder *UserMemoryFinder) *UserFinderProxy {
	return &UserFinderProxy{
		dbFinder:     dbFinder,
		memoryFinder: memoryFinder,
	}
}

func (p *UserFinderProxy) Find(userId string) (*User, error) {
	if user, err := p.memoryFinder.Find(userId); err != nil {
		return nil, fmt.Errorf("memory lookup error: %w", err)
	} else if user != nil {
		fmt.Printf("User %s found in cache\n", userId)
		return user, nil
	}

	user, err := p.dbFinder.Find(userId)
	if err != nil {
		return nil, fmt.Errorf("db lookup error: %w", err)
	}

	// Cache the result if found
	if user != nil {
		fmt.Printf("User %s fetched from DB and cached\n", userId)
		p.memoryFinder.Add(*user)
		return user, nil
	}

	return nil, fmt.Errorf("user %s not found\n", userId)

}
