package misc

import (
	"fmt"
	"strings"
)

// PrintStruct pretty prints a struct
func PrintStruct(s interface{}) {
	v := fmt.Sprintf("%+v", s)
	for _, s := range strings.Split(v, " ") {
		fmt.Println(s)
	}
}
