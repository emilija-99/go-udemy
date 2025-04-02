package main

import (
	"errors"
	"fmt"
	"log"
)

type TruckNormal struct {
	id    string
	cargo int
}

type TruckElectric struct {
	id      string
	cargo   int
	battery float64
}

type TruckInterface interface {
	LoadCargo() error
	UnloadCargo() error
}

var (
	ErrNotImplemented = errors.New("NOT IMPLEMENTED")
	ErrTruckNotFound  = errors.New("TRUCK NOT FOUND")
)

func (t *TruckNormal) LoadCargo() error {
	t.cargo += 1
	return nil
}

func (t *TruckNormal) UnloadCargo() error {
	t.cargo = 0
	return nil
}

func (t *TruckElectric) LoadCargo() error {
	t.cargo += 1
	t.battery -= 1
	return nil
}

func (t *TruckElectric) UnloadCargo() error {
	t.cargo = 0
	t.battery += 1
	return nil
}

// loading and unloading of a truck
// error always comes last (string, float64, error)
func processTruck(truck TruckInterface) error {
	fmt.Printf("Processing truck: %+v\n", truck)

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("error loading cargo: %w", err)
	}

	if err := truck.UnloadCargo(); err != nil {
		return fmt.Errorf("error unloading cargo: %w", err)
	}

	return nil // no error

}

func main() {
	nt := TruckNormal{id: "1"}
	et := TruckElectric{id: "2", cargo: 100, battery: 100}

	// empty interface
	// person := make(map[string]any, 0) same as interface{] introduced in Go 1.18}
	person := make(map[string]interface{}, 0)
	person["name"] = "Ema"
	person["age"] = 25

	age, exists := person["age"].(int) // type assertion
	if !exists {
		log.Fatalf("Error: %s", "Age not found")
		return
	}

	log.Println("Age: ", age)

	err := processTruck(&nt)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	err = processTruck(&et)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	log.Println("Cargo Normal: ", nt.cargo)
	log.Println("Cargo Electric: ", et.battery)
}
