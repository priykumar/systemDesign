package main

import (
	"fmt"
	"sync"
)

// This is a thread-safe singleton implementation in Go.
// It ensures that only one instance of dbConnection is created,
// even when multiple goroutines try to access it concurrently.

type dbConnection struct {
	url      string
	userName string
	password string
}

var mu = &sync.Mutex{}
var dbClient *dbConnection

func getInstance(url, user, pwd string, wg *sync.WaitGroup) *dbConnection {
	defer wg.Done()

	if dbClient == nil {
		mu.Lock()
		defer mu.Unlock()
		if dbClient == nil {
			dbClient = &dbConnection{
				url:      url,
				userName: user,
				password: pwd,
			}
			fmt.Println("DB client created!!")
		} else {
			fmt.Println("DB client already exist1")
		}
	} else {
		fmt.Println("DB Client already exist")
	}
	return dbClient
}

func main() {
	var wg sync.WaitGroup
	for range 10 {
		wg.Add(1)
		go getInstance("url", "user1", "user123", &wg)
	}

	wg.Wait()
}
