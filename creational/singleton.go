package creational

import "sync"

type Singleton interface {
	DoWork()
}

type singleton struct {
}

func (s *singleton) DoWork() {
}

var (
	instance *singleton
	once     sync.Once
)

func GetSingletonInstance() Singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
