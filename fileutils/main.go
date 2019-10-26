package fileutils

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Readfile(filename string) []byte {
	dat, err := ioutil.ReadFile(filename)
	Check(err)
	return dat
}

func ReadfileStream(filename string) {
	const BufferSize = 10000
	file, err := os.Open(filename)
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
