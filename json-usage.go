package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string  `json:"first_name"` //`metadata for json`
	Age       int     `json:"age,omitempty"`
	Email     string  `json:"email,omitempty"`   // if string is empty then omit from json...
	Address   Address `json:"address,omitempty"` //Nested
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

func main() {
	//person := Person{FirstName: "John", Age: 30, Email: "somebody@somewhere.com"}
	//person := Person{FirstName: "John", Age: 30}
	person := Person{FirstName: "John"}

	// convert to JSON:
	jsonStr, err := json.Marshal(person)
	if err != nil {
		log.Println("Failed to marshal object.", err)
	}
	log.Println("JSON representation of the object:", string(jsonStr))

	person2 := Person{FirstName: "Greg", Address: Address{City: "New York", State: "NY"}}
	jsonStr2, _ := json.Marshal(person2)
	log.Println("JSON representation of the object:", string(jsonStr2))

	// converting from JSON:
	// tic (`) below, means raw string. Note: "zip" will be ignored.
	jsonData := `{"first_name": "Mark", "age": 100, "address": {"city": "Middletown", "state":"NJ", "zip":12345}}`
	var inPerson Person
	if err := json.Unmarshal([]byte(jsonData), &inPerson); err != nil {
		log.Fatalf("Failed to convert provded json to the instnce of Person.", err)
	}

	log.Println(inPerson)

	//Handling lists:
	listOfAddrs := []Address{
		{City: "Buffalo", State: "NY"},
		{City: "New York", State: "NY"},
		{City: "Yonkers", State: "NY"},
		{City: "Middletown", State: "NJ"},
		{City: "Miami", State: "FL"},
	}
	fmt.Println(listOfAddrs)
	jsonList, _ := json.Marshal(listOfAddrs)
	fmt.Println(string(jsonList))

	// Handling unknown JSON objecs:
	var data []map[string]interface{} // list of maps of pair(key -> any data)
	err = json.Unmarshal(jsonList, &data)
	if err != nil {
		log.Fatalln("Failed to unmarshal.", err)
	}

	fmt.Println(data)
	for i, e := range data {
		log.Printf("item %d: %v\n", i, e)
	}
}
