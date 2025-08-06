package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Person struct {
	XMLName   xml.Name `xml:"person"`     // after unmarshalling will have name of the type ("person" [!])
	FirstName string   `xml:"first_name"` //`metadata for json`
	Age       int      `xml:"age,omitempty"`
	Email     string   `xml:"email,omitempty"`   // if string is empty then omit from json...
	Address   Address  `xml:"address,omitempty"` //Nested
}

type Address struct {
	//XMLName xml.Name `xml:"address"` //embedded fields does not have a root and name is defined in the parrent!
	City  string `xml:"city"`
	State string `xml:"state"`
}

func Check(err error) {
	if err != nil {
		log.Panicln("Error detected:", err)
	}
}

func main() {
	//person := Person{FirstName: "John", Age: 30, Email: "somebody@somewhere.com"}
	person := Person{FirstName: "John", Age: 30}
	//person := Person{FirstName: "John"}

	log.Println("Object:", person)
	// convert to JSON:
	xmlStr, err := xml.Marshal(person)
	Check(err)
	log.Printf("XML representation of the object:\n%s\n", string(xmlStr))

	xmlStr1, err := xml.MarshalIndent(person, " ", "  ")
	Check(err)
	log.Printf("Indented XML representation of the object:\n%s\n", string(xmlStr1))

	person2 := Person{FirstName: "Greg", Address: Address{City: "New York", State: "NY"}}
	xmlStr2, _ := xml.MarshalIndent(person2, " ", "  ")
	log.Printf("XML representation of the object:\n%s\n", string(xmlStr2))

	// converting from XML:
	// tic (`) below, means raw string. Note: "zip" will be ignored.
	//jsonData := `{"first_name": "Mark", "age": 100, "address": {"city": "Middletown", "state":"NJ", "zip":12345}}`

	fmt.Println("Try to unmarshal XML (No indentation!):\n", string(xmlStr2))
	var inPerson Person
	Check(xml.Unmarshal(xmlStr2, &inPerson))

	log.Println("Object from XML:\n", inPerson)

	//Handling lists:
	listOfAddrs := []Address{
		{City: "Buffalo", State: "NY"},
		{City: "New York", State: "NY"},
		{City: "Yonkers", State: "NY"},
		{City: "Middletown", State: "NJ"},
		{City: "Miami", State: "FL"},
	}
	fmt.Println(listOfAddrs)
	xmlList, _ := xml.Marshal(listOfAddrs)
	fmt.Println(string(xmlList))

	// // Handling unknown JSON objecs:
	// var data []map[string]interface{} // list of maps of pair(key -> any data)
	// err = xml.Unmarshal(xmlList, &data)
	// if err != nil {
	// 	log.Fatalln("Failed to unmarshal.", err)
	// }

	// fmt.Println(data)
	// for i, e := range data {
	// 	log.Printf("item %d: %v\n", i, e)
	// }

	book := Book{
		ISBN:  "11-22-33-44-55",
		Title: "Surprise",
	}
	bookXml, err := xml.MarshalIndent(book, "", " ")
	Check(err)

	fmt.Println("Object:", book)
	fmt.Println("XML:\n", string(bookXml))
}

// Attributes: <book isbn="12rtrtb54"</book>
type Book struct {
	XMLName xml.Name `xml:"book"`
	ISBN    string   `xml:"isbn,attr"` //it became attribute
	Title   string   `xml:"title"`
}
