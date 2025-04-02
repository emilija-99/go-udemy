package main

import (
	"testing"
)

func TestAddTruck(t *testing.T) {
	manager := newTruckManager()
	manager.AddTruck("truck - 1", 100)

	if len(manager.trucks) != 1 {
		t.Errorf("expected 1 truck, you got it: %d", len(manager.trucks))
	}
}

func TestGetTruck(t *testing.T) {
	manager := newTruckManager()
	manager.AddTruck("getMe", 100)

	truck, error := manager.GetTruck("getMe")
	if error != nil {
		t.Errorf("error getting truck by id")
	}

	if truck.ID != "getMe" {
		t.Errorf("expected truck with ID %s", truck.ID)
	}
}

func TestRemoveTruck(t *testing.T) {
	manager := newTruckManager()

	manager.AddTruck("key", 100)

	manager.RemoveTruck("key")

	_, error := manager.GetTruck("key")

	if error != ErrTruckNotFound {
		t.Errorf("excepted 'truck not found' error, but recived: %v", error)
	}

	if len(manager.trucks) != 0 {
		t.Errorf("excepted 0 trucks but received: %d", len(manager.trucks))
	}
}

func TestUpdateTruckCargo(t *testing.T) {
	manager := newTruckManager()
	manager.AddTruck("ID", 200)

	manager.UpdateTruckCargo("ID", 1)

	truck, err := manager.GetTruck("ID")

	if err != nil {
		t.Errorf("excepted no error, but received %v", err)
	}
	if truck.Cargo != 1 {
		t.Errorf("expected truck with cargo: 1, but received truck with cargo: %d", truck.Cargo)
	}
}

/*
This is a test for concurency.
Maps are not concurency.

How to handle race condition with maps:
Functional and mutable code
or mutex.
*/
func TestConcurrentUpdate(t *testing.T) {
	manager := newTruckManager()

	manager.AddTruck("1", 100)

	const numGoroutines = 100
	const iterations = 100

	done := make(chan bool)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			for j := 0; j < iterations; j++ {
				truck, _ := manager.GetTruck("1")
				manager.UpdateTruckCargo("1", truck.Cargo+1)
			}
			done <- true
		}()
	}
	// Wait for all goroutines to complete
	for i := 0; i < numGoroutines; i++ {
		<-done
	}
}
