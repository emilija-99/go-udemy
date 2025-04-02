package main

import (
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

func processTruck(t interface{}) error {
	if normalTruck, ok := t.(*NormalTruck); ok {
		fmt.Printf("Processing NormalTruck: %+v\n", normalTruck)
	}
	if electricTruck, ok := t.(*ElectricTruck); ok {
		fmt.Printf("Processing ElectricTruck: %+v\n", electricTruck)
	}
	return nil
}

func processFleet(trucks []interface{}) error {
	// time.Sleep(150 * time.Millisecond)

	var waitGroup sync.WaitGroup
	// waitGroup.Add(len(trucks));

	for _, t := range trucks {
		waitGroup.Add(1) // increment the wait group
		// processing go routines
		// go processTruck(t)
		go func(t interface{}) {
			if err := processTruck(t); err != nil {
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
Concurrent programming
- go routines
Basic explanation is that is that it's a program where we have multiple
tasks running simultaneously

- sync go routines
*/
func main() {

	fleet := []interface{}{
		&NormalTruck{ID: 1, cargo: 0},
		&ElectricTruck{ID: 2, cargo: 0, battery: 100},
		&NormalTruck{ID: 3, cargo: 0},
		&ElectricTruck{ID: 4, cargo: 0, battery: 100},
		&NormalTruck{ID: 1, cargo: 0},
		&ElectricTruck{ID: 2, cargo: 0, battery: 100},
		&NormalTruck{ID: 3, cargo: 0},
		&ElectricTruck{ID: 4, cargo: 0, battery: 100},
		&NormalTruck{ID: 1, cargo: 0},
		&ElectricTruck{ID: 2, cargo: 0, battery: 100},
		&NormalTruck{ID: 3, cargo: 0},
		&ElectricTruck{ID: 4, cargo: 0, battery: 100},
		&NormalTruck{ID: 1, cargo: 0},
		&ElectricTruck{ID: 2, cargo: 0, battery: 100},
		&NormalTruck{ID: 3, cargo: 0},
		&ElectricTruck{ID: 4, cargo: 0, battery: 100},
		&NormalTruck{ID: 1, cargo: 0},
		&ElectricTruck{ID: 2, cargo: 0, battery: 100},
		&NormalTruck{ID: 3, cargo: 0},
		&ElectricTruck{ID: 4, cargo: 0, battery: 100},
		&NormalTruck{ID: 1, cargo: 0},
		&ElectricTruck{ID: 2, cargo: 0, battery: 100},
		&NormalTruck{ID: 3, cargo: 0},
		&ElectricTruck{ID: 4, cargo: 0, battery: 100},
	}

	if err := processFleet(fleet); err != nil {
		fmt.Printf("error processing fleet: %v\n", err)
		return
	}

	fmt.Println("Fleet processed successfully")

}
