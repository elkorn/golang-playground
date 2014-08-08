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

func TestByte2Bin(t *testing.T) {
	tests := []struct {
		input    byte
		expected []byte
	}{
		{1, []byte{0, 0, 0, 1}},
		{2, []byte{0, 0, 1, 0}},
		{3, []byte{0, 0, 1, 1}},
		{4, []byte{0, 1, 0, 0}},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, byte2bin(test.input))
	}
}

func TestShift(t *testing.T) {
	// How could I forget how byte shifting works. :)
	x := byte(13)
	assert.Equal(t, 1, x&1)
	x = x >> 1
	assert.Equal(t, 0, x&1)
	x = x >> 1
	assert.Equal(t, 1, x&1)
	x = x >> 1
	assert.Equal(t, 1, x&1)

	assert.Equal(t, 0, (15>>4)&1)
}

func TestByte2X(t *testing.T) {
	tests := []struct {
		input    byte
		expected []byte
	}{
		{1, []byte("   x")},
		{2, []byte("  x ")},
		{3, []byte("  xx")},
		{4, []byte(" x  ")},
		{15, []byte("xxxx")},
		{255, []byte("xxxxxxxx")},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, byte2x(test.input))
	}
}

func TestBytes2X(t *testing.T) {
	tests := []byteTest{
		{[]byte("FF"), []byte("xxxxxxxx")},
		{[]byte("FF 81"), []byte("xxxxxxxx\nx      x")},
	}

	for _, test := range tests {
		assert.Equal(t, string(test.expected), string(bytes2x(test.input)))
	}
}

func TestMain(t *testing.T) {
	main()
}
