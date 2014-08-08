package c171

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type byteTest struct {
	input    []byte
	expected []byte
}

func TestHex2Byte(t *testing.T) {
	tests := []byteTest{
		{[]byte("01"), []byte{1}},
		{[]byte("0F"), []byte{15}},
		{[]byte("10"), []byte{16}},
		{[]byte("FF"), []byte{255}},
		{[]byte("FFFF"), []byte{255, 255}},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, hex2byte(test.input))
	}
}

// func TestMain(t *testing.T) {
// 	main()
// }
