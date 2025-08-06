package main

import "fmt"

func main() {
	var age uint8 = 10
	fmt.Printf("age: %d, %4d, %-4d, %x, %X, %#X\n", age, age, age, age, age, age)
	s := fmt.Sprint("Hello", "world", "\n", 12, 34)
	fmt.Println(s)

	s = fmt.Sprintln("Hello", "world", "\n", 12, 34) // adding spaces and \n
	fmt.Println(s)

	s = fmt.Sprintf("%s, %v, %#x, %v, %d, PI type: %T \n", "Hello", "world", "\n", 12, 34, 3.14)
	fmt.Println(s)

	var name string
	fmt.Print("Enter name and age:")
	//fmt.Scan(&name, &age)
	//fmt.Scanln(&name, &age)
	fmt.Scanf("%s %d", &name, &age)
	fmt.Println("Hello", name)
	fmt.Println("your age is", age)

	//Error formatting
	if err := checkAge(&age); err != nil {
		fmt.Println(err)
	}

}

func checkAge(age *uint8) error {
	if *age < 18 {
		return fmt.Errorf("Too young to drive. You (%d) are under 18.", *age)
	}
	return nil
}
