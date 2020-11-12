package commands

import (
	"fmt"
	"github.com/alessio/go-plugins-ex/registry"
	"github.com/spf13/cobra"
	"sort"
)

// rootCmd represents the base command when called without any subcommands
var demoCmd = &cobra.Command{
	Use:   "demo",
	Short: "demo plugins",
	Long:  `Configure and execute modules one after another.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := registry.LoadFromDir(pluginsDir); err != nil {
			return err
		}

		modules := registry.Modules()

		sort.Strings(sort.StringSlice(modules))

		for _, name := range modules {
			mod, _ := registry.Lookup(name)
			fmt.Println("\tConfigure(), err =", mod.Configure())
			fmt.Println("\tStart(), err = ", mod.Start())
			fmt.Println("\tStop(), err = ", mod.Stop())
		}

		return nil
	},
}
