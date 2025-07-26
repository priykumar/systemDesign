package singleton

import (
	"sync"
	"testing"
)

func TestSingletonInstance(t *testing.T) {
	const goroutines = 100
	var wg sync.WaitGroup
	instances := make(chan *dbConnection, goroutines)

	for range goroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			inst := GetInstance("url", "admin", "pass")
			instances <- inst
		}()
	}

	wg.Wait()
	close(instances)

	var first *dbConnection
	for inst := range instances {
		if first == nil {
			first = inst
			continue
		}
		if first != inst {
			t.Errorf("Expected same instance, got different ones")
		}
	}
}
