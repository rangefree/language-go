package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fileName := "out.txt"
	if err := PrepareFile(fileName); err != nil {
		panic(err)
	}
	fmt.Println()

	_, err := ReadFileAtOnce(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println()

	_, err = ReadLineByLine(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println()

}

func PrepareFile(fileName string) (err error) {
	//Create file:
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Failed to create file:", err)
		return err
	}

	fmt.Printf("File %s is opened.\n", fileName)
	defer func() {
		file.Close() //close when leaving the scope
		fmt.Printf("File %s is closed.\n", fileName)
	}() // it is function call

	// Write bytes:
	text := "This file was created from go language.\n"
	file.WriteString(text)
	file.Write([]byte(text))

	// Write through writer
	data := []byte("Text to write from the buffer: Pretend that we have some pretty long textual data which we have somewhere and want to write to the file.\n")
	fWriter := bufio.NewWriter(file)
	fWriter.Write(data)
	fWriter.Flush()

	// Read data in chunks and writing it to the file:
	reader := bufio.NewReader(strings.NewReader(string(data)))
	buffer := make([]byte, 10) // my sliding buffer
	chunkNumber := 0
	for {
		count, err := reader.Read(buffer)
		fmt.Printf("Writing data chunk # %d: %d bytes\n", chunkNumber, count)
		if err != nil {
			if count == 0 && err.Error() == "EOF" {
				fmt.Println("Last chunk processed.")
				break
			}
			return err
		}
		chunkNumber++

		count, err = file.Write(buffer[0:count])
		if err != nil {
			fmt.Println("Failed to write from buffer: ", err)
			return err
		}
		fmt.Printf("Wrote %d byte from the buffer into the file %s.\n", count, fileName)
	}
	return nil
}

func ReadFileAtOnce(fileName string) (total int, err error) {
	total = 0
	file, err := os.Open(fileName)
	if err != nil {
		return total, err
	}
	fmt.Println("File", fileName, "is opened.")
	defer func() {
		file.Close()
		fmt.Println("File", fileName, "is closed.")
	}()

	buf := make([]byte, 1024)
	total, err = file.Read(buf)
	if err != nil {
		fmt.Println("Failed to read data from file", fileName, ":", err)
		return total, err
	}

	fmt.Printf("Read %d bytes from file\n", total)
	fmt.Println(string(buf))
	return total, err
}

func ReadLineByLine(fileName string) (total int, err error) {
	total = 0
	file, err := os.Open(fileName)
	if err != nil {
		return total, err
	}
	fmt.Println("File", fileName, "is opened.")
	defer func() {
		file.Close()
		fmt.Println("File", fileName, "is closed.")
	}()

	scanner := bufio.NewScanner(file)
	lineNum := 0
	lineLen := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineNum++
		lineLen = len(line)
		fmt.Printf("%4d,%4d: %s\n", lineNum, lineLen, line)
		total += lineLen + 1 /*carriage return (\n) from file*/
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Read %d bytes from file\n", total)
	return total, err
}
