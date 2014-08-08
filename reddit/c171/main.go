package c171

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
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

func hex2byte(src []byte) []byte {
	result := make([]byte, len(src)/2)
	_, err := hex.Decode(result, src)
	if nil != err {
		log.Fatal(err)
	}

	return result
}

func byte2bin(src byte) []byte {
	result := make([]byte, 4)
	param := src
	for i := 0; i < 4; i++ {
		result[3-i] = param & 1
		param = param >> 1
	}

	return result
}

func main() {
	file, _ := ioutil.ReadFile("input")
	for _, v := range file {
		fmt.Print(translate[v])
	}
}
