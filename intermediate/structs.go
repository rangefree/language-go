package main

import (
	"fmt"
)

type Person struct {
	fName         string
	lName         string
	age           uint8
	addr          Address // nested struct
	PhoneHomeCell         // EMBEDDED STRUCT (used for "inheritance") anonymous struct - fields accessed directly
}

type Address struct {
	city    string
	country string
}

type PhoneHomeCell struct {
	home string
	cell string
}

func (a PhoneHomeCell) GetPhones() string {
	return fmt.Sprintf("Home: %s, Cell: %s", a.home, a.cell)
}

func (a PhoneHomeCell) Show() string {
	return a.GetPhones()
}

func (a Person) Show() string {
	ret := fmt.Sprintf("%s:\n\tage:\t%v\n\t%s", a.FullName(), a.age, a.GetPhones())
	return ret
}

// method(s) should be outside the struct
func (a Person) FullName() string { return a.fName + " " + a.lName }

func (pA *Person) IncrAge() { pA.age = pA.age + 1 }

type Rectangle struct {
	length float64
	width  float64
}

func (a Rectangle) Area() float64 {
	return a.width * a.length

}

func (pA *Rectangle) Scale(factor float64) {
	pA.length *= factor
	pA.width *= factor
}

// define my integer type:
type MyInt int

func (a MyInt) IsPositive() bool {
	return !(a > 0)
}

type Shape struct {
	Rectangle
}

func main() {

	p := Person{fName: "john", lName: "Doe", age: 100, addr: Address{city: "London", country: "UK"},
		PhoneHomeCell: PhoneHomeCell{home: "000-1111", cell: "111-2222"},
	}
	fmt.Println(p)

	user := struct {
		name  string
		email string
	}{name: "user123", email: "whatever@whatever.com"}

	fmt.Println(user)
	fmt.Println("p.FullName() =", p.FullName())

	p.IncrAge()
	fmt.Println(p)
	fmt.Println(p.addr)

	p.addr.city = "whatever"
	p.cell = "3333-4444"
	fmt.Println(p.addr)
	fmt.Println(p)
	fmt.Println("Calling method of the embedded structure through owner:", p.GetPhones())
	fmt.Println("Calling overwritten method Show() of the owner:\n", p.Show())

	rect := Rectangle{length: 10, width: 9}
	area := rect.Area()
	fmt.Printf("Area of rectangle %6.2f x %6.2f is %6.2f\n", rect.length, rect.width, area)
	rect.Scale(2)
	fmt.Printf("Area of the rectangle, scaled by factor of 2: %6.2f\n", rect.Area())

	shape := Shape{Rectangle: Rectangle{length: 5, width: 10}}
	fmt.Printf("Calling Area() method on Shape (%3.2f x %3.2f) type: %3.2f",
		shape.length, shape.width, shape.Area())

}
