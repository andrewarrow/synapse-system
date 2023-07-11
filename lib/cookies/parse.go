package cookies

import (
	"fmt"
	"io/ioutil"
)

func Parse(filename string) {
	b, _ := ioutil.ReadFile(filename)
	s := string(b)
	fmt.Println(len(s))
}
