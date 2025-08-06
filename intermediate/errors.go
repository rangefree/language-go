package main

import (
	"errors"
	"fmt"
)

func do_something(i int) error {
	if i < 0 {
		fmt.Printf("do_something(%d) failed\n", i)

		// & allows usage of the error as receiver pointer in Error() function
		// it will be complaint with error interface when used in generic way
		return &myError{id: 1, msg: "Something went wrong"}
	}

	fmt.Printf("do_something(%d) did something\n", i)
	return nil
}

func read_data() error {
	if err := read_config(); err != nil {
		return fmt.Errorf("read_data(): %w", err)
	}
	fmt.Println("read_data() finished successfully")
	return nil
}

func read_config() error {
	return errors.New("read_config(): config error")
}

func Eval(err error, ret ...interface{}) []interface{} {
	if err != nil {
		fmt.Println("Sample of generic error handling: \n\t", err)
		return nil
	}
	return ret
}

func main() {
	if err := do_something(1); err != nil {
		fmt.Println(err)
	}

	if err := do_something(-1); err != nil {
		fmt.Println(err)
	}

	if err := read_data(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("---------------------------------")
	Eval(do_something(1))
	Eval(do_something(-1))
	Eval(read_data())
}

type myError struct {
	msg string
	id  uint32
}

// follow error interface requirement: Error() function should be provided for myError
func (a *myError) Error() string {
	return fmt.Sprintf("Error [%d] - %s", a.id, a.msg)
}

// func (a *myError) error() error {
// 	return errors.New(a.Error())
// }
