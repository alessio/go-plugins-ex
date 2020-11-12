package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/alessio/go-plugins-ex/module"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("auth: ")
	log.SetOutput(os.Stderr)
}

var (
	state  mod
	Loader loader           //nolint: unused
	_      module.Interface = mod{}
)

type loader string

func (loader) Load() module.Interface { return state }

type mod struct{}

func (mod) Configure() error {
	log.Println("mod is now configured. Ready to start.")

	return nil
}

func (mod) Start() error {
	log.Println("Start() called")

	return nil
}

func (mod) Stop() error {
	log.Println("Stop() called")

	return nil
}

func (mod) Name() string        { return "auth" }
func (mod) Permissions() string { return "io:r" }
func (m mod) Command() *cobra.Command {
	return &cobra.Command{
		Use:   m.Name(),
		Short: `Demo auth command`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SetOut(os.Stdout)
			cmd.Println("args: ", args)

			return nil
		},
	}
}
