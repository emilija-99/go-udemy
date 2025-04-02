package main

import (
	"testing"
)

// go test -v *.go - run more files
func TestMain(t *testing.T) {
	t.Run("Process Trunk", func(t *testing.T) {
		t.Run("Should load and unload truck cargo", func(t *testing.T) {
			nt := TruckNormal{id: "1", cargo: 42}
			et := TruckElectric{id: "2", cargo: 100, battery: 100}

			err := processTruck(&nt)
			if err != nil {
				t.Fatalf("Error: %s", err)
			}

			err = processTruck(&et)
			if err != nil {
				t.Fatalf("Error: %s", err)
			}

			// asserting
			if nt.cargo != 0 {
				t.Fatalf("Expected cargo to be 0, got %d", nt.cargo)
			}

			if et.battery != 100 {
				t.Fatalf("Expected battery to be -2, got %f", et.battery)
			}
		})
	})
}
