package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	out := make(chan string, 1)

	// go routine
	go func() {
		// defer expression are called at the end and they are used for cleanup
		defer f.Close()
		defer close(out) // close channel

		str := ""
		for {
			data := make([]byte, 8)
			n, err := f.Read(data)
			if err != nil {
				break
			}

			data = data[:n]
			if i := bytes.IndexByte(data, '\n'); i != -1 {
				str += string(data[:i])
				data = data[i+1:]
				out <- str
			}

			str += string(data)

		}

		if len(str) == 0 {
			out <- str
		}
	}()

	return out
}

func main() {
	fmt.Println("I hope I get the job!")

	f, err := os.Open("message.txt")
	if err != nil {
		log.Fatal("error", "error", err)
	}

	lines := getLinesChannel(f)

	for line := range lines {
		fmt.Printf("read: %s\n", line)
	}

}
