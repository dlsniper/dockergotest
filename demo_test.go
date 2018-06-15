package dockergotest

import (
	"testing"
)

func TestMyFunc1(t *testing.T) {
	t.Skip("ignored")
}

func TestMyFunc2(t *testing.T) {
	return
}

func TestMyFunc3(t *testing.T) {
	t.Fail()
}
