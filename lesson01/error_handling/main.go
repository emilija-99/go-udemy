package main

import (
	"errors"
	"fmt"
	"log"
)

type Truck struct {
	id string
}

var (
	ErrNotImplemented = errors.New("NOT IMPLEMENTED")
	ErrTruckNotFound  = errors.New("TRUCK NOT FOUND")
)

func (t *Truck) LoadCargo() error {
	return ErrTruckNotFound
}

// loading and unloading of a truck
// error always comes last (string, float64, error)
func processTruck(truck Truck) error {
	fmt.Println("Processing truck", truck.id)

	// if truck == (Truck{}) {
	// 	return fmt.Errorf("TRUCK OBJECT IS EMPTY.")
	// }

	// if truck.id == "" {
	// 	return fmt.Errorf("ID NOT FOUND.")
	// }

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo: %w", err)
	}

	return ErrNotImplemented // no error

}

func main() {
	trucks := []Truck{
		{id: "truck1"},
		{id: "truck2"},
		{id: "truck3"},
		{id: ""},
		{},
	}

	for _, truck := range trucks {
		// Println do not print variable, Printf does
		fmt.Printf("Truck %s arrived.\n", truck.id)
		if err := processTruck(truck); err != nil {
			// if errors.Is(err, ErrNotImplemented) {
			// 	fmt.Println("Error: Not implemented")
			// }
			// if errors.Is(err, ErrTruckNotFound) {
			// 	fmt.Println("Error: Truck not found")
			// }
			log.Fatalf("Error processing truck: %v", err)
		}

		// if err != nil {
		// 	fmt.Println("Error: ", err)
		// }
		// processTruck(truck)
	}
}
