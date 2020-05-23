package fileutils

import (
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
func ReadFileStream(filepath string, bufferSize int, f func([]byte)) {
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

		f(buffer[:bytesread])
	}
}
