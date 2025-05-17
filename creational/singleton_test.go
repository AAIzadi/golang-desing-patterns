package creational

import (
	"sync"
	"testing"
)

func TestGetInstance(t *testing.T) {
	instance1 := GetSingletonInstance()
	instance2 := GetSingletonInstance()

	if instance1 != instance2 {
		t.Error("GetSingletonInstance returned different instances, expected the same instance")
	}
}

func TestConcurrentAccess(t *testing.T) {
	const numRoutines = 100
	var wg sync.WaitGroup
	wg.Add(numRoutines)

	instances := make([]Singleton, numRoutines)

	for i := 0; i < numRoutines; i++ {
		go func(index int) {
			defer wg.Done()
			instances[index] = GetSingletonInstance()
		}(i)
	}

	wg.Wait()

	// Verify all instances are the same
	firstInstance := instances[0]
	for i, instance := range instances {
		if instance != firstInstance {
			t.Errorf("Instance %d is not the same as the first instance", i)
		}
	}
}

func TestDoWork(t *testing.T) {
	instance := GetSingletonInstance()

	t.Run("DoesNotPanic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("DoWork panicked: %v", r)
			}
		}()
		instance.DoWork()
	})

}

func TestSingletonInterface(t *testing.T) {
	instance := GetSingletonInstance()
	if _, ok := instance.(Singleton); !ok {
		t.Error("GetSingletonInstance returned value that doesn't implement Singleton interface")
	}
}

func BenchmarkGetInstance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GetSingletonInstance()
	}
}
