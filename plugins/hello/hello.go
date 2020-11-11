package main

import (
	"errors"
	"log"
	"os"

	"github.com/alessio/go-plugins-ex/module"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("hello: ")
	log.SetOutput(os.Stderr)
}

var state *mod

func Module() module.Interface {
	if state == nil {
		state = new(mod)
	}

	return state
}

type mod struct {
	configured bool
}

func (p *mod) Configure() error {
	log.Println("mod is now configured. Read to start.")

	p.configured = true

	return nil
}

func (p *mod) Start() error {
	if !p.configured {
		return ErrNotConfigured
	}

	log.Println("Start() called")

	return nil
}

func (p *mod) Stop() error {
	log.Println("Stop() called")

	return nil
}

func (*mod) Name() string        { return "hello" }
func (*mod) Permissions() string { return "io:rw]" }

var ErrNotConfigured = errors.New("module is not configured")
