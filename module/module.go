package module

import "github.com/spf13/cobra"

type Interface interface {
	Name() string
	Permissions() string
	Configure() error
	Start() error
	Stop() error
	Command() *cobra.Command
}

type Loader interface {
	Load() Interface
}
