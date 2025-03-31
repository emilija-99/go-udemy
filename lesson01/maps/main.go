package main

import (
	"fmt"
	"maps"
	"sync"
	"time"
)

func main() {

	// MAPS are not concurance
	m := make(map[string]int)
	// "string":1 -> map[ket]value
	// maps return boolean value if key exists or not
	// a, exists := m["a"]

	if _, ok := m["key"]; ok {
		fmt.Println("key exists")
	} else {
		fmt.Println("key does not exists")
		m["key"] = 1
		fmt.Println("m['key'] : ", m["key"])
	}

	if _, ok := m["a"]; ok {
		fmt.Println("key exists")
		fmt.Println("m['a'] : ", m["a"])
	} else {
		fmt.Println("key does not exists")
		m["a"] = 1
		fmt.Println("m['a'] : ", m["a"])
	}

	fmt.Println("m: ", m)

	// delete key from map
	delete(m, "key")
	fmt.Println("m after delete key: ", m)

	// clone map
	m2 := maps.Clone(m)

	fmt.Println("after clone m into m2: ", m2)

	// example using Insert function
	// Insert function is particulary useful when you want to merge maps or add
	// multiple entries at once.
	// iter.Seq2 type represents a sequence that can yield key-value pairs

	seq := func(yield func(string, int) bool) {
		yield("key1", 1)
		yield("key2", 2)
		yield("key3", 3)
	}

	maps.Insert(m2, seq)

	fmt.Println("m2 after insert: ", m2)

	m2_keys := maps.Keys(m2)
	fmt.Println("m2 keys: ", m2_keys)

	eq := maps.Equal(m, m2) // check if maps are equal
	fmt.Println("m and m2 are equal: ", eq)

	// Race condition example
	var wg sync.WaitGroup
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(1 * time.Second)
			m[fmt.Sprintf("key-%d", i)] = i
		}(i)
	}
	wg.Wait()

	fmt.Println("m", m)

	// clear map
	clear(m)
	fmt.Println("m after clear: ", m)

}
