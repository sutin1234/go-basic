package oscar

// _ ไม่ต้องสนใจค่านี้
// := shorthand assign value

import (
	"encoding/csv"
	"log"
	"os"
)

func ActorWhoGotMoreThanOne(filename string) map[string]int {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)
	data, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	cache := map[string]int{}
	for _, record := range data[1:] {
		cache[record[3]]++ // append key to map
	}

	for name, count := range cache {
		if count <= 1 {
			delete(cache, name) // delete key from map
		}
	}

	return cache

}
