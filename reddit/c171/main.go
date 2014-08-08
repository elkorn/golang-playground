package c171

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

const SPACE = byte(' ')
const ENTER = byte('\n')
const HALF_BYTE = 4

func hex2byte(src []byte) []byte {
	result := make([]byte, len(src)/2)
	_, err := hex.Decode(result, src)
	if nil != err {
		log.Fatal(err)
	}

	return result
}

func byte2out(src, zero, one byte) []byte {
	size := HALF_BYTE
	if src > 15 {
		size = size + HALF_BYTE
	}

	result := make([]byte, size)
	mapping := []byte{zero, one}
	param := src
	for i := 0; i < size; i++ {
		result[size-i-1] = mapping[param&1]
		param = param >> 1
	}

	return result
}

func byte2bin(src byte) []byte {
	return byte2out(src, 0, 1)
}

func byte2x(src byte) []byte {
	return byte2out(src, byte(' '), byte('x'))
}

func bytes2x(src []byte) []byte {
	var result bytes.Buffer
	n, start := len(src), 0
	writeSlice := func(slice []byte) {
		for _, b := range hex2byte(slice) {
			result.Write(byte2x(b))
		}
	}

	for i := 0; i < n; i++ {
		if SPACE == src[i] {
			writeSlice(src[start:i])
			result.WriteByte(ENTER)
			start = i + 1
		}
	}

	writeSlice(src[start:n])

	return result.Bytes()
}

func main() {
	file, _ := os.Open("input")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(string(bytes2x(scanner.Bytes())))
	}
}
