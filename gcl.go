// Package gcl implements the stuff needed.
// WARNING: THIS API MIGHT NOT BE STABLE. DO NOT USE IT YET.
package gcl // import "github.com/conmon/gcl"

import "io"

type Parser struct {
}

// NewParser returns a new Parser
func NewParser() *Parser {
	return new(Parser)
}

// Add adds the new file to the list of data which will be
func (p Parser) Add(f string, r io.Reader) error {
	return nil
}

// Parse will parse all the data into v, or return an error if one occured.
func (p Parser) Parse(v interface{}) error {
	return nil
}
