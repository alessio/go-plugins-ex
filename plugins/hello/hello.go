package main

import (
	"errors"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/alessio/go-plugins-ex/module"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("hello: ")
	log.SetOutput(os.Stderr)
}

var (
	state  *mod
	Loader loader           //nolint: unused
	_      module.Interface = &mod{}
)

type loader string

func (loader) Load() module.Interface {
	if state == nil {
		state = &mod{}
	}

	return state
}

type mod struct {
	configured bool
}

func (m *mod) Configure() error {
	log.Println("mod is now configured. Ready to start.")

	m.configured = true

	return nil
}

func (m *mod) Start() error {
	if !m.configured {
		return ErrNotConfigured
	}

	log.Println("Start() called")

	return nil
}

func (m *mod) Stop() error {
	log.Println("Stop() called")

	return nil
}

func (m *mod) Name() string        { return "hello" }
func (m *mod) Permissions() string { return "io:rw]" }
func (m *mod) Command() *cobra.Command {
	return &cobra.Command{
		Use:   m.Name(),
		Short: `Demo hello command`,
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SetOut(os.Stdout)
			cmd.Println("hello ", args[0])

			return nil
		},
	}
}

var ErrNotConfigured = errors.New("module is not configured")
