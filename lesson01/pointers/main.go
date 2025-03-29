package main

import "fmt"

type Truck struct {
	ID    int
	cargo int
}

func main() {
	truckID := 42
	// 42 address of truckID: 0xc000102088

	anotherTruckID := &truckID // pointer to int -> truckID
	// same as: var anotherTruckID *int = &truckID

	fmt.Println("Truck", truckID, "address of truckID:", &truckID)
	fmt.Println("Another Truck", anotherTruckID, "address of anotherTruckID:", &anotherTruckID)
	/*
		Truck 42 address of truckID: 0xc00000a118
		Another Truck 0xc00000a118 address of anotherTruckID: 0xc0000540d0
	*/
	// dereferencing means to access the value of the pointer
	fmt.Println("Dereferencing anotherTruckID:", *anotherTruckID)

	/*
		Here we are changing the value of truckID but not the address
		AnotherTruck is a pointer to truckID so anotherTruckID will also change
	*/
	fmt.Println("Changing the value of truckID:")
	truckID = 100
	fmt.Println("Truck", truckID, "address of truckID:", &truckID)
	fmt.Println("Another Truck", anotherTruckID, "address of anotherTruckID:", &anotherTruckID, "value of anotherTruckID:", *anotherTruckID)

	/*
		Truck 42 address of truckID: 0xc00000a118
		Another Truck 0xc00000a118 address of anotherTruckID: 0xc0000540d0
		Dereferencing anotherTruckID: 42
		Changing the value of truckID:
		Truck 100 address of truckID: 0xc00000a118
		Another Truck 0xc00000a118 address of anotherTruckID: 0xc0000540d0 value of anotherTruckID: 100
	*/

	fmt.Println("Working with functions and pointers:")

	truck := Truck{ID: 42, cargo: 0}
	fmt.Printf("Truck %+v, address of truck: %p\n", truck, &truck)

	fillTheTruck(truck)

	fmt.Printf("Truck %+v, address of truck: %p\n", truck, &truck)

	/*
		Working with functions and pointers:
		Truck {ID:42 cargo:0}, address of truck: 0xc00000a160
		Truck {ID:42 cargo:0}, address of truck: 0xc00000a160

		Sending the truck as a parameter to the function
		fillTheTruck will create a copy of the truck
		so the original truck will not be modified
	*/

	fillTheTruckPointer(&truck)

	fmt.Printf("Truck %+v, address of truck: %p\n", truck, &truck)

	/*
		Truck {ID:42 cargo:0}, address of truck: 0xc00000a160
		Truck {ID:42 cargo:100}, address of truck: 0xc00000a160

		Function will access the address in the memory of the Truck
		and then change the actual value of the truck
	*/
}

func fillTheTruckPointer(truck *Truck) {
	truck.cargo = 100
}

func fillTheTruck(truck Truck) {
	truck.cargo = 100

	// truck is a copy of the original truck
	// so the truck we are modifying is not the original truck
	// and there having different addresses
	fmt.Printf("Truck %+v, address of truck: %p\n", truck, &truck)

	/*
		Truck {ID:42 cargo:0}, address of truck: 0xc00000a160
		Truck {ID:42 cargo:100}, address of truck: 0xc00000a190
	*/
}
