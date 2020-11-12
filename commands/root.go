/*
Copyright © 2020 Alessio Treglia <alessio@debian.org>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const progName = "go-plugins-ex"

var (
	cfgFile    string
	configDir  string
	pluginsDir string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   progName,
	Short: "sample application",
	Long:  `This program demos Go plugins.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	configDir = ConfigDir()
	pluginsDir = filepath.Join(configDir, "plugins")
	configFile := filepath.Join(configDir, "config")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", configFile, "configuratiojn file")
	rootCmd.PersistentFlags().StringVar(&pluginsDir, "plugins-dir", pluginsDir, "plugins directory")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(listCmd, demoCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".go-plugins-ex" (without extension).
		viper.AddConfigPath(ConfigDir())
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		_, _ = fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	// make sure plugins dir exists
	if err := os.MkdirAll(pluginsDir, 0755); err != nil {
		panic(err)
	}
}

func ConfigDir() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	return filepath.Join(configDir, progName)
}
