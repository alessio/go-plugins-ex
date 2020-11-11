package main

import (
	"log"
	"os"

	"github.com/alessio/go-plugins-ex/module"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("auth: ")
	log.SetOutput(os.Stderr)
}

var state *mod

func Module() module.Interface {
	if state == nil {
		state = new(mod)
	}

	return state
}

type mod struct{}

func (p *mod) Configure() error {
	log.Println("mod is now configured. Ready to start.")

	return nil
}

func (p *mod) Start() error {
	log.Println("Start() called")

	return nil
}

func (p *mod) Stop() error {
	log.Println("Stop() called")

	return nil
}

func (*mod) Name() string        { return "auth" }
func (*mod) Permissions() string { return "io:r" }
