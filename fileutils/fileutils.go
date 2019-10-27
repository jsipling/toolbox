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

func Readfile(filepath string) []byte {
	dat, err := ioutil.ReadFile(filepath)
	check(err)
	return dat
}

func ReadfileStream(filepath string) {
	const BufferSize = 10000
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	buffer := make([]byte, BufferSize)

	for {
		bytesread, err := file.Read(buffer)

		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}

			break
		}

		fmt.Println("bytes read: ", bytesread)
	}
}
