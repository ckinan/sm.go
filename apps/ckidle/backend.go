package main

import "fmt"

type Backend interface {
	execute(action Action)
}

type SyntheticBackend struct {
}

func NewSyntheticBackend() *SyntheticBackend {
	return &SyntheticBackend{}
}

// execute runs the action determined by the state machine
func (be *SyntheticBackend) execute(action Action) {
	fmt.Printf("executing action=%s\n", action)
}
