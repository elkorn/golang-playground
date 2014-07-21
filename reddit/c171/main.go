package c171

import (
	"fmt"
	"io/ioutil"
)

var translate map[byte]string = map[byte]string{
	'0':  "    ",
	'1':  "   x",
	'2':  "  x ",
	'3':  "  xx",
	'4':  " x  ",
	'5':  " x x",
	'6':  " xx ",
	'7':  " xxx",
	'8':  "x   ",
	'9':  "x  x",
	'A':  "x x ",
	'B':  "x xx",
	'C':  "xx  ",
	'D':  "xx x",
	'E':  "xxx ",
	'F':  "xxxx",
	' ':  "\n",
	'\n': "\n\n",
}

func main() {
	file, _ := ioutil.ReadFile("input")
	for _, v := range file {
		fmt.Print(translate[v])
	}
}
