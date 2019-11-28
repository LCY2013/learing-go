package client

import (
	"module1/ch15/services"
	"testing"
)

func TestServices(t *testing.T) {
	t.Log(services.Fibonacci(5))
}
