package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var demoCmd = &cobra.Command{
	Use:   "demo",
	Short: "demo plugins",
	Long:  `Configure and execute modules one after another.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		plugins, err := availablePlugins()
		if err != nil {
			return err
		}

		var executables Plugins
		for _, plug := range plugins {
			if !plug.Compatible() {
				cmd.PrintErrf("%s is not binary compatible, ignoring", plug.Filename())
				continue
			}

			executables = append(executables, plug)
		}

		for _, plug := range executables {
			fmt.Println("Configure(), err =", plug.Configure())
			fmt.Println("Start(), err = ", plug.Start())
			fmt.Println("Stop(), err = ", plug.Stop())
		}

		return nil
	},
}
