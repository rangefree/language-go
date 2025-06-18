package main

import (
	"fmt"
	"maps"
)

func main() {
	// var mapVar map[keyType]valueType
	var myMap map[string]int     // declare nil (!) map - unusable without next line!
	myMap = make(map[string]int) //WARNING!!! allocate storage for map !
	fmt.Println(myMap)

	myMap["Key1"] = 1
	myMap["key2"] = 2

	fmt.Println(myMap)
	fmt.Println("len(myMap) =", len(myMap))

	fmt.Println("myMap[\"Key100\"] =", myMap["Key100"])
	fmt.Println(myMap)

	//check if value exists:
	key := "Key100"
	_, ok := myMap[key]
	fmt.Printf("Value for key (%s) exists? ", key)
	fmt.Println(ok)

	key = "Key1"
	_, ok = myMap[key]
	fmt.Printf("Value for key (%s) exists? ", key)
	fmt.Println(ok)

	myMap2 := map[string]int{"a": 1, "b": 2}
	myMap3 := map[string]int{"a": 1, "b": 2}

	fmt.Println("are maps equal?", maps.Equal(myMap2, myMap3))
	fmt.Println("are maps equal?", maps.Equal(myMap2, myMap2))
	fmt.Println("are maps equal?", maps.Equal(myMap, myMap2))

	for k, v := range myMap {
		fmt.Printf("MAP[%s] = %d \n", k, v)
	}
}
