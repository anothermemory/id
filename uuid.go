package id

import "github.com/satori/go.uuid"

type generatorUUID struct {
}

// Generate returns new generated ID (UUID V4) as string.
func (generatorUUID) Generate() string {
	return uuid.NewV4().String()
}

// NewUUID returns Generator which will return new UUID V4 as generated values.
func NewUUID() Generator {
	return &generatorUUID{}
}
