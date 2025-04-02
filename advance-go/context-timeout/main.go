package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

type Truck struct {
	ID    int
	cargo int
}

type NormalTruck struct {
	ID    int
	cargo int
}
type ElectricTruck struct {
	ID      int
	cargo   int
	battery float64
}

type contextKey string

var UserIDKey contextKey = "userID"

func processTruck(ctx, t interface{}) error {
	if normalTruck, ok := t.(*NormalTruck); ok {
		fmt.Printf("Processing NormalTruck: %+v\n", normalTruck)
	}
	if electricTruck, ok := t.(*ElectricTruck); ok {
		fmt.Printf("Processing ElectricTruck: %+v\n", electricTruck)
	}
	return nil
}

func processFleet(ctx context.Context, trucks []interface{}) error {

	// time.Sleep(150 * time.Millisecond)
	// fmt.Printf("Processing fleet with context ", ctx)

	// access user id from context
	// userID := ctx.Value(UserIDKey)
	// fmt.Println("User ID: ", userID)

	// context with timeout
	ctx, canceled := context.WithTimeout(ctx, time.Second*2)
	defer canceled() // cleanup the context

	delay := time.Second * 3
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(delay):
		break
	}

	/*
		C:\Users\emari\Documents\udemy_go_backend\go_lang_udemy\lesson01\context-timeout>go run .
		error processing fleet: context deadline exceeded
	*/
	var waitGroup sync.WaitGroup
	// waitGroup.Add(len(trucks));

	for _, t := range trucks {
		waitGroup.Add(1) // increment the wait group
		// processing go routines
		// go processTruck(t)
		go func(t interface{}) {
			if err := processTruck(ctx, t); err != nil {
				log.Println(err)
			}
			waitGroup.Done() // go routine is done
		}(t)

		time.Sleep(50 * time.Millisecond)
		fmt.Println("Fleet processed")
	}

	waitGroup.Wait() // wait for all go routines to finish

	return nil
}

/*
Package context defines the Context type, which carries deadlines,
canceling signals, and other request-scoped values across API boundaries
and between processes.
*/
func main() {

	ctx := context.Background()
	// or
	// ctx := context.TODO();

	ctx = context.WithValue(ctx, UserIDKey, 42)
	/*
		context.Background returns a non-nil, empty Context. It is never canceled,
		has no values, and has no deadline.
	*/

	// conxtex - first argument in function signature

	fleet := []interface{}{
		&NormalTruck{ID: 1, cargo: 0},
		&ElectricTruck{ID: 2, cargo: 0, battery: 100},
		&NormalTruck{ID: 3, cargo: 0},
		&ElectricTruck{ID: 4, cargo: 0, battery: 100},
	}

	if err := processFleet(ctx, fleet); err != nil {
		fmt.Printf("error processing fleet: %v\n", err)
		return
	}

	fmt.Println("Fleet processed successfully")

}
