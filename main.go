package main

import "log"

func main() {
	m := map[string]bool{
		"ku": true,
	}
	log.Println(m["d"])

	m22 := map[string]struct{}{
		"ku": {},
	}
	log.Println(m22["d"])

	m3 := map[string]bool{
		"ku": false,
	}
	log.Println(m3["ku"])  // false
	log.Println(m3["kud"]) // false
}
