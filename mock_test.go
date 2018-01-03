package id_test

import (
	"fmt"
	"testing"

	"github.com/anothermemory/id"
	"github.com/stretchr/testify/assert"
)

func TestNewMock(t *testing.T) {
	g := id.NewMock("123")
	assert.Equal(t, "123", g.Generate())
}

func TestGeneratorMock_Generate_SameValue(t *testing.T) {
	g := id.NewMock("123")
	assert.Equal(t, g.Generate(), g.Generate())
}

func ExampleNewMock() {
	g := id.NewMock("123")
	fmt.Println(g.Generate())
	// Output: 123
}
