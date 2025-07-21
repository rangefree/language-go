package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(strings.NewReader("Hello, bufio package. Lets try it!"))
	data := make([]byte, 20)

	for {
		count, err := reader.Read(data)

		if err != nil {
			if count == 0 && err.Error() == "EOF" {
				fmt.Println("Done reading.")
				break
			}
			panic(err)
		}

		fmt.Printf("Read %d byte: [%s]\n", count, data[:count])
	}
	// Note: reader cannot use without resetting. It will return error = "EOF" after loop

	// n, err := reader.Read(data)
	// if err != nil { 	panic(err) }
	// fmt.Printf("Read %d byte: [%s]\n", n, data[:n])

	// n, err = reader.Read(data)
	// if err != nil { 	panic(err) }
	// fmt.Printf("Read %d byte: [%s]\n", n, data[:n])

	fmt.Println()

	writer := bufio.NewWriter(os.Stdout)
	//dataOut := make([]byte, 20)
	//i := 0
	for v := byte(0); v < 80; v++ {
		writer.WriteByte(v)
		writer.WriteByte(' ')
	}

	// FLUSH the buffer!!! It will write content to the target
	if err := writer.Flush(); err != nil {
		fmt.Println("Failed to flush:", err)
	}
	fmt.Println()

	if n, err := writer.WriteString("Some string to write through writer...\n"); err != nil {
		panic(err)
	} else {
		writer.Flush()
		fmt.Println("Writer wrote", n, "bytes to the target.")
	}

	// if i == 0 {
	// 	writer.WriteByte(i)
	// }
	// 	dataOut[i] := rand.Intn(256)
	// 	i++
	// }
	// }

}
