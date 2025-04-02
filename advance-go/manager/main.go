package main

import (
	"errors"
	"sync"
)

var ErrTruckNotFound = errors.New("truck not found")

type FleetManager interface {
	AddTruck(id string, cargo int) error
	GetTruck(id string) (Truck, error)
	RemoveTruck(id string) error
	UpdateTruckCargo(id string, cargo int) error
}

type Truck struct {
	ID    string
	Cargo int
}

type truckManager struct {
	trucks       map[string]*Truck
	sync.RWMutex // composing our manager with mutex
}

func newTruckManager() truckManager {
	return truckManager{
		trucks: make(map[string]*Truck),
	}
}

func (m *truckManager) AddTruck(id string, cargo int) error {
	m.Lock()         // lock map, only one goroutine will access to map and lock it
	defer m.Unlock() // unlock for other goroutines be able to access datastructure maps
	m.trucks[id] = &Truck{ID: id, Cargo: cargo}

	return nil
}

func (m *truckManager) GetTruck(id string) (Truck, error) {
	// lock for reading, goroutines have in same time reading but not writting
	m.RLock()
	defer m.RUnlock()

	// pointer is reference to memory
	truck, ok := m.trucks[id]

	if !ok {
		return Truck{}, ErrTruckNotFound
	}
	// derefrence
	return *truck, nil
}

func (m *truckManager) RemoveTruck(id string) error {
	m.Lock()
	defer m.Unlock()

	delete(m.trucks, id)

	return nil
}

func (m *truckManager) UpdateTruckCargo(id string, cargo int) error {
	m.Lock()
	defer m.Unlock()

	m.trucks[id].Cargo = cargo
	return nil
}

func main() {}
