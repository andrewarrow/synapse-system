package cookies

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Parse(filename string) {
	b, _ := ioutil.ReadFile(filename)
	s := string(b)
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		tokens := strings.Split(line, "\t")
		if len(tokens) == 1 {
			continue
		}
		name := tokens[0]
		value := tokens[1]
		escapedName := strings.Replace(name, `"`, `\"`, -1)
		escapedValue := strings.Replace(value, `"`, `\"`, -1)

		format := `c = Cookie{"%s","%s"}`
		fmt.Println(fmt.Sprintf(format, escapedName, escapedValue))
		fmt.Println("cookies = append(cookies, c)")
	}
}
