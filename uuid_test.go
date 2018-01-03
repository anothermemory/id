package id_test

import (
	"fmt"
	"testing"

	"github.com/anothermemory/id"
	"github.com/stretchr/testify/assert"
)

func TestNewUUID(t *testing.T) {
	g := id.NewUUID()
	assert.NotEmpty(t, g.Generate())
}

func TestGeneratorUUID_Generate(t *testing.T) {
	g := id.NewUUID()
	assert.NotEqual(t, g.Generate(), g.Generate())
}

func ExampleNewUUID() {
	g := id.NewUUID()
	fmt.Println(g.Generate())
}

