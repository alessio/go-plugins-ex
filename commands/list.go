package commands

import (
	"fmt"
	"github.com/alessio/go-plugins-ex/registry"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list plugins",
	Long:  `This command lists available plugins.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := registry.LoadFromDir(pluginsDir); err != nil {
			return err
		}

		for _, modName := range registry.Modules() {
			fmt.Println(modName)
		}

		return nil
	},
}
