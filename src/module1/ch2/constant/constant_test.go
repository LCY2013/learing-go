package constant

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

/*const (
	Monday = 1
	Tuesday = 2
	Wednesday = 3
)*/

const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
	t.Log(Readable, Writeable, Executable)
}

func TestConstantTry1(t *testing.T) {
	const num = 7 //0111
	t.Log(num&Readable == Readable, num&Writeable == Writeable, num&Executable == Executable)
	const number = 1 //0001
	t.Log(number&Readable == Readable, number&Writeable == Writeable, number&Executable == Executable)
}
