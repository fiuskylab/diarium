package main

import (
	"fmt"

	"github.com/fiuskylab/diarium"
)

func main() {
	client := diarium.NewDefaultClient()

	err := fmt.Errorf("Foo bar")

	if err != nil {
		client.Log(diarium.Error, err.Error(), nil)
		// handles error
	}
}
