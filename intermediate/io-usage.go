package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFromReader(r io.Reader) {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	Check(err)
	fmt.Println(string(buf[:n]))
}

func writeToWriterw(w io.Writer, data string) {
	_, err := w.Write([]byte(data))
	Check(err)
}

func closeTarget(c io.Closer) {
	Check(c.Close())
}

func bufExample() {
	var buf bytes.Buffer // var on "stack"
	buf.WriteString("Some shitty text!")
	fmt.Println(buf.String())

}

func multiReaderExample() {
	r1 := strings.NewReader("Silly text here")
	r2 := strings.NewReader("Sample if you can call it that")
	mr := io.MultiReader(r1, r2)
	pBuf := new(bytes.Buffer) // allocated on "heap"
	n, err := pBuf.ReadFrom(mr)
	Check(err)
	fmt.Println("size =", n, "content:", pBuf.String())

}

func pipeExample() {
	r, w := io.Pipe()

	// GO routine (function which is executed immediately in the separate thread):
	go func() {
		w.Write([]byte("textToPipe"))
		w.Close()
	}()

	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	fmt.Println(buf.String())

}

func writeToFile(filePath string, data string) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	Check(err)
	defer closeTarget(file) //file.Close()

	writer := io.Writer(file)
	n, err := writer.Write([]byte(data))
	Check(err)
	fmt.Println("Wrote", n, "bytes.", data)
}

func main() {
	fmt.Println("Reading from reader:")
	readFromReader(strings.NewReader("Another silly example from the go course."))

	fmt.Println("writing to the writer:")
	var writer bytes.Buffer
	writeToWriterw(&writer, "Hello Writer")
	fmt.Println(writer.String())

	fmt.Println("Buffer example:")
	bufExample()

	fmt.Println("Multi reader example:")
	multiReaderExample()

	fmt.Println("Pipe example:")
	pipeExample()

	fmt.Println("Writing to file:")
	writeToFile("sample.txt", "Some stuff to write...")

}
