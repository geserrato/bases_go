package main

import "strconv"

func valoresMultiples(name, lastName string, age int) (map[string]string, map[string]string) {
	map1 := make(map[string]string)
	map2 := make(map[string]string)

	map1["name"] = name
	map1["lastName"] = lastName
	map1["age"] = strconv.Itoa(age)

	map2["name"] = name
	map2["lastName"] = lastName
	map2["age"] = strconv.Itoa(age)

	return map1, map2
}

func main() {
	name, lastName, age := "John", "Doe", 30
	map1, map2 := valoresMultiples(name, lastName, age)
	println("Map 1:", map1["name"], map1["lastName"], map1["age"])
	println("Map 2:", map2["name"], map2["lastName"], map2["age"])
}
