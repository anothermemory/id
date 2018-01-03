// Package id contains interface and set of implementations for generating different identifiers. For production usage
// identifiers must be mostly unique. For testing purposes it's often much easier to use mocked generator so it will
// generate the same id each time.
package id

// Generator represents interface which can be used to generate random ID values. Actual generated values depends
// on used implementation
type Generator interface {
	// Generate returns new generated ID as string
	Generate() string
}
