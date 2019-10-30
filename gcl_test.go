package gcl

import "os"

func ExampleParser() {
	p := NewParser()
	f, err := os.Open("foo.gcl")
	if err != nil {
		panic(err)
	}

	var data = struct {
		Bar string `gcl:"bar"`
	}{}

	if err := p.Add(f.Name(), f); err != nil {
		panic(err)
	}

	if err := p.Parse(&data); err != nil {
		panic(err)
	}
}
