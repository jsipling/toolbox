package fileutils

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// ReadFile reads the whole file into memory
func ReadFile(filepath string) []byte {
	dat, err := ioutil.ReadFile(filepath)
	check(err)
	return dat
}

// ReadFileStream reads the file as a stream, keeping memory use low
func ReadFileStream(filepath string, bufferSize int, fn func([]byte)) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	buffer := make([]byte, bufferSize)

	for {
		bytesread, err := file.Read(buffer)

		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}

			break
		}

		fn(buffer[:bytesread])
	}
}

// ReadFileWithReadString will read line by line without the 65536 character limit
func ReadFileWithReadString(filepath string, fn func(string)) (err error) {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)
	var line string
	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		// Process the line here.
		fn(line)

		if err != nil {
			break
		}
	}
	if err != io.EOF {
		fmt.Printf(" > Failed with error: %v\n", err)
		return err
	}
	return
}
